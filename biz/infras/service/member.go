package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/member"
	"saas/pkg/db/ent"
	member2 "saas/pkg/db/ent/member"
	"saas/pkg/db/ent/membercontract"
	"saas/pkg/db/ent/memberdetails"
	"saas/pkg/minio"

	"saas/pkg/db/ent/predicate"
	"saas/pkg/encrypt"
	"strconv"
	"time"
)

type Member struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

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

func NewMember(ctx context.Context, c *app.RequestContext) do.Member {
	return &Member{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

//func (m Member) ProductSearch(members []int64) (info *do.ProductInfo, err error) {
//	//m.db.MemberProduct.Query().Where(memberproduct.MemberIDIn(members...)).All(m.ctx)
//	m.db.Debug().MemberProduct.Query().
//		Modify(func(s *sql.Selector) {
//			s.GroupBy(memberproduct.FieldProductID)
//			s.Having(
//				sql.In(
//					memberproduct.FieldMemberID,
//					members,
//				),
//			)
//		}).
//		All(m.ctx)
//	return
//}
//
//func (m Member) PropertySearch(memberProducts []int64) (info *do.PropertyInfo, err error) {
//	m.db.MemberProductProperty.Query().Where(memberproductproperty.MemberProductIDIn(memberProducts...)).All(m.ctx)
//
//	return
//}

func (m Member) CreateMember(req member.CreateOrUpdateMemberReq) error {

	parsedTime, _ := time.Parse(time.DateTime, req.Birthday)

	var gender int64
	if req.Gender == "女性" {
		gender = 0
	} else if req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
	}

	tx, err := m.db.Tx(m.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	mer := tx.Member.Create().
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetNickname(req.Name).
		SetStatus(req.Status).
		SetName(req.Name).
		SetCondition(0)

	if req.Password != "" {
		password, _ := encrypt.Crypt(req.Mobile)
		mer.SetPassword(password)
	}
	noe, err := mer.Save(m.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}

	_, err = tx.MemberDetails.Create().
		SetEmail(req.Email).
		SetWecom(req.Wecom).
		SetCreateID(req.CreateId).
		SetGender(gender).
		SetBirthday(parsedTime).
		SetInfo(noe).
		Save(m.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create Member Details failed"))
		return err
	}

	//_, err = tx.Face.Create().
	//	SetMemberFaces(noe).
	//	Save(m.ctx)
	//
	//if err != nil {
	//	err = rollback(tx, errors.Wrap(err, "create Face failed"))
	//	return err
	//}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (m Member) UpdateMember(req member.CreateOrUpdateMemberReq) error {

	tx, err := m.db.Tx(m.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	up := tx.Member.Update().
		Where(member2.IDEQ(req.ID)).
		SetAvatar(req.Avatar).
		SetStatus(req.Status).
		SetMobile(req.Mobile)

	if req.Password != "" {
		password, _ := encrypt.Crypt(req.Password)
		up.SetPassword(password)
	}

	_, err = up.Save(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}

	_, err = tx.MemberDetails.Update().
		Where(memberdetails.MemberIDEQ(req.ID)).
		SetEmail(req.Email).
		Save(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (m Member) DeleteMember(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m Member) MemberList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error) {
	var predicates []predicate.Member
	if req.Name != "" {
		predicates = append(predicates, member2.NameEQ(req.Name))
	}
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

	m.cache.SetWithTTL("memberInfo"+strconv.Itoa(int(info.Id)), &info, 1, 1*time.Hour)
	return
}

func (m Member) entMemberInfo(v ent.Member) *member.MemberInfo {
	d, _ := m.db.MemberDetails.Query().Where(memberdetails.MemberIDEQ(v.ID)).First(m.ctx)

	var gender string
	if d.Gender == 0 {
		gender = "女性"
	} else if d.Gender == 1 {
		gender = "男性"
	} else {
		gender = "保密"
	}
	var age int64
	if !d.Birthday.IsZero() {
		age = int64(time.Now().Sub(d.Birthday).Hours() / 24 / 365)
	}
	return &member.MemberInfo{
		Gender:    gender,
		Avatar:    minio.URLconvert(m.ctx, m.c, v.Avatar),
		Name:      v.Name,
		Mobile:    v.Mobile,
		Id:        v.ID,
		Nickname:  v.Nickname,
		Condition: v.Condition,
		Age:       age,
		Birthday:  d.Birthday.Format(time.DateTime),
	}
}

func (m Member) MemberSearch(option string, value string) (memberInfo *member.MemberInfo, err error) {
	//memberInfo = new(member.MemberInfo)
	//switch option {
	//case "1":
	//	id, err := strconv.ParseInt(value, 10, 64)
	//	if err != nil {
	//		return memberInfo, errors.New("会员账号ID不正确")
	//	}
	//	return m.Info(id)
	//case "2":
	//	id, err := m.db.Member.Query().Where(member2.Mobile(value)).FirstID(m.ctx)
	//	if err != nil {
	//		return memberInfo, errors.New("未找到该会员")
	//	}
	//	return m.Info(id)
	//case "3":
	//	p, err := m.db.MemberProduct.Query().Where(memberproduct.SnEQ(value)).First(m.ctx)
	//	if err != nil {
	//		return memberInfo, errors.New("未找到该会员商品")
	//	}
	//	return m.Info(p.MemberID)
	//case "4":
	//	p, err := m.db.MemberProductProperty.Query().Where(memberproductproperty.SnEQ(value)).First(m.ctx)
	//	if err != nil {
	//		return memberInfo, errors.New("未找到该会员商品属性")
	//	}
	//	return m.Info(p.MemberID)
	//default:
	//	return memberInfo, errors.New("搜索类型不规范")
	//}
	return
}

func (m Member) ChangePassword(ID int64, oldPassword, newPassword string) error {

	targetMember, err := m.db.Member.Query().Where(member2.IDEQ(ID)).First(m.ctx)
	if err != nil {
		return errors.Wrap(err, "targetUser not found")
	}
	// check old password
	if ok := encrypt.VerifyPassword(oldPassword, targetMember.Password); !ok {
		err = errors.New("wrong old password")
		return err
	}
	// update password
	password, _ := encrypt.Crypt(newPassword)
	_, err = m.db.Member.Update().Where(member2.IDEQ(ID)).SetPassword(password).Save(m.ctx)

	return err
}

func (m Member) UpdateMemberStatus(ID int64, status int64) error {
	_, err := m.db.Member.Update().Where(member2.IDEQ(ID)).SetStatus(status).Save(m.ctx)
	return err
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
func (m Member) MemberNode(ID int64) (info *member.MemberNode, err error) {

	memberNode, exist := m.cache.Get("memberNode" + strconv.Itoa(int(ID)))
	if exist {
		if u, ok := memberNode.(*member.MemberNode); ok {
			return u, nil
		}
	}
	memberEnt, err := m.db.Member.Query().Where(member2.IDEQ(ID)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member failed")
		return info, err
	}
	memberDetails, err := m.db.MemberDetails.Query().Where(memberdetails.MemberIDEQ(ID)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member Details failed")
		return info, err
	} else {

	}
	info.CreatedAt = memberEnt.CreatedAt.Format(time.DateTime)
	//进馆最后期限时间
	info.EntryDeadlineTime = memberDetails.EntryDeadlineTime.Format(time.DateTime)
	//最后一次进馆时间
	info.EntryLastTime = memberDetails.EntryLastTime.Format(time.DateTime)
	//最后一次上课时间
	info.ClassLastTime = memberDetails.ClassLastTime.Format(time.DateTime)

	////消费总金额
	info.MoneySum = memberDetails.MoneySum
	////首次的产品
	//info.ProductId = memberDetails.ProductId
	//info.ProductName = memberDetails.MoneySu,,,m
	//首次消费场馆
	info.ProductVenue = memberDetails.ProductVenue
	//info.ProductVenueName = memberDetails.MoneySum
	//进馆总次数
	info.EntrySum = memberDetails.EntrySum

	//关联员工
	info.RelationUid = memberDetails.CreateID
	//info.RelationUname = memberDetails.MoneySum
	//关联会员
	info.RelationMid = memberDetails.RelationMid
	//info.RelationMname = memberDetails.MoneySum
	info.CreatedAt = memberDetails.CreatedAt.String()
	info.UpdatedAt = memberDetails.UpdatedAt.String()
	return
}
