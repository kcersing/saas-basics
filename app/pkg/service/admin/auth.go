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

type Auth struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (a Auth) UpdateApiAuth(roleIDStr string, infos []*do.ApiAuthInfo) error {
	//TODO implement me
	panic("implement me")
}

func (a Auth) ApiAuth(roleIDStr string) (infos []*do.ApiAuthInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (a Auth) UpdateMenuAuth(roleID uint64, menuIDs []uint64) error {
	//TODO implement me
	panic("implement me")
}

func (a Auth) MenuAuth(roleID uint64) (menuIDs []uint64, err error) {
	//TODO implement me
	panic("implement me")
}

func NewAuth(ctx context.Context, c *app.RequestContext) do.Auth {
	return &Auth{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
