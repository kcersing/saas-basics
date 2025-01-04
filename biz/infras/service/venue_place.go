package service

import (
	"context"
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/venueplace"
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/venue"

	"saas/biz/dal/db/ent"
	"time"
)

func (v Venue) CreatePlace(req *venue.VenuePlaceInfo) error {
	_, err := v.db.VenuePlace.Create().
		SetName(req.Name).
		SetPic(req.Pic).
		SetVenueID(req.VenueId).
		SetStatus(req.Status).
		SetNumber(req.Number).
		SetInformation(req.Information).
		SetClassify(req.Classify).
		SetIsShow(req.IsShow).
		SetIsAccessible(req.IsAccessible).
		SetIsBooking(req.IsBooking).
		SetType(req.Type).
		SetSeat(req.Seat).
		AddProductIDs(req.ProductIds...).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue failed")
		return err
	}

	return nil
}

func (v Venue) UpdatePlace(req *venue.VenuePlaceInfo) error {
	_, err := v.db.VenuePlace.Update().
		Where(venueplace.IDEQ(req.ID)).
		SetName(req.Name).
		SetPic(req.Pic).
		SetVenueID(req.VenueId).
		SetStatus(req.Status).
		SetNumber(req.Number).
		SetInformation(req.Information).
		SetClassify(req.Classify).
		SetIsShow(req.IsShow).
		SetIsAccessible(req.IsAccessible).
		SetIsBooking(req.IsBooking).
		SetType(req.Type).
		SetSeat(req.Seat).
		AddProductIDs(req.ProductIds...).
		Save(v.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Venue Place failed")
		return err
	}

	return nil
}
func (v Venue) PlaceDel(id int64) (err error) {
	_, err = v.db.VenuePlace.Update().Where(venueplace.IDEQ(id)).SetDelete(1).Save(v.ctx)
	return err
}

func (v Venue) PlaceList(req *venue.VenuePlaceListReq) (list []*venue.VenuePlaceInfo, total int, err error) {
	var predicates []predicate.VenuePlace

	if req.Name != "" {
		predicates = append(predicates, venueplace.NameEQ(req.Name))
	}
	if req.Status > 0 {
		predicates = append(predicates, venueplace.StatusEQ(req.Status))
	}
	if req.Type > 0 {
		predicates = append(predicates, venueplace.TypeEQ(req.Type))
	}
	predicates = append(predicates, venueplace.Delete(0))
	l, err := v.db.VenuePlace.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(v.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Venue Place list failed")
		return list, total, err
	}

	for _, p := range l {
		lt := entVenuePlaceInfo(p)
		list = append(list, lt)
	}

	total, _ = v.db.VenuePlace.Query().Where(predicates...).Count(v.ctx)
	return
}

func entVenuePlaceInfo(v *ent.VenuePlace) *venue.VenuePlaceInfo {

	createdAt := v.CreatedAt.Format(time.DateTime)
	updatedAt := v.UpdatedAt.Format(time.DateTime)
	var products []*base.List
	var productIds []int64
	productAll, _ := v.QueryProducts().All(context.Background())
	if len(productAll) > 0 {
		for _, p := range productAll {
			products = append(products, &base.List{
				ID:   p.ID,
				Name: p.Name,
			})
			productIds = append(productIds, p.ID)
		}
	}

	return &venue.VenuePlaceInfo{
		ID:           v.ID,
		Name:         v.Name,
		Pic:          v.Pic,
		VenueId:      v.VenueID,
		Number:       v.Number,
		Status:       v.Status,
		Information:  v.Information,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		Classify:     v.Classify,
		IsShow:       v.IsShow,
		IsAccessible: v.IsAccessible,
		Seat:         v.Seat,
		Products:     products,
		ProductIds:   productIds,
		IsBooking:    v.IsBooking,
		Type:         v.Type,
	}
}

func (v Venue) UpdatePlaceStatus(id int64, status int64) error {
	_, err := v.db.VenuePlace.Update().Where(venueplace.IDEQ(id)).SetStatus(status).Save(v.ctx)
	return err
}
