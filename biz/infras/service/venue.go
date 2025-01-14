package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/dictionarydetail"
	"saas/biz/dal/db/ent/predicate"
	venue2 "saas/biz/dal/db/ent/venue"
	"saas/biz/dal/db/ent/venueplace"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/venue"
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

func (v Venue) VenueDel(id int64) (err error) {
	_, err = v.db.Venue.Update().Where(venue2.IDEQ(id)).SetDelete(1).Save(v.ctx)
	return err
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

	info = v.entVenueInfo(one)
	//err = copier.CopyWithOption(info, one, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	//hlog.Info("err :", err)
	//if err != nil {
	//	err = errors.Wrap(err, "copy venue info failed")
	//	return info, err
	//}

	v.cache.SetWithTTL("venueInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
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

	info = entVenuePlaceInfo(one)

	v.cache.SetWithTTL("placeInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
	return
}

func (v Venue) Create(req *venue.VenueInfo) error {
	_, err := v.db.Venue.Create().
		SetName(req.Name).
		SetAddress(req.Address).
		SetAddressDetail(req.AddressDetail).
		SetLatitude(req.Latitude).
		SetLongitude(req.Longitude).
		SetMobile(req.Mobile).
		SetPic(req.Pic).
		SetEmail(req.Email).
		SetInformation(req.Information).
		SetClassify(req.Classify).
		SetType(req.Type).
		SetSeal(req.Seal).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) Update(req *venue.VenueInfo) error {
	_, err := v.db.Venue.Update().
		Where(venue2.IDEQ(req.ID)).
		SetName(req.Name).
		SetAddress(req.Address).
		SetAddressDetail(req.AddressDetail).
		SetLatitude(req.Latitude).
		SetLongitude(req.Longitude).
		SetMobile(req.Mobile).
		SetPic(req.Pic).
		SetEmail(req.Email).
		SetInformation(req.Information).
		SetStatus(req.Status).
		SetClassify(req.Classify).
		SetType(req.Type).
		SetSeal(req.Seal).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) List(req *venue.VenueListReq) (list []*venue.VenueInfo, total int, err error) {
	var predicates []predicate.Venue

	if req.Name != "" {
		predicates = append(predicates, venue2.NameEQ(req.Name))
	}
	if req.Status > 0 {
		predicates = append(predicates, venue2.StatusEQ(req.Status))
	}
	predicates = append(predicates, venue2.Delete(0))
	venueList, err := v.db.Venue.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(venue2.FieldID)).
		Limit(int(req.PageSize)).All(v.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Venue list failed")
		return list, total, err
	}

	for _, ve := range venueList {
		list = append(list, v.entVenueInfo(ve))
	}

	total, _ = v.db.Venue.Query().Where(predicates...).Count(v.ctx)
	return
}

func (v Venue) UpdateVenueStatus(id int64, status int64) error {
	_, err := v.db.Venue.Update().Where(venue2.IDEQ(id)).SetStatus(status).Save(v.ctx)
	return err
}

func (ve Venue) entVenueInfo(v *ent.Venue) *venue.VenueInfo {
	createdAt := v.CreatedAt.Format(time.DateTime)
	updatedAt := v.UpdatedAt.Format(time.DateTime)

	var dAll []*base.List
	ddAll, _ := ve.db.DictionaryDetail.Query().Where(dictionarydetail.IDIn(v.Classify...)).All(ve.ctx)
	if ddAll != nil {
		for _, item := range ddAll {
			contract := &base.List{
				ID:   item.ID,
				Name: item.Title,
			}
			dAll = append(dAll, contract)
		}
	}

	return &venue.VenueInfo{
		ID:            v.ID,
		Name:          v.Name,
		Address:       v.Address,
		Pic:           v.Pic,
		Mobile:        v.Mobile,
		Latitude:      v.Latitude,
		Longitude:     v.Longitude,
		Status:        v.Status,
		Email:         v.Email,
		AddressDetail: v.AddressDetail,
		Information:   v.Information,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		Classify:      v.Classify,
		Type:          v.Type,
		Seal:          v.Seal,
		ClassifyName:  dAll,
	}
}
