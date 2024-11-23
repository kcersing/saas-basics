package service

import (
	"context"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/membercontract"
	"saas/pkg/db/ent/memberproduct"
	"saas/pkg/db/ent/memberproductproperty"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/venue"
)

type MemberProduct struct {
	ctx   context.Context
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m MemberProduct) ContractList(req do.MemberContractListReq) (resp []*do.MemberContractInfo, total int, err error) {
	var predicates []predicate.MemberContract
	if req.MemberId > 0 {
		predicates = append(predicates, membercontract.MemberID(req.MemberId))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, membercontract.VenueID(req.VenueId))
	}
	if req.ContractId > 0 {
		predicates = append(predicates, membercontract.ContractID(req.ContractId))
	}

	lists, err := m.db.MemberContract.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Contract list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Member Contract info failed")
		return resp, 0, err
	}
	for i, v := range lists {
		vInfo, err := NewVenue(m.ctx, m.c).VenueInfo(v.VenueID)
		if err == nil {
			resp[i].VenueName = vInfo.Name
		}
		first, err := v.QueryContent().First(m.ctx)
		if err != nil {
			return nil, 0, err
		}
		resp[i].Content = first.Content
		resp[i].SignImg = first.SignImg
	}

	return
}

func (m MemberProduct) ProductList(req do.MemberProductListReq) (resp []*do.MemberProductInfo, total int, err error) {
	var predicates []predicate.MemberProduct
	if req.MemberId > 0 {
		predicates = append(predicates, memberproduct.MemberID(req.MemberId))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, memberproduct.VenueID(req.VenueId))
	}
	if req.Name != "" {
		predicates = append(predicates, memberproduct.Name(req.Name))
	}
	if req.Status > 0 {
		predicates = append(predicates, memberproduct.Status(req.Status))
	}

	lists, err := m.db.MemberProduct.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Product list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Member Product info failed")
		return resp, 0, err
	}
	for i, v := range lists {
		vInfo, err := NewVenue(m.ctx, m.c).VenueInfo(v.VenueID)
		if err == nil {
			resp[i].VenueName = vInfo.Name
		}
		resp[i].StatusName = do.MPStatusNames[do.MPStatus(v.Status)]
		list, _, err := m.PropertyList(do.MemberPropertyListReq{
			MemberProductId: v.ID,
			Page:            1,
			PageSize:        99,
		})
		if err == nil {
			copier.Copy(&resp[i].Property, &list)
		}

	}

	total, _ = m.db.MemberProduct.Query().Where(predicates...).Count(m.ctx)
	return
}

func (m MemberProduct) ProductDetail(ID int64) (info *do.MemberProductInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) ProductUpdate(req do.MemberProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) ProductUpdateStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) PropertyList(req do.MemberPropertyListReq) (resp []*do.MemberPropertyInfo, total int, err error) {
	var predicates []predicate.MemberProductProperty
	if req.MemberId > 0 {
		predicates = append(predicates, memberproductproperty.MemberID(req.MemberId))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, memberproductproperty.HasVenuesWith(venue.ID(req.VenueId)))
	}
	if req.MemberProductId > 0 {
		predicates = append(predicates, memberproductproperty.MemberProductID(req.MemberProductId))
	}
	if req.Type != "" {
		predicates = append(predicates, memberproductproperty.Type(req.Type))
	}
	if req.Name != "" {
		predicates = append(predicates, memberproductproperty.Name(req.Name))
	}
	if req.Status > 0 {
		predicates = append(predicates, memberproductproperty.Status(req.Status))
	}

	lists, err := m.db.MemberProductProperty.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Product list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Member Product info failed")
		return resp, 0, err
	}

	for i, v := range lists {
		venues, err := v.QueryVenues().Select(venue.FieldID, venue.FieldName).All(m.ctx)
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

	total, _ = m.db.MemberProductProperty.Query().Where(predicates...).Count(m.ctx)
	return
}

func (m MemberProduct) PropertyDetail(ID int64) (info *do.MemberPropertyInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) PropertyUpdate(req do.MemberPropertyInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) PropertyUpdateStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) Create(req do.MemberProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) Update(req do.MemberProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) List(req do.MemberProductListReq) (resp []*do.MemberProductInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) UpdateStatus(ID int64, status int64) error {
	_, err := m.db.MemberProduct.Update().Where(memberproduct.IDEQ(ID)).SetStatus(int64(status)).Save(m.ctx)
	return err
}

func (m MemberProduct) InfoByID(ID int64) (roleInfo *do.MemberProductInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func NewMemberProduct(ctx context.Context) do.MemberProduct {
	return &MemberProduct{
		ctx:   ctx,
		salt:  "",
		db:    mysql.DB,
		cache: cache.Cache,
	}
}
