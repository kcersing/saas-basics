package do

import "rpc_gen/kitex_gen/company/venue"

type Venue interface {
	CreateVenue(req *venue.VenueInfo) error
	UpdateVenue(req *venue.VenueInfo) error
	VenueList(req *venue.VenueListReq) (list []*venue.VenueInfo, total int, err error)
	UpdateVenueStatus(id int64, status int64) error

	CreateVenuePlace(req *venue.VenuePlaceInfo) error
	UpdateVenuePlace(req *venue.VenuePlaceInfo) error
	VenuePlaceList(req *venue.VenuePlaceListReq) (list []*venue.VenuePlaceInfo, total int, err error)
	UpdateVenuePlaceStatus(id int64, status int64) error

	VenueInfo(id int64) (info *venue.VenueInfo, err error)
	VenuePlaceInfo(id int64) (info *venue.VenuePlaceInfo, err error)
}
