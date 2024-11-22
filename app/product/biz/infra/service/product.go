package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
	"saas/pkg/db/ent/user"
	"saas/pkg/db/ent/venue"
	"strconv"
	"time"
)

type Product struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (p Product) Info(id int64) (info *do.ProductInfo, err error) {
	inter, exist := p.cache.Get("productInfo" + strconv.Itoa(int(id)))
	if exist {
		if v, ok := inter.(*do.ProductInfo); ok {
			return v, nil
		}
	}
	one, err := p.db.Product.Query().Where(product.IDEQ(id)).First(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "get product failed")
		return info, err
	}

	info = &do.ProductInfo{
		ID:          one.ID,
		Name:        one.Name,
		Pic:         one.Pic,
		Description: one.Description,
		Price:       one.Price,
		Stock:       one.Stock,
		Status:      one.Status,
		CreateID:    one.CreateID,
	}
	if one.CreateID != 0 {
		create, err := p.db.User.Query().Where(user.IDEQ(one.CreateID)).First(p.ctx)
		if err != nil {
			err = errors.Wrap(err, "get user failed")
			return info, err
		}
		info.CreateName = create.Nickname
	}

	p.cache.SetWithTTL("productInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
	return
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
		AddVenues(p.db.Venue.Query().Where(venue.IDIn(req.VenueId...)).AllX(p.ctx)...).
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
		ClearVenues().
		AddVenues(p.db.Venue.Query().Where(venue.IDIn(req.VenueId...)).AllX(p.ctx)...).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "update Product Property failed")
		return err
	}

	return nil
}
func (p Product) UpdatePropertyStatus(ID int64, status int64) error {
	_, err := p.db.ProductProperty.Update().Where(productproperty.IDEQ(ID)).SetStatus(int64(status)).Save(p.ctx)
	return err
}
func (p Product) DeleteProperty(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) PropertyList(req do.ProductListReq) (resp []*do.PropertyInfo, total int, err error) {

	hlog.Info(req)
	var predicates []predicate.ProductProperty
	if req.Name != "" {
		predicates = append(predicates, productproperty.NameEQ(req.Name))
	}
	if len(req.Status) > 0 {
		predicates = append(predicates, productproperty.StatusIn(req.Status...))
	}
	if len(req.CreatedTime) > 0 {
		s, _ := time.ParseInLocation(time.DateTime, req.CreatedTime[0], time.Local)
		d, _ := time.ParseInLocation(time.DateTime, req.CreatedTime[1], time.Local)
		predicates = append(predicates, productproperty.CreatedAtGTE(s))
		predicates = append(predicates, productproperty.CreatedAtLTE(d))
	}

	if len(req.Venue) > 0 {
		predicates = append(predicates, productproperty.HasVenuesWith(venue.IDIn(req.Venue...)))
	}

	if req.Type != "" {
		predicates = append(predicates, productproperty.TypeEQ(req.Type))
	}

	lists, err := p.db.ProductProperty.
		Query().
		Where(predicates...).
		Order(ent.Desc(productproperty.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Product info failed")
		return resp, 0, err
	}

	for i, v := range lists {
		venues, err := v.QueryVenues().Select(venue.FieldID, venue.FieldName).All(p.ctx)
		if err == nil {
			var ven []do.PropertyVenue
			err = copier.Copy(&ven, &venues)
			resp[i].Venue = ven

			for i2, v := range ven {
				if i2 == 0 {
					resp[i].Venues = v.Name
				} else {
					resp[i].Venues += ", " + v.Name
				}
				resp[i].VenueId = append(resp[i].VenueId, v.ID)
			}

		}
	}

	total, _ = p.db.ProductProperty.Query().Where(predicates...).Count(p.ctx)
	return

}

// propertys
func (p Product) Create(req do.CreateOrUpdateProduct) error {
	var property []int64
	property = append(property, req.CardProperty)
	property = append(property, req.CourseProperty...)
	property = append(property, req.ClassProperty...)
	properties, err := p.db.ProductProperty.Query().Where(
		productproperty.IDIn(property...),
	).All(p.ctx)
	if err != nil {
		return errors.Wrap(err, "未找到属性")
	}
	_, err = p.db.Product.Create().
		SetName(req.Name).
		SetPic(req.Pic).
		SetDescription(req.Description).
		SetPrice(req.Price).
		SetStock(req.Stock).
		AddPropertys(properties...).
		SetCreateID(req.CreateID).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Product failed")
		return err
	}

	return nil
}

func (p Product) Update(req do.ProductInfo) error {
	_, err := p.db.Product.Update().
		Where(product.IDEQ(req.ID)).
		SetName(req.Name).
		SetPic(req.Pic).
		SetDescription(req.Description).
		SetStatus(req.Status).
		SetPrice(req.Price).
		SetStock(req.Stock).
		SetCreateID(req.CreateID).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "update Product failed")
		return err
	}

	return nil
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
		Order(ent.Desc(product.FieldCreatedAt)).
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

	for i, v := range lists {
		//propertys

		ds, _ := v.QueryPropertys().IDs(p.ctx)

		all, err := p.db.ProductProperty.Query().Where(productproperty.IDIn(ds...)).All(p.ctx)
		if err != nil {
			return nil, 0, err
		}

		for _, p := range all {
			pi := do.PropertyInfo{
				ID:       p.ID,
				Name:     p.Name,
				Price:    p.Price,
				Duration: p.Duration,
				Length:   p.Length,
				Count:    p.Count,
				Type:     p.Type,
				Data:     p.Data,
				Status:   p.Status,
			}
			if p.Type == "card" {
				resp[i].CardProperty = append(resp[i].CourseProperty, pi)
			}
			if p.Type == "course" {
				resp[i].CourseProperty = append(resp[i].CourseProperty, pi)
			}
			if p.Type == "class" {
				resp[i].ClassProperty = append(resp[i].CourseProperty, pi)
			}
		}

		//v.CourseProperty
		//v.ClassProperty
		//v.CardProperty
	}

	total, _ = p.db.Product.Query().Where(predicates...).Count(p.ctx)
	return

}

func (p Product) UpdateStatus(ID int64, status int64) error {
	_, err := p.db.Product.Update().Where(product.IDEQ(ID)).SetStatus(int64(status)).Save(p.ctx)
	return err
}

func NewProduct(ctx context.Context, c *app.RequestContext) do.Product {
	return &Product{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
