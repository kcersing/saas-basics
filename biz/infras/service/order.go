package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/order"
)

type Order struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewOrder(ctx context.Context, c *app.RequestContext) do.Order {
	return &Order{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (o Order) Info(id int64) (info *order.OrderInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (o Order) Update(req order.UpdateOrderReq) error {
	//TODO implement me
	panic("implement me")
}

func (o Order) List() (resp []*order.OrderInfo, total int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (o Order) UpdateStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (o Order) CreateContestOrder() (resp order.OrderInfo, err error) {
	//TODO implement me
	panic("implement me")
}
