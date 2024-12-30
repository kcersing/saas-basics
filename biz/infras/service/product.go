package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/productcourses"

	"saas/biz/dal/db/ent/predicate"
	product2 "saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/user"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/product"
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

func (p Product) CreateProduct(req product.CreateOrUpdateProductReq) error {
	var createdId int64
	userId, exist := p.c.Get("userId")
	if exist || userId != nil {
		uId, ok := userId.(string)
		if ok {
			createdId, _ = strconv.ParseInt(uId, 10, 64)
		}
	}
	signSalesAt, _ := time.Parse(time.DateTime, req.SignSalesAt)
	endSalesAt, _ := time.Parse(time.DateTime, req.EndSalesAt)

	_, err := p.db.Product.Create().
		SetName(req.Name).
		SetPic(req.Pic).
		SetDescription(req.Description).
		SetStock(req.Stock).
		SetStatus(req.Status).
		SetDuration(req.Duration).
		SetLength(req.Length).
		SetType(req.Type).
		SetDeadline(req.Deadline).
		SetSales(req.Sales).
		SetIsSales(req.IsSales).
		SetSignSalesAt(signSalesAt).
		SetEndSalesAt(endSalesAt).
		SetCreatedID(createdId).
		AddTagIDs(req.TagId...).
		AddContractIDs(req.ContractId...).
		SetIsLessons(req.IsLessons).
		SetSubType(req.SubType).
		SetPrice(req.Price).
		SetTimes(req.Times).
		SetIsCourse(req.IsCourse).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Product failed")
		return err
	}
	if len(req.Courses) > 0 {
		go func() {
			for _, v := range req.Courses {
				p.db.ProductCourses.Create().
					SetProductID(v.ID).
					SetType(v.Type).
					SetName(v.Name).
					SetNumber(v.Number).
					Save(p.ctx)
			}
		}()
	}
	if len(req.Lessons) > 0 {
		go func() {
			for _, v := range req.Lessons {
				p.db.ProductCourses.Create().
					SetProductID(v.ID).
					SetType(v.Type).
					SetName(v.Name).
					Save(p.ctx)
			}
		}()
	}

	return nil
}

func (p Product) UpdateProduct(req product.CreateOrUpdateProductReq) error {
	//var createdId int64
	//userId, exist := p.c.Get("userId")
	//	if exist || userId != nil {
	//	createdId = userId.(int64)
	//}
	signSalesAt, _ := time.Parse(time.DateTime, req.SignSalesAt)
	endSalesAt, _ := time.Parse(time.DateTime, req.EndSalesAt)
	_, err := p.db.Product.Update().
		Where(product2.IDEQ(req.ID)).
		SetName(req.Name).
		SetPic(req.Pic).
		SetDescription(req.Description).
		SetStock(req.Stock).
		SetStatus(req.Status).
		SetDuration(req.Duration).
		SetLength(req.Length).
		SetType(req.Type).
		SetDeadline(req.Deadline).
		SetSales(req.Sales).
		SetIsSales(req.IsSales).
		SetSignSalesAt(signSalesAt).
		SetEndSalesAt(endSalesAt).
		//SetCreatedID(createdId).
		AddTagIDs(req.TagId...).
		AddContractIDs(req.ContractId...).
		SetIsLessons(req.IsLessons).
		SetSubType(req.SubType).
		SetPrice(req.Price).
		SetTimes(req.Times).
		SetIsCourse(req.IsCourse).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "update Product failed")
		return err
	}
	if len(req.Courses) > 0 {
		p.db.ProductCourses.Delete().Where(productcourses.ProductIDEQ(req.ID)).Exec(p.ctx)
		go func() {
			for _, v := range req.Courses {
				p.db.ProductCourses.Create().
					SetProductID(v.ID).
					SetType(v.Type).
					SetName(v.Name).
					SetNumber(v.Number).
					Save(p.ctx)
			}
		}()
	}
	if len(req.Lessons) > 0 {
		p.db.ProductCourses.Delete().Where(productcourses.ProductIDEQ(req.ID)).Exec(p.ctx)
		go func() {
			for _, v := range req.Lessons {
				p.db.ProductCourses.Create().
					SetProductID(v.ID).
					SetType(v.Type).
					SetName(v.Name).
					Save(p.ctx)
			}
		}()
	}
	return nil
}

func (p Product) DeleteProduct(id int64) error {
	_, err := p.db.Product.Update().Where(product2.IDEQ(id)).SetDelete(1).Save(p.ctx)
	return err
}

func (p Product) UpdateProductStatus(ID int64, status int64) error {
	_, err := p.db.Product.Update().Where(product2.IDEQ(ID)).SetStatus(status).Save(p.ctx)
	return err
}

func (p Product) ProductInfo(id int64) (resp *product.ProductInfo, err error) {
	resp = new(product.ProductInfo)

	productInter, exist := p.cache.Get("productInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := productInter.(*product.ProductInfo); ok {
			return u, nil
		}
	}
	productEnt, err := p.db.Product.Query().Where(product2.IDEQ(id)).First(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "get product failed")
		return resp, err
	}
	resp = p.entProductInfo(productEnt)
	p.cache.SetWithTTL("productInfo"+strconv.Itoa(int(resp.ID)), &resp, 1, 1*time.Hour)

	return
}

func (p Product) ProductList(req *product.ProductListReq) (resp []*product.ProductInfo, total int, err error) {
	var predicates []predicate.Product
	if req.Name != "" {
		predicates = append(predicates, product2.NameEQ(req.Name))
	}
	if len(req.Status) > 0 {
		predicates = append(predicates, product2.StatusIn(req.Status...))
	}

	if req.Type != "" {
		predicates = append(predicates, product2.TypeEQ(req.Type))
	}
	if req.SubType != "" {
		predicates = append(predicates, product2.SubTypeEQ(req.SubType))
	}
	predicates = append(predicates, product2.Delete(0))
	lists, err := p.db.Product.Query().Where(predicates...).
		Order(ent.Desc(product2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Product list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, p.entProductInfo(v))
	}

	total, _ = p.db.Product.Query().Where(predicates...).Count(p.ctx)
	return
}

func (p Product) entProductInfo(v *ent.Product) *product.ProductInfo {
	var tags, contracts []*base.List
	tagAll, _ := v.QueryTags().All(p.ctx)
	if tagAll != nil {
		for _, item := range tagAll {
			tag := &base.List{
				ID:   item.ID,
				Name: item.Title,
			}
			tags = append(tags, tag)
		}
	}

	contractsAll, _ := v.QueryContracts().All(p.ctx)
	if contractsAll != nil {
		for _, item := range contractsAll {
			contract := &base.List{
				ID:   item.ID,
				Name: item.Name,
			}
			contracts = append(contracts, contract)
		}
	}

	var created string

	first, _ := p.db.User.Query().Where(user.IDEQ(v.CreatedID)).First(p.ctx)
	if first != nil {
		created = first.Name
	}

	var courses, lessons []*base.CourseList
	coursesAll, _ := v.QueryCourses().All(p.ctx)
	if len(coursesAll) > 0 {
		for _, ve := range coursesAll {
			courses = append(courses, &base.CourseList{
				ID:     ve.CoursesID,
				Name:   ve.Name,
				Type:   ve.Type,
				Number: ve.Number,
			},
			)

		}
	}
	lessonsAll, _ := v.QueryLessons().All(p.ctx)
	if len(lessonsAll) > 0 {
		for _, ve := range lessonsAll {
			lessons = append(lessons, &base.CourseList{
				ID:     ve.CoursesID,
				Name:   ve.Name,
				Type:   ve.Type,
				Number: ve.Number,
			},
			)

		}
	}

	return &product.ProductInfo{
		ID:          v.ID,
		Name:        v.Name,
		Pic:         v.Pic,
		Description: v.Description,
		Stock:       v.Stock,
		Status:      v.Status,
		Duration:    v.Duration,
		Length:      v.Length,
		Type:        v.Type,
		Deadline:    v.Deadline,
		Sales:       v.Sales,
		IsSales:     v.IsSales,
		SignSalesAt: v.SignSalesAt.Format(time.DateTime),
		EndSalesAt:  v.EndSalesAt.Format(time.DateTime),
		CreatedAt:   v.CreatedAt.Format(time.DateTime),
		UpdatedAt:   v.UpdatedAt.Format(time.DateTime),
		CreatedId:   v.CreatedID,
		CreatedName: created,
		Tags:        tags,
		Contracts:   contracts,

		IsLessons: v.IsLessons,
		SubType:   v.SubType,
		Price:     v.Price,
		Times:     v.Times,

		IsCourse: v.IsCourse,
		Courses:  courses,
		Lessons:  lessons,
	}
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
