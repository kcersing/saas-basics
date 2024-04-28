package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/encrypt"

	"github.com/pkg/errors"
)

type Member struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m Member) Create(req do.MemberInfo) error {
	password, _ := encrypt.Crypt(req.Password)
	_, err := m.db.Member.Create().
		SetPassword(password).
		SetStatus(req.Status).
		SetMobile(req.Mobile).
		SetEmail(req.Email).
		Save(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}
	return nil
}

func (m Member) Update(req do.MemberInfo) error {
	password, _ := encrypt.Crypt(req.Password)
	_, err := m.db.Member.Update().
		Where(member.IDEQ(req.ID)).
		SetPassword(password).
		SetStatus(req.Status).
		SetMobile(req.Mobile).
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
	hlog.Info(req.PageSize)
	lists, err := m.db.Member.Query().Where(predicates...).
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
	total, _ = m.db.Member.Query().Where(predicates...).Count(m.ctx)
	return
}

func (m Member) Info(ID int64) (memberInfo *do.MemberInfo, err error) {
	//TODO implement me
	panic("implement me")
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
