package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/product"
	"saas/pkg/db/ent/productproperty"
)

type Product struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (p Product) CreateProperty(req do.PropertyInfo) error {

	_, err := p.db.ProductProperty.Create().
		SetName(req.Name).
		SetType(req.Type).
		SetLength(req.Length).
		SetStatus(req.Status).
		SetDuration(req.Duration).
		SetCount(req.Count).
		SetPrice(req.Price).
		SetData("").
		Save(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}
	return nil

}

func (p Product) UpdateProperty(req do.PropertyInfo) error {
	_, err := p.db.ProductProperty.Update().
		Where(productproperty.IDEQ(req.ID)).
		SetName(req.Name).
		SetType(req.Type).
		SetLength(req.Length).
		SetStatus(req.Status).
		SetDuration(req.Duration).
		SetCount(req.Count).
		SetPrice(req.Price).
		SetData("").
		Save(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Product Property failed")
		return err
	}

	return nil
}

func (p Product) DeleteProperty(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) PropertyList(req do.ProductListReq) (resp []*do.PropertyInfo, total int, err error) {

	var predicates []predicate.ProductProperty
	if req.Name != "" {
		predicates = append(predicates, productproperty.NameEQ(req.Name))
	}
	lists, err := p.db.ProductProperty.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy product info failed")
		return resp, 0, err
	}
	total, _ = p.db.ProductProperty.Query().Where(predicates...).Count(p.ctx)
	return

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
	var predicates []predicate.Product
	if req.Name != "" {
		predicates = append(predicates, product.NameEQ(req.Name))
	}
	lists, err := p.db.Product.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy product info failed")
		return resp, 0, err
	}
	total, _ = p.db.Product.Query().Where(predicates...).Count(p.ctx)
	return

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
