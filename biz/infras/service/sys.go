package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/contract"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/product"
	user2 "saas/biz/dal/db/ent/user"
	venue2 "saas/biz/dal/db/ent/venue"
	"saas/biz/dal/db/ent/venueplace"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/base"
)

type Sys struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (s Sys) RoleList(req base.ListReq) (list []do.SysList, total int64, err error) {
	var predicates []predicate.Role

	lists, err := s.db.Role.Query().Where(predicates...).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Role list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Remark,
		})
	}
	total = int64(len(list))

	return
}

func (s Sys) PlaceList(req base.ListReq) (list []do.SysList, total int64, err error) {
	var predicates []predicate.VenuePlace

	if req.VenueId > 0 {
		predicates = append(predicates, venueplace.VenueID(req.VenueId))
	}

	lists, err := s.db.VenuePlace.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get VenuePlace list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))

	return
}

func (s Sys) ProductList(req base.ListReq) (list []do.SysList, total int64, err error) {
	var predicates []predicate.Product

	if req.Name != "" {
		predicates = append(predicates, product.NameEQ(req.Name))
	}
	if req.Type != "" {
		predicates = append(predicates, product.TypeEQ(req.Type))
	}
	if req.SubType != "" {
		predicates = append(predicates, product.SubTypeEQ(req.SubType))
	}
	lists, err := s.db.Product.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))
	return
}

func (s Sys) VenueList(req base.ListReq) (list []do.SysList, total int64, err error) {
	var predicates []predicate.Venue

	if req.Name != "" {
		predicates = append(predicates, venue2.Name(req.Name))
	}
	if req.Type != "" {
		predicates = append(predicates, venue2.Type(req.Type))
	}
	lists, err := s.db.Venue.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))

	return
}

func (s Sys) MemberList(req base.ListReq) (list []do.SysList, total int64, err error) {
	var predicates []predicate.Member

	if req.Name != "" {
		predicates = append(predicates, member.Name(req.Name))
	}
	if req.Mobile != "" {
		predicates = append(predicates, member.Mobile(req.Mobile))
	}

	//if req.VenueId > 0 {
	//	predicates = append(predicates, member.VenueID(req.VenueId))
	//}
	lists, err := s.db.Debug().Member.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
			Key:  v.Mobile,
		})
	}
	total = int64(len(list))
	return
}

func (s Sys) ContractList(req base.ListReq) (list []do.SysList, total int64, err error) {
	var predicates []predicate.Contract

	if req.Name != "" {
		predicates = append(predicates, contract.Name(req.Name))
	}

	predicates = append(predicates, contract.VenueID(req.VenueId))

	lists, err := s.db.Contract.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))
	return
}

func (s Sys) StaffList(req base.ListReq) (list []do.SysList, total int64, err error) {
	var predicates []predicate.User

	if req.Name != "" {
		predicates = append(predicates, user2.Name(req.Name))
	}
	predicates = append(predicates, user2.VenueID(req.VenueId))
	lists, err := s.db.User.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))
	return
}

func NewSys(ctx context.Context, c *app.RequestContext) do.Sys {
	return &Sys{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
