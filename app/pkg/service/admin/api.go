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

type Api struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (a Api) Create(req do.ApiInfo) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) Update(req do.ApiInfo) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) Delete(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (a Api) List(req do.ListApiReq) (resp []*do.ApiInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func NewApi(ctx context.Context, c *app.RequestContext) do.Api {
	return &Api{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
