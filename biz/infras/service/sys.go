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
	"saas/biz/dal/db/ent/dictionarydetail"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/memberprofile"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/role"
	user2 "saas/biz/dal/db/ent/user"
	venue2 "saas/biz/dal/db/ent/venue"
	"saas/biz/dal/db/ent/venueplace"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/sys"
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
	predicates = append(predicates, role.Delete(0))
	lists, err := s.db.Role.Query().Where(predicates...).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Role list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			Id:   v.ID,
			Name: v.Remark,
		})
	}
	total = int64(len(list))

	return
}

func (s Sys) PlaceList(req sys.SysPlaceListReq) (list []do.SysPlaceList, total int64, err error) {
	var predicates []predicate.VenuePlace

	if req.VenueId > 0 {
		predicates = append(predicates, venueplace.VenueID(req.VenueId))
	}
	if req.Type > 0 {
		predicates = append(predicates, venueplace.TypeEQ(req.Type))
	}
	if req.Classify > 0 {
		predicates = append(predicates, venueplace.ClassifyEQ(req.Classify))
	}
	if len(req.ProductId) > 0 {
		predicates = append(predicates, venueplace.HasProductsWith(product.IDIn(req.ProductId...)))
	}

	predicates = append(predicates, venueplace.Delete(0))
	lists, err := s.db.VenuePlace.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get VenuePlace list failed")
		return nil, 0, err
	}
	for _, v := range lists {

		products, _ := v.QueryProducts().Select("product_id").Ints(s.ctx)

		list = append(list, do.SysPlaceList{
			Id:       v.ID,
			Name:     v.Name,
			Products: products,
			Seat:     v.Seat,
			Number:   v.Number,
			Classify: v.Classify,
			Type:     v.Type,
		})
	}
	total = int64(len(list))

	return
}

func (s Sys) ProductList(req sys.SysProductListReq) (list []do.SysProductList, total int64, err error) {
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
	if req.VenueId > 0 {
		predicates = append(predicates, product.VenueID(req.VenueId))
	}
	predicates = append(predicates, product.Delete(0))
	lists, err := s.db.Product.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}

	for _, v := range lists {
		tags, _ := v.QueryTags().Select("id").Ints(s.ctx)

		list = append(list, do.SysProductList{
			Id:     v.ID,
			Name:   v.Name,
			Length: v.Length,
			Tags:   tags,
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
	predicates = append(predicates, venue2.Delete(0))
	lists, err := s.db.Venue.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))

	return
}

func (s Sys) MemberList(req sys.SysMemberListReq) (list []do.SysMemberList, total int64, err error) {
	var predicates []predicate.Member

	if req.Name != "" {
		predicates = append(predicates, member.HasMemberProfileWith(memberprofile.Name(req.Name)))
	}
	if req.Condition > 0 {
		predicates = append(predicates, member.HasMemberProfileWith(memberprofile.Condition(req.Condition)))
	}
	if req.Mobile != "" {
		predicates = append(predicates, member.Mobile(req.Mobile))
	}

	if req.VenueId > 0 {
		predicates = append(predicates, member.HasMemberProfileWith(memberprofile.VenueID(req.VenueId)))
	}

	if len(req.ProductId) > 0 {
		predicates = append(predicates, member.HasMemberProductsWith(memberproduct.ProductIDIn(req.ProductId...)))
	}
	predicates = append(predicates, member.Delete(0))
	lists, err := s.db.Member.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		pro, _ := v.QueryMemberProfile().First(s.ctx)
		if pro != nil {
			list = append(list, do.SysMemberList{
				Id:     v.ID,
				Name:   pro.Name,
				Mobile: v.Mobile,
			})
		}

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
	predicates = append(predicates, contract.Delete(0))
	lists, err := s.db.Contract.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		list = append(list, do.SysList{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	total = int64(len(list))
	return
}

func (s Sys) StaffList(req sys.SysStaffListReq) (list []do.SysStaffList, total int64, err error) {
	var predicates []predicate.User

	if req.Name != "" {
		predicates = append(predicates, user2.Name(req.Name))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, user2.VenueID(req.VenueId))
	}
	if len(req.TagId) > 0 {
		predicates = append(predicates, user2.HasTagsWith(dictionarydetail.IDIn(req.TagId...)))
	}
	predicates = append(predicates, user2.Type(1))
	predicates = append(predicates, user2.Delete(0))
	lists, err := s.db.User.Query().Where(predicates...).All(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}
	for _, v := range lists {
		tags, _ := v.QueryTags().Select("id").Ints(s.ctx)

		list = append(list, do.SysStaffList{
			Id:   v.ID,
			Name: v.Name,
			Tags: tags,
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
