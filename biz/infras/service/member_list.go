package service

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/membercontract"
	"saas/biz/dal/db/ent/memberprofile"
	"saas/biz/dal/db/ent/predicate"
	user2 "saas/biz/dal/db/ent/user"
	"saas/biz/dal/minio"
	"saas/idl_gen/model/member"
	"saas/pkg/enums"
	"strconv"
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
	mInfo, _ := v.QueryMember().First(m.ctx)
	//mpInfo, _ := v.QueryMemberProduct().First(m.ctx)
	return &member.MemberContractInfo{
		Name:       &v.Name,
		MemberId:   &v.MemberID,
		MemberName: &mInfo.Name,
		VenueId:    &v.VenueID,
		VenueName:  &vInfo.Name,
		//MemberProductId:   &v.MemberProductID,
		//MemberProductName: nil,
		ContractId: &v.ContractID,
		Sign:       &v.Sign,
		SignImg:    &first.SignImg,
		Content:    &first.Content,
	}
}

func (m Member) MemberFullList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error) {
	var predicates []predicate.Member
	if req.Name != "" {
		predicates = append(predicates, member2.NameEQ(req.Name))
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
	predicates = append(predicates, member2.Condition(2))
	lists, err := m.db.Member.Query().Where(predicates...).
		Order(ent.Desc(member2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Member info failed")
		return resp, 0, err
	}

	for i, v := range lists {
		resp[i], _ = m.MemberInfo(v.ID)
	}

	total, _ = m.db.Member.Query().Where(predicates...).Count(m.ctx)
	return
}

func (m Member) MemberPotentialList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error) {
	var predicates []predicate.Member
	if req.Name != "" {
		predicates = append(predicates, member2.NameEQ(req.Name))
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
	predicates = append(predicates, member2.Condition(1))
	lists, err := m.db.Member.Query().Where(predicates...).
		Order(ent.Desc(member2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Member info failed")
		return resp, 0, err
	}

	for i, v := range lists {
		resp[i], _ = m.MemberInfo(v.ID)
	}

	total, _ = m.db.Member.Query().Where(predicates...).Count(m.ctx)
	return
}
func (m Member) MemberInfo(id int64) (info *member.MemberInfo, err error) {

	info = new(member.MemberInfo)

	userInterface, exist := m.cache.Get("memberInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := userInterface.(*member.MemberInfo); ok {
			return u, nil
		}
	}

	memberEnt, err := m.db.Member.Query().Where(member2.IDEQ(id)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member failed")
		return info, err
	}

	info = m.entMemberInfo(*memberEnt)

	m.cache.SetWithTTL("memberInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
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
		ID:            v.ID,
		Name:          v.Name,
		Username:      v.Username,
		Mobile:        v.Mobile,
		Avatar:        minio.URLconvert(m.ctx, m.c, v.Avatar),
		Condition:     v.Condition,
		ConditionName: enums.ReturnMemberConditionValues(v.Condition),
		Status:        v.Status,
		CreatedAt:     v.CreatedAt.Format(time.DateTime),
		UpdatedAt:     v.UpdatedAt.Format(time.DateTime),
		Profile: &member.MemberProfile{
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
			MoneySum:          d.MoneySum,
			ProductId:         d.ProductID,
			ProductName:       d.ProductName,
			ProductVenue:      d.ProductVenue,
			ProductVenueName:  d.ProductVenueName,
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
