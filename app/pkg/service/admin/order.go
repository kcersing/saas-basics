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

type Order struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (o Order) Create(req do.OrderInfo) error {
	//TODO implement me
	panic("implement me")
}

func (o Order) Update(req do.OrderInfo) error {
	//TODO implement me
	panic("implement me")
}

func (o Order) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (o Order) List(req do.OrderListReq) (resp []*do.OrderInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (o Order) UpdateStatus(ID int64, status int8) error {
	//TODO implement me
	panic("implement me")
}

func (o Order) InfoByID(ID int64) (roleInfo *do.OrderInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func NewOrder(ctx context.Context, c *app.RequestContext) do.Order {
	return &Order{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
