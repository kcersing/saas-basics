package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/admin/pkg/minio"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/memberdetails"
	"saas/pkg/db/ent/memberproduct"
	"saas/pkg/db/ent/memberproductproperty"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/encrypt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Member struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m Member) Create(req do.CreateOrUpdateMemberReq) error {
	tx, err := m.db.Tx(m.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	password, _ := encrypt.Crypt(req.Mobile)
	noe, err := tx.Member.Create().
		SetPassword(password).
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetNickname(req.Nickname).
		SetStatus(req.Status).
		SetName(req.Name).
		SetCondition(0).
		Save(m.ctx)
	if err != nil {
		err = rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}

	parsedTime, _ := time.Parse(time.DateOnly, req.Birthday)

	var gender int64
	if req.Gender == "女性" {
		gender = 0
	} else if req.Gender == "男性" {
		gender = 1
	} else {
		gender = 2
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
		err = rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (m Member) Update(req do.CreateOrUpdateMemberReq) error {
	password, _ := encrypt.Crypt(req.Password)
	_, err := m.db.Member.Update().
		Where(member.IDEQ(req.ID)).
		SetPassword(password).
		SetStatus(req.Status).
		SetMobile(req.Mobile).
		Save(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}

	_, err = m.db.MemberDetails.Update().
		Where(memberdetails.MemberIDEQ(req.ID)).
		SetEmail(req.Email).
		Save(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}

	return nil
}

func (m Member) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m Member) List(req do.MemberListReq) (resp []*do.MemberInfo, total int, err error) {
	var predicates []predicate.Member
	if req.Name != "" {
		predicates = append(predicates, member.NameEQ(req.Name))
	}
	lists, err := m.db.Member.Query().Where(predicates...).
		Order(ent.Desc(member.FieldID)).
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

	for _, v := range resp {
		v.Avatar = minio.URLconvert(m.ctx, m.c, v.Avatar)
	}

	total, _ = m.db.Member.Query().Where(predicates...).Count(m.ctx)
	return
}

func (m Member) Info(ID int64) (memberInfo *do.MemberInfo, err error) {

	memberInfo = new(do.MemberInfo)

	userInterface, exist := m.cache.Get("memberInfo" + strconv.Itoa(int(ID)))
	if exist {
		if u, ok := userInterface.(*do.MemberInfo); ok {
			return u, nil
		}
	}
	memberEnt, err := m.db.Member.Query().Where(member.IDEQ(ID)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member failed")
		return memberInfo, err
	}
	memberDetails, err := m.db.MemberDetails.Query().Where(memberdetails.MemberIDEQ(ID)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member Details failed")
		return memberInfo, err
	}
	err = copier.Copy(&memberInfo, &memberDetails)
	if err != nil {
		err = errors.Wrap(err, "copy member info failed")
		return memberInfo, err
	}

	memberInfo.Avatar = minio.URLconvert(m.ctx, m.c, memberEnt.Avatar)
	memberInfo.Name = memberEnt.Name
	memberInfo.Mobile = memberEnt.Mobile
	memberInfo.ID = memberEnt.ID
	memberInfo.Nickname = memberEnt.Nickname
	memberInfo.CreatedAt = memberEnt.CreatedAt.Format(time.DateTime)
	memberInfo.Condition = memberEnt.Condition
	memberInfo.Birthday = memberDetails.Birthday.Format(time.DateOnly)

	if memberDetails.Gender == 0 {
		memberInfo.Gender = "女性"
	} else if memberDetails.Gender == 1 {
		memberInfo.Gender = "男性"
	} else {
		memberInfo.Gender = "保密"
	}

	if !memberDetails.Birthday.IsZero() {
		memberInfo.Age = int64(time.Now().Sub(memberDetails.Birthday).Hours() / 24 / 365)
	}

	m.cache.SetWithTTL("memberInfo"+strconv.Itoa(int(memberInfo.ID)), &memberInfo, 1, 1*time.Hour)
	return

}

func (m Member) Search(option string, value string) (memberInfo *do.MemberInfo, err error) {
	memberInfo = new(do.MemberInfo)
	switch option {
	case "1":
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return memberInfo, errors.New("会员账号ID不正确")
		}
		return m.Info(id)
	case "2":
		id, err := m.db.Member.Query().Where(member.Mobile(value)).FirstID(m.ctx)
		if err != nil {
			return memberInfo, errors.New("未找到该会员")
		}
		return m.Info(id)
	case "3":
		p, err := m.db.MemberProduct.Query().Where(memberproduct.SnEQ(value)).First(m.ctx)
		if err != nil {
			return memberInfo, errors.New("未找到该会员商品")
		}
		return m.Info(p.MemberID)
	case "4":
		p, err := m.db.MemberProductProperty.Query().Where(memberproductproperty.SnEQ(value)).First(m.ctx)
		if err != nil {
			return memberInfo, errors.New("未找到该会员商品属性")
		}
		return m.Info(p.MemberID)
	default:
		return memberInfo, errors.New("搜索类型不规范")
	}

}
func (m Member) ChangePassword(ID int64, oldPassword, newPassword string) error {

	targetMember, err := m.db.Member.Query().Where(member.IDEQ(ID)).First(m.ctx)
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
	_, err = m.db.Member.Update().Where(member.IDEQ(ID)).SetPassword(password).Save(m.ctx)

	return err
}

func (m Member) UpdateStatus(ID int64, status int64) error {
	_, err := m.db.Member.Update().Where(member.IDEQ(ID)).SetStatus(status).Save(m.ctx)
	return err
}

func NewMember(ctx context.Context, c *app.RequestContext) do.Member {
	return &Member{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
