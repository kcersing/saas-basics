package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type Login struct {
	ctx  context.Context
	c    *app.RequestContext
	salt string
	db   *ent.Client
}

func (l Login) Login(username, password string) (res *do.LoginResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (l Login) LoginByOAuth(provider, credential string) (res *do.LoginResp, err error) {
	//TODO implement me
	panic("implement me")
}

func NewLogin(ctx context.Context, c *app.RequestContext) do.Login {
	return &Login{
		ctx:  ctx,
		c:    c,
		salt: config.GlobalServerConfig.MysqlInfo.Salt,
		db:   infras.DB,
	}
}
