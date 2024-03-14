package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type Member struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m Member) Create(req do.MemberInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m Member) Update(req do.MemberInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m Member) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m Member) List(req do.MemberListReq) (resp []*do.MemberInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m Member) UpdateStatus(ID int64, status int8) error {
	//TODO implement me
	panic("implement me")
}

func (m Member) InfoByID(ID int64) (roleInfo *do.MemberInfo, err error) {
	//TODO implement me
	panic("implement me")
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
