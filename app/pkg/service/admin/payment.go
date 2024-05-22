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

type Payment struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (p Payment) UnifyPay(req do.UnifyPayReq) error {
	//TODO implement me
	panic("implement me")
}

func (p Payment) UniqueOutOrderNo(outOrderNo string) error {
	//TODO implement me
	panic("implement me")
}

func (p Payment) QRPay(req do.QRPayReq) {
	//TODO implement me
	panic("implement me")
}

func (p Payment) QueryOrder(outOrderNo string) {
	//TODO implement me
	panic("implement me")
}

func (p Payment) CloseOrder(outOrderNo string) {
	//TODO implement me
	panic("implement me")
}

func NewPayment(ctx context.Context, c *app.RequestContext) do.Payment {
	return &Payment{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
