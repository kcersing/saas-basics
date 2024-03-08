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

type Product struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (p Product) CreateProperty(req do.PropertyInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) UpdateProperty(req do.PropertyInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) DeleteProperty(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) PropertyList(req do.ProductListReq) (resp []*do.PropertyInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (p Product) Create(req do.ProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) Update(req do.ProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) List(req do.ProductListReq) (resp []*do.ProductInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (p Product) UpdateStatus(ID int64, status int8) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) InfoByID(ID int64) (roleInfo *do.ProductInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func NewProduct(ctx context.Context, c *app.RequestContext) do.Product {
	return &Product{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
