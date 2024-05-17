package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/contract"
	"saas/pkg/db/ent/dictionary"
	"saas/pkg/db/ent/dictionarydetail"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/product"
	"saas/pkg/db/ent/productproperty"
	"saas/pkg/db/ent/user"
	venue2 "saas/pkg/db/ent/venue"
)

type Sys struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (s Sys) ProductList(req do.SysListReq) (list []do.SysList, err error) {
	var predicates []predicate.Product

	if req.Name != "" {
		predicates = append(predicates, product.NameEQ(req.Name))
	}
	lists, err := s.db.Product.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return
}

func (s Sys) PropertyList(req do.SysListReq) (list []do.SysList, err error) {
	var predicates []predicate.ProductProperty

	if req.Name != "" {
		predicates = append(predicates, productproperty.NameEQ(req.Name))
	}
	lists, err := s.db.ProductProperty.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return
}

func (s Sys) PropertyType(req do.SysListReq) (list []do.SysList, err error) {
	var predicates []predicate.DictionaryDetail

	if req.Name != "" {
		predicates = append(predicates, dictionarydetail.ValueEQ(req.Name))
	}
	if req.DictionaryId != 0 {
		predicates = append(predicates, dictionarydetail.HasDictionaryWith(dictionary.IDEQ(req.DictionaryId)))
	}

	lists, err := s.db.DictionaryDetail.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Value,
		})
	}

	return
}

func (s Sys) VenueList(req do.SysListReq) (list []do.SysList, err error) {
	var predicates []predicate.Venue

	if req.Name != "" {
		predicates = append(predicates, venue2.Name(req.Name))
	}

	lists, err := s.db.Venue.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return
}

func (s Sys) MemberList(req do.SysListReq) (list []do.SysList, err error) {
	var predicates []predicate.Member

	if req.Name != "" {
		predicates = append(predicates, member.Name(req.Name))
	}
	if req.Mobile != "" {
		predicates = append(predicates, member.Mobile(req.Mobile))
	}

	lists, err := s.db.Member.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return
}

func (s Sys) ContractList(req do.SysListReq) (list []do.SysList, err error) {
	var predicates []predicate.Contract

	if req.Name != "" {
		predicates = append(predicates, contract.Name(req.Name))
	}

	lists, err := s.db.Contract.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return
}

func (s Sys) StaffList(req do.SysListReq) (list []do.SysList, err error) {
	var predicates []predicate.User

	if req.Name != "" {
		predicates = append(predicates, user.Nickname(req.Name))
	}

	lists, err := s.db.User.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Nickname,
		})
	}

	return
}

func NewSys(ctx context.Context, c *app.RequestContext) do.Sys {
	return &Sys{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
