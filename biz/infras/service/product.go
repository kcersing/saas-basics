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
	"saas/idl_gen/model/admin/product"
)

type Product struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (p Product) CreateProduct(req product.ProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) UpdateProduct(req product.ProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) DeleteProduct(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) UpdateProductStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) ProductList(req *product.ProductListReq) (resp []*product.ProductInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (p Product) ProductInfo(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) ProductListExport(req *product.ProductListReq) (resp string, err error) {
	//TODO implement me
	panic("implement me")
}

func NewProduct(ctx context.Context, c *app.RequestContext) do.Product {
	return &Product{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
