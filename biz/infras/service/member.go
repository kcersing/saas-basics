package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/memberdetails"
	"saas/biz/dal/db/ent/memberprofile"
	user2 "saas/biz/dal/db/ent/user"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/member"
	"saas/pkg/encrypt"
	"saas/pkg/enums"
	"time"
)

type Member struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m Member) UpdateMemberFollow(req *member.UpdateMemberFollowReq) (err error) {

	user, err := m.db.User.Query().Where(user2.IDEQ(req.FollowId)).First(m.ctx)
	if user == nil {
		return errors.New("未找到该员工")
	}

	_, err = m.db.MemberDetails.
		Update().
		Where(memberdetails.MemberID(req.MemberId)).
		SetRelationUID(req.FollowId).
		SetRelationUname(user.Name).
		Save(m.ctx)
	return err
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

	birthday, _ := time.Parse(time.DateTime, req.Birthday)

	var gender = enums.ReturnMemberGenderKey(req.Gender)

	tx, err := m.db.Tx(m.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	noe, err := tx.Member.Create().
		SetMobile(req.Mobile).
		SetUsername(req.Mobile).
		SetName(req.Name).
		SetCondition(1).
		Save(m.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}

	_, err = tx.MemberProfile.Create().
		SetMember(noe).
		SetGrade(gender).
		SetBirthday(birthday).
		SetSource(req.Source).
		SetGrade(req.Grade).
		SetIntention(req.Intention).
		Save(m.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create Member Profile failed"))
		return err
	}

	_, err = tx.MemberDetails.Create().
		SetMember(noe).
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

	birthday, _ := time.Parse(time.DateTime, req.Birthday)

	var gender = enums.ReturnMemberGenderKey(req.Gender)

	tx, err := m.db.Tx(m.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	//if req.Password != "" {
	//	password, _ := encrypt.Crypt(req.Password)
	//}

	_, err = tx.Member.Update().
		Where(member2.IDEQ(req.ID)).
		SetMobile(req.Mobile).
		SetUsername(req.Mobile).
		SetName(req.Name).
		SetCondition(1).
		Save(m.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}

	_, err = tx.MemberProfile.Update().
		Where(memberprofile.MemberID(req.ID)).
		SetGrade(gender).
		SetBirthday(birthday).
		SetSource(req.Source).
		SetGrade(req.Grade).
		SetIntention(req.Intention).
		Save(m.ctx)

	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create Member Profile failed"))
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (m Member) DeleteMember(id int64) error {
	_, err := m.db.Member.Update().Where(member2.IDEQ(id)).SetDelete(1).Save(m.ctx)
	return err
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
