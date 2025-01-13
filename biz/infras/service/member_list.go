package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/membercontract"
	"saas/biz/dal/db/ent/memberprofile"
	"saas/biz/dal/db/ent/predicate"
	user2 "saas/biz/dal/db/ent/user"
	"saas/biz/dal/minio"
	"saas/idl_gen/model/member"
	"saas/pkg/enums"
	"saas/pkg/utils"
	"time"
)

func (m Member) ContractList(req member.MemberContractListReq) (resp []*member.MemberContractInfo, total int, err error) {
	var predicates []predicate.MemberContract
	if req.MemberId > 0 {
		predicates = append(predicates, membercontract.MemberID(req.MemberId))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, membercontract.VenueID(req.VenueId))
	}
	if req.ContractId > 0 {
		predicates = append(predicates, membercontract.ContractID(req.ContractId))
	}
	predicates = append(predicates, membercontract.Delete(0))
	lists, err := m.db.MemberContract.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Contract list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, m.entMemberContractInfo(*v))
	}

	return
}

func (m Member) entMemberContractInfo(v ent.MemberContract) *member.MemberContractInfo {
	vInfo, _ := NewVenue(m.ctx, m.c).VenueInfo(v.VenueID)
	first, _ := v.QueryContent().First(m.ctx)
	mInfo, _ := m.db.MemberProfile.Query().Where(memberprofile.MemberID(v.MemberID)).First(m.ctx)
	mpInfo, _ := v.QueryMemberProduct().First(m.ctx)
	return &member.MemberContractInfo{
		Name:              v.Name,
		MemberId:          v.MemberID,
		MemberName:        mInfo.Name,
		VenueId:           v.VenueID,
		VenueName:         vInfo.Name,
		MemberProductId:   v.MemberProductID,
		MemberProductName: mpInfo.Name,
		ContractId:        v.ContractID,
		Sign:              v.Sign,
		SignImg:           first.SignImg,
		Content:           first.Content,
		CreatedAt:         v.CreatedAt.Format(time.DateTime),
		UpdatedAt:         v.UpdatedAt.Format(time.DateTime),
	}
}

func (m Member) MemberFullList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error) {
	var predicates []predicate.Member
	if req.Name != "" {
		predicates = append(predicates, member2.HasMemberProfileWith(memberprofile.Name(req.Name)))
	}

	if req.Mobile != "" {
		predicates = append(predicates, member2.MobileEQ(req.Mobile))
	}

	if req.Source > 0 {
		predicates = append(predicates, member2.HasMemberProfileWith(memberprofile.SourceEQ(req.Source)))
	}
	if req.Intention > 0 {
		predicates = append(predicates, member2.HasMemberProfileWith(memberprofile.IntentionEQ(req.Intention)))
	}
	if req.CreatedId > 0 {
		predicates = append(predicates, member2.CreatedIDEQ(req.CreatedId))
	}
	if req.StartCreatedAt != "" && req.EndCreatedAt != "" {
		startAt, _ := time.Parse(time.DateTime, req.StartCreatedAt)
		endAt, _ := time.Parse(time.DateTime, req.EndCreatedAt)

		predicates = append(predicates, member2.CreatedAtGTE(startAt))
		predicates = append(predicates, member2.CreatedAtLTE(endAt))
	}

	predicates = append(predicates, member2.Delete(0))
	predicates = append(predicates, member2.HasMemberProfileWith(memberprofile.Condition(2)))
	lists, err := m.db.Member.Query().Where(predicates...).
		Order(ent.Desc(member2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, m.entMemberInfo(*v))
	}

	total, _ = m.db.Member.Query().Where(predicates...).Count(m.ctx)
	return
}
func (m Member) MemberFullListExport(req member.MemberListReq) (string, error) {

	exportFilePath, domain := utils.ExportFilePath("线索导出")

	resp, total, _ := m.MemberPotentialList(req)

	if total == 0 {
		return "", errors.New("暂无数据")
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return "", err
	}
	tale := []interface{}{
		"学员姓名",
		"手机号",
		"来源",
		"成为学员时间",
		"性别",
		"出生日期",
		"年纪",
	}

	err = f.SetSheetRow("Sheet1", cell, &tale)
	if err != nil {
		return "", err
	}

	for idx, row := range resp {
		cell, err := excelize.CoordinatesToCellName(1, idx+2)
		if err != nil {
			return "", err
		}
		r := []interface{}{
			row.Name,
			row.Mobile,
			row.Profile.SourceName,
			row.Details.FirstAt,
			row.Profile.Gender,
			row.Profile.Birthday,
			row.Profile.GradeName,
		}

		err = f.SetSheetRow("Sheet1", cell, &r)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	//Save spreadsheet by the given path.
	if err := f.SaveAs(exportFilePath); err != nil {
		fmt.Println(err)
	}
	return domain, nil
}

func (m Member) MemberPotentialList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error) {
	var predicates []predicate.Member
	if req.Name != "" {
		predicates = append(predicates, member2.HasMemberProfileWith(memberprofile.NameEQ(req.Name)))
	}

	if req.Mobile != "" {
		predicates = append(predicates, member2.MobileEQ(req.Mobile))
	}

	if req.Source > 0 {
		predicates = append(predicates, member2.HasMemberProfileWith(memberprofile.SourceEQ(req.Source)))
	}
	if req.Intention > 0 {
		predicates = append(predicates, member2.HasMemberProfileWith(memberprofile.IntentionEQ(req.Intention)))
	}
	if req.CreatedId > 0 {
		predicates = append(predicates, member2.CreatedIDEQ(req.CreatedId))
	}
	if req.StartCreatedAt != "" && req.EndCreatedAt != "" {
		startAt, _ := time.Parse(time.DateTime, req.StartCreatedAt)
		endAt, _ := time.Parse(time.DateTime, req.EndCreatedAt)

		predicates = append(predicates, member2.CreatedAtGTE(startAt))
		predicates = append(predicates, member2.CreatedAtLTE(endAt))
	}

	predicates = append(predicates, member2.Delete(0))
	predicates = append(predicates, member2.HasMemberProfileWith(memberprofile.Condition(1)))
	lists, err := m.db.Member.Query().Where(predicates...).
		Order(ent.Desc(member2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, m.entMemberInfo(*v))
	}

	total, _ = m.db.Member.Query().Where(predicates...).Count(m.ctx)
	return
}

func (m Member) MemberPotentialListExport(req member.MemberListReq) (string, error) {

	exportFilePath, domain := utils.ExportFilePath("线索导出")

	resp, total, _ := m.MemberPotentialList(req)

	if total == 0 {
		return "", errors.New("暂无数据")
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return "", err
	}
	tale := []interface{}{
		"学生姓名",
		"手机号",
		"来源",
		"意向",
		"添加时间",
		"添加人",
		"跟进人",
		"性别",
		"出生日期",
		"年纪",
	}

	err = f.SetSheetRow("Sheet1", cell, &tale)
	if err != nil {
		return "", err
	}

	for idx, row := range resp {
		cell, err := excelize.CoordinatesToCellName(1, idx+2)
		if err != nil {
			return "", err
		}
		r := []interface{}{
			row.Name,
			row.Mobile,
			row.Profile.SourceName,
			row.Profile.IntentionName,
			row.CreatedAt,
			row.Details.CreatedName,
			row.Details.RelationUname,
			row.Profile.Gender,
			row.Profile.Birthday,
			row.Profile.GradeName,
		}

		err = f.SetSheetRow("Sheet1", cell, &r)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	//Save spreadsheet by the given path.
	if err := f.SaveAs(exportFilePath); err != nil {
		fmt.Println(err)
	}
	return domain, nil
}

func (m Member) MemberInfo(id int64) (info *member.MemberInfo, err error) {

	memberEnt, err := m.db.Member.Query().Where(member2.IDEQ(id)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member failed")
		return nil, err
	}

	info = m.entMemberInfo(*memberEnt)

	return
}

func (m Member) entMemberInfo(v ent.Member) *member.MemberInfo {

	d, _ := v.QueryMemberDetails().First(m.ctx)
	p, _ := v.QueryMemberProfile().First(m.ctx)

	var gender = enums.ReturnMemberGenderValues(p.Gender)

	var age int64
	if !p.Birthday.IsZero() {
		age = int64(time.Now().Sub(p.Birthday).Hours() / 24 / 365)
	}

	var gradeName, intentionName, sourceName, createdName string

	if p.Grade > 0 {
		gradeName = NewDictionary(m.ctx, m.c).GetDictionaryDetailTitle(p.Grade)
	}
	if p.Intention > 0 {
		intentionName = NewDictionary(m.ctx, m.c).GetDictionaryDetailTitle(p.Intention)
	}
	if p.Source > 0 {
		sourceName = NewDictionary(m.ctx, m.c).GetDictionaryDetailTitle(p.Source)
	}
	if v.CreatedID > 0 {
		createdName = m.db.User.Query().Where(user2.IDEQ(v.CreatedID)).FirstX(m.ctx).Name
	}

	return &member.MemberInfo{
		ID:       v.ID,
		Name:     p.Name,
		Username: v.Username,
		Mobile:   v.Mobile,
		Avatar:   minio.URLconvert(m.ctx, m.c, v.Avatar),

		Status:    v.Status,
		CreatedAt: v.CreatedAt.Format(time.DateTime),
		UpdatedAt: v.UpdatedAt.Format(time.DateTime),
		Profile: &member.MemberProfile{
			Condition:        p.Condition,
			ConditionName:    enums.ReturnMemberConditionValues(p.Condition),
			MobileAscription: p.MobileAscription,
			FatherName:       p.FatherName,
			MotherName:       p.MotherName,
			Grade:            p.Grade,
			Intention:        p.Intention,
			Source:           p.Source,
			GradeName:        gradeName,
			IntentionName:    intentionName,
			SourceName:       sourceName,
			Email:            p.Email,
			Gender:           gender,
			Age:              age,
			Wecom:            p.Wecom,
			Birthday:         p.Birthday.Format(time.DateOnly),
			ID:               p.ID,
		},
		Details: &member.MemberDetails{
			MoneySum:    d.MoneySum,
			ProductId:   d.ProductID,
			ProductName: d.ProductName,

			EntrySum:          d.EntrySum,
			EntryLastTime:     d.EntryLastTime.Format(time.DateTime),
			EntryDeadlineTime: d.EntryDeadlineTime.Format(time.DateTime),
			ClassLastTime:     d.ClassLastTime.Format(time.DateTime),
			RelationUid:       d.RelationUID,
			RelationUname:     d.RelationUname,
			RelationMid:       d.RelationMid,
			RelationMname:     d.RelationMame,
			CreatedId:         v.CreatedID,
			CreatedName:       createdName,
			FirstAt:           d.FirstTime.Format(time.DateTime),
		},
	}

}

func (m Member) MemberPrivacy(ID int64) (info *member.MemberPrivacy, err error) {
	//	faceFirst, err := m.db.Face.Query().Where(face.MemberIDEQ(id)).First(m.ctx)
	//	if err != nil {
	//		err = errors.Wrap(err, "get member face failed")
	//		return info, err
	//	} else {
	//		info.IdentityCard = faceFirst.IdentityCard
	//		info.FaceIdentityCard = faceFirst.FaceIdentityCard
	//		info.BackIdentityCard = faceFirst.BackIdentityCard
	//		info.FacePic = faceFirst.FacePic
	//		info.FaceEigenvalue = faceFirst.FaceEigenvalue
	//		info.FacePicUpdatedTime = faceFirst.FacePicUpdatedTime.Format(time.DateTime)
	//	}
	return
}
