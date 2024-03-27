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
	"saas/pkg/db/ent/venue"
	"saas/pkg/db/ent/venueplace"
)

type Venue struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (v Venue) Create(req *do.VenueInfo) error {
	_, err := v.db.Venue.Create().
		SetName(req.Name).
		SetAddress(req.Address).
		SetAddressDetail(req.AddressDetail).
		SetLatitude(req.Latitude).
		SetLongitude(req.Longitude).
		SetMobile(req.Mobile).
		SetPic(req.Pic).
		SetInformation(req.Information).
		SetStatus(req.Status).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) Update(req *do.VenueInfo) error {
	_, err := v.db.Venue.Update().
		Where(venue.IDEQ(req.ID)).
		SetName(req.Name).
		SetAddress(req.Address).
		SetAddressDetail(req.AddressDetail).
		SetLatitude(req.Latitude).
		SetLongitude(req.Longitude).
		SetMobile(req.Mobile).
		SetPic(req.Pic).
		SetInformation(req.Information).
		SetStatus(req.Status).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) List(req *do.VenueListReq) (list []*do.VenueInfo, total int, err error) {
	var predicates []predicate.Venue

	if req.Name != "" {
		predicates = append(predicates, venue.NameEQ(req.Name))
	}

	venueList, err := v.db.Venue.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(v.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Venue list failed")
		return list, total, err
	}

	err = copier.Copy(&list, &venueList)
	if err != nil {
		err = errors.Wrap(err, "copy Venue info failed")
		return list, 0, err
	}
	total, _ = v.db.Venue.Query().Where(predicates...).Count(v.ctx)
	return
}

func (v Venue) UpdateVenueStatus(id int64, status int64) error {
	_, err := v.db.Venue.Update().Where(venue.IDEQ(id)).SetStatus(status).Save(v.ctx)
	return err
}
func (v Venue) CreatePlace(req *do.VenuePlaceInfo) error {
	_, err := v.db.VenuePlace.Create().
		SetName(req.Name).
		SetPic(req.Pic).
		SetVenueID(req.VenueId).
		SetStatus(req.Status).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) UpdatePlace(req *do.VenuePlaceInfo) error {
	_, err := v.db.VenuePlace.Update().
		Where(venueplace.IDEQ(req.ID)).
		SetName(req.Name).
		SetPic(req.Pic).
		SetVenueID(req.VenueId).
		SetStatus(req.Status).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue Place failed")
		return err
	}

	return nil
}

func (v Venue) PlaceList(req *do.VenuePlaceListReq) (list []*do.VenuePlaceInfo, total int, err error) {
	var predicates []predicate.VenuePlace

	if req.Name != "" {
		predicates = append(predicates, venueplace.NameEQ(req.Name))
	}

	l, err := v.db.VenuePlace.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(v.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Venue Place list failed")
		return list, total, err
	}

	err = copier.Copy(&list, &l)
	if err != nil {
		err = errors.Wrap(err, "copy Venue Place info failed")
		return list, 0, err
	}
	total, _ = v.db.VenuePlace.Query().Where(predicates...).Count(v.ctx)
	return
}

func (v Venue) UpdatePlaceStatus(id int64, status int64) error {
	_, err := v.db.VenuePlace.Update().Where(venueplace.IDEQ(id)).SetStatus(status).Save(v.ctx)
	return err
}
func NewVenue(ctx context.Context, c *app.RequestContext) do.Venue {
	return &Venue{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}