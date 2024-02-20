package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type Captcha struct {
	ctx  context.Context
	c    *app.RequestContext
	salt string
	db   *ent.Client
}

func (c Captcha) GetCaptcha() (id, b64s string, err error) {
	//TODO implement me
	panic("implement me")
}

func NewCaptcha(ctx context.Context, c *app.RequestContext) do.Captcha {
	return &Captcha{
		ctx:  ctx,
		c:    c,
		salt: config.GlobalServerConfig.MysqlInfo.Salt,
		db:   infras.DB,
	}
}
