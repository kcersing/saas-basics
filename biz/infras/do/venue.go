package do

import "saas/idl_gen/model/venue"

type Venue interface {
	Create(req *venue.VenueInfo) error
	Update(req *venue.VenueInfo) error
	List(req *venue.VenueListReq) (list []*venue.VenueInfo, total int, err error)
	UpdateVenueStatus(id int64, status int64) error

	CreatePlace(req *venue.VenuePlaceInfo) error
	UpdatePlace(req *venue.VenuePlaceInfo) error
	PlaceList(req *venue.VenuePlaceListReq) (list []*venue.VenuePlaceInfo, total int, err error)
	UpdatePlaceStatus(id int64, status int64) error

	VenueInfo(id int64) (info *venue.VenueInfo, err error)
	PlaceInfo(id int64) (info *venue.VenuePlaceInfo, err error)
}
