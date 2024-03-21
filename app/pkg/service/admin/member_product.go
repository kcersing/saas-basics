package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/memberproduct"
)

type MemberProduct struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m MemberProduct) Create(req do.MemberProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) Update(req do.MemberProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) List(req do.MemberProductListReq) (resp []*do.MemberProductInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) UpdateStatus(ID int64, status int64) error {
	_, err := m.db.MemberProduct.Update().Where(memberproduct.IDEQ(ID)).SetStatus(int64(status)).Save(m.ctx)
	return err
}

func (m MemberProduct) InfoByID(ID int64) (roleInfo *do.MemberProductInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func NewMemberProduct(ctx context.Context, c *app.RequestContext) do.MemberProduct {
	return &MemberProduct{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
