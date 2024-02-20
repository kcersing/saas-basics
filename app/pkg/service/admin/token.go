package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type Token struct {
	ctx  context.Context
	c    *app.RequestContext
	salt string
	db   *ent.Client
}

func (t Token) Create(req *do.TokenInfo) error {
	//TODO implement me
	panic("implement me")
}

func (t Token) Update(req *do.TokenInfo) error {
	//TODO implement me
	panic("implement me")
}

func (t Token) IsExistByUserID(userID uint64) bool {
	//TODO implement me
	panic("implement me")
}

func (t Token) Delete(userID uint64) error {
	//TODO implement me
	panic("implement me")
}

func (t Token) List(req *do.TokenListReq) (res []*do.TokenInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func NewToken(ctx context.Context, c *app.RequestContext) do.Token {
	return &Token{
		ctx:  ctx,
		c:    c,
		salt: config.GlobalServerConfig.MysqlInfo.Salt,
		db:   infras.DB,
	}
}
