package service

//
//import (
//	"context"
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/dgraph-io/ristretto"
//	"github.com/pkg/errors"
//	"saas/app/admin/config"
//	"saas/app/admin/infras"
//	"saas/app/pkg/do"
//	"saas/biz/dal/db/ent"
//	"saas/biz/dal/db/entcontract"
//	"saas/biz/dal/db/entdictionary"
//	"saas/biz/dal/db/entdictionarydetail"
//	"saas/biz/dal/db/entmember"
//	"saas/biz/dal/db/entpredicate"
//	"saas/biz/dal/db/entproduct"
//	"saas/biz/dal/db/entproductproperty"
//	"saas/biz/dal/db/entuser"
//	venue2 "saas/biz/dal/db/entvenue"
//	"saas/biz/dal/db/entvenueplace"
//	"strconv"
//)
//
//type Sys struct {
//	ctx   context.Context
//	c     *app.RequestContext
//	salt  string
//	db    *ent.Client
//	cache *ristretto.Cache
//}
//
//func (s Sys) RoleList(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.Role
//
//	lists, err := s.db.Role.Query().Where(predicates...).All(s.ctx)
//	if err != nil {
//		err = errors.Wrap(err, "get Role list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Remark,
//		})
//	}
//	total = int64(len(list))
//
//	return
//}
//
//func (s Sys) PlaceList(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.VenuePlace
//
//	if req.Venue > 0 {
//		predicates = append(predicates, venueplace.VenueID(req.Venue))
//	}
//
//	lists, err := s.db.VenuePlace.Query().Where(predicates...).All(s.ctx)
//
//	if err != nil {
//		err = errors.Wrap(err, "get VenuePlace list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Name,
//		})
//	}
//	total = int64(len(list))
//
//	return
//}
//
//func (s Sys) ProductList(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.Product
//
//	if req.Name != "" {
//		predicates = append(predicates, product.NameEQ(req.Name))
//	}
//	lists, err := s.db.Product.Query().Where(predicates...).All(s.ctx)
//
//	if err != nil {
//		err = errors.Wrap(err, "get product list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Name,
//		})
//	}
//	total = int64(len(list))
//	return
//}
//
//func (s Sys) PropertyList(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.ProductProperty
//
//	if req.Name != "" {
//		predicates = append(predicates, productproperty.NameEQ(req.Name))
//	}
//	if req.Type != "" {
//		predicates = append(predicates, productproperty.TypeEQ(req.Type))
//	}
//	if req.Product > 0 {
//		predicates = append(predicates, productproperty.HasProductWith(product.IDEQ(req.Product)))
//	}
//	if req.Venue > 0 {
//		predicates = append(predicates, productproperty.HasVenuesWith(venue2.IDEQ(req.Venue)))
//	}
//
//	lists, err := s.db.ProductProperty.Query().Where(predicates...).All(s.ctx)
//
//	if err != nil {
//		err = errors.Wrap(err, "get product list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		price := strconv.FormatFloat(v.Price, 'f', -1, 64)
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Name,
//			Key:  price,
//		})
//	}
//	total = int64(len(list))
//	return
//}
//
//func (s Sys) PropertyType(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.DictionaryDetail
//
//	if req.Name != "" {
//		predicates = append(predicates, dictionarydetail.ValueEQ(req.Name))
//	}
//	if req.DictionaryId != 0 {
//		predicates = append(predicates, dictionarydetail.HasDictionaryWith(dictionary.IDEQ(req.DictionaryId)))
//	}
//
//	lists, err := s.db.DictionaryDetail.Query().Where(predicates...).All(s.ctx)
//
//	if err != nil {
//		err = errors.Wrap(err, "get product list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Value,
//			Key:  v.Key,
//		})
//	}
//	total = int64(len(list))
//	return
//}
//
//func (s Sys) VenueList(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.Venue
//
//	if req.Name != "" {
//		predicates = append(predicates, venue2.Name(req.Name))
//	}
//
//	lists, err := s.db.Venue.Query().Where(predicates...).All(s.ctx)
//
//	if err != nil {
//		err = errors.Wrap(err, "get product list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Name,
//		})
//	}
//	total = int64(len(list))
//
//	return
//}
//
//func (s Sys) MemberList(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.Member
//
//	if req.Name != "" {
//		predicates = append(predicates, member.Name(req.Name))
//	}
//	if req.Mobile != "" {
//		predicates = append(predicates, member.Mobile(req.Mobile))
//	}
//
//	if req.Product > 0 {
//		ints, err := s.db.Debug().MemberProductProperty.Query().Where().GroupBy(member.FieldID).Ints(s.ctx)
//		if err != nil {
//			return nil, 0, err
//		}
//
//		var mids []int64
//		for _, v := range ints {
//			mids = append(mids, int64(v))
//		}
//		predicates = append(predicates, member.IDIn(mids...))
//	}
//	lists, err := s.db.Debug().Member.Query().Where(predicates...).All(s.ctx)
//
//	if err != nil {
//		err = errors.Wrap(err, "get product list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Name,
//			Key:  v.Mobile,
//		})
//	}
//	total = int64(len(list))
//	return
//}
//
//func (s Sys) ContractList(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.Contract
//
//	if req.Name != "" {
//		predicates = append(predicates, contract.Name(req.Name))
//	}
//
//	lists, err := s.db.Contract.Query().Where(predicates...).All(s.ctx)
//
//	if err != nil {
//		err = errors.Wrap(err, "get product list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Name,
//		})
//	}
//	total = int64(len(list))
//	return
//}
//
//func (s Sys) StaffList(req do.SysListReq) (list []do.SysList, total int64, err error) {
//	var predicates []predicate.User
//
//	if req.Name != "" {
//		predicates = append(predicates, user.Nickname(req.Name))
//	}
//
//	lists, err := s.db.User.Query().Where(predicates...).All(s.ctx)
//
//	if err != nil {
//		err = errors.Wrap(err, "get product list failed")
//		return nil, 0, err
//	}
//	for _, v := range lists {
//		list = append(list, do.SysList{
//			ID:   v.ID,
//			Name: v.Nickname,
//		})
//	}
//	total = int64(len(list))
//	return
//}
//
//func NewSys(ctx context.Context, c *app.RequestContext) do.Sys {
//	return &Sys{
//		ctx:   ctx,
//		c:     c,
//		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
//		db:    infras.DB,
//		cache: infras.Cache,
//	}
//}
