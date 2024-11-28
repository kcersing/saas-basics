package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/venue"

	"saas/biz/dal/cache"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/predicate"
	venue2 "saas/pkg/db/ent/venue"
	"saas/pkg/db/ent/venueplace"
	"strconv"
	"time"
)

type Venue struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewVenue(ctx context.Context, c *app.RequestContext) do.Venue {
	return &Venue{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (v Venue) VenueInfo(id int64) (info *venue.VenueInfo, err error) {
	inter, exist := v.cache.Get("venueInfo" + strconv.Itoa(int(id)))
	if exist {
		if v, ok := inter.(*venue.VenueInfo); ok {
			return v, nil
		}
	}
	one, err := v.db.Venue.Query().Where(venue2.IDEQ(id)).First(v.ctx)
	if err != nil {
		err = errors.Wrap(err, "get venue failed")
		return info, err
	}
	createdAt := one.CreatedAt.Format(time.DateTime)
	updatedAt := one.UpdatedAt.Format(time.DateTime)
	info = &venue.VenueInfo{
		ID:        &one.ID,
		Name:      &one.Name,
		Address:   &one.Address,
		Pic:       &one.Pic,
		Mobile:    &one.Mobile,
		Latitude:  &one.Latitude,
		Longitude: &one.Longitude,
		Status:    &one.Status,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}
	//err = copier.CopyWithOption(info, one, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	//hlog.Info("err :", err)
	//if err != nil {
	//	err = errors.Wrap(err, "copy venue info failed")
	//	return info, err
	//}

	v.cache.SetWithTTL("venueInfo"+strconv.Itoa(int(*info.ID)), &info, 1, 1*time.Hour)
	return
}

func (v Venue) PlaceInfo(id int64) (info *venue.VenuePlaceInfo, err error) {
	inter, exist := v.cache.Get("placeInfo" + strconv.Itoa(int(id)))
	if exist {
		if v, ok := inter.(*venue.VenuePlaceInfo); ok {
			return v, nil
		}
	}
	one, err := v.db.VenuePlace.Query().Where(venueplace.IDEQ(id)).First(v.ctx)
	if err != nil {
		err = errors.Wrap(err, "get venue place failed")
		return info, err
	}

	err = copier.Copy(&info, &one)
	if err != nil {
		err = errors.Wrap(err, "copy venue place info failed")
		return info, err
	}

	v.cache.SetWithTTL("placeInfo"+strconv.Itoa(int(*info.ID)), &info, 1, 1*time.Hour)
	return
}

func (v Venue) Create(req *venue.VenueInfo) error {
	_, err := v.db.Venue.Create().
		SetName(*req.Name).
		SetAddress(*req.Address).
		SetAddressDetail(*req.AddressDetail).
		SetLatitude(*req.Latitude).
		SetLongitude(*req.Longitude).
		SetMobile(*req.Mobile).
		SetPic(*req.Pic).
		SetEmail(*req.Email).
		SetInformation(*req.Information).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) Update(req *venue.VenueInfo) error {
	_, err := v.db.Venue.Update().
		Where(venue2.IDEQ(*req.ID)).
		SetName(*req.Name).
		SetAddress(*req.Address).
		SetAddressDetail(*req.AddressDetail).
		SetLatitude(*req.Latitude).
		SetLongitude(*req.Longitude).
		SetMobile(*req.Mobile).
		SetPic(*req.Pic).
		SetEmail(*req.Email).
		SetInformation(*req.Information).
		SetStatus(*req.Status).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) List(req *venue.VenueListReq) (list []*venue.VenueInfo, total int, err error) {
	var predicates []predicate.Venue

	if *req.Name != "" {
		predicates = append(predicates, venue2.NameEQ(*req.Name))
	}

	venueList, err := v.db.Venue.Query().Where(predicates...).
		Offset(int(*req.Page-1) * int(*req.PageSize)).
		Limit(int(*req.PageSize)).All(v.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Venue list failed")
		return list, total, err
	}

	err = copier.Copy(&list, &venueList)
	if err != nil {
		err = errors.Wrap(err, "copy Venue info failed")
		return list, 0, err
	}
	for i, v := range venueList {
		createdAt := v.CreatedAt.Format(time.DateTime)
		list[i].CreatedAt = &createdAt
	}

	total, _ = v.db.Venue.Query().Where(predicates...).Count(v.ctx)
	return
}

func (v Venue) UpdateVenueStatus(id int64, status int64) error {
	_, err := v.db.Venue.Update().Where(venue2.IDEQ(id)).SetStatus(status).Save(v.ctx)
	return err
}
func (v Venue) CreatePlace(req *venue.VenuePlaceInfo) error {
	_, err := v.db.VenuePlace.Create().
		SetName(*req.Name).
		SetPic(*req.Pic).
		SetVenueID(*req.VenueId).
		SetStatus(*req.Status).
		SetNumber(*req.Number).
		SetInformation(*req.Information).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) UpdatePlace(req *venue.VenuePlaceInfo) error {
	_, err := v.db.VenuePlace.Update().
		Where(venueplace.IDEQ(*req.ID)).
		SetName(*req.Name).
		SetPic(*req.Pic).
		SetVenueID(*req.VenueId).
		SetStatus(*req.Status).
		SetNumber(*req.Number).
		SetInformation(*req.Information).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue Place failed")
		return err
	}

	return nil
}

func (v Venue) PlaceList(req *venue.VenuePlaceListReq) (list []*venue.VenuePlaceInfo, total int, err error) {
	var predicates []predicate.VenuePlace

	if *req.Name != "" {
		predicates = append(predicates, venueplace.NameEQ(*req.Name))
	}

	l, err := v.db.VenuePlace.Query().Where(predicates...).
		Offset(int(*req.Page-1) * int(*req.PageSize)).
		Limit(int(*req.PageSize)).All(v.ctx)
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
