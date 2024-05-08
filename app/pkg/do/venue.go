package do

type Venue interface {
	Create(req *VenueInfo) error
	Update(req *VenueInfo) error
	List(req *VenueListReq) (list []*VenueInfo, total int, err error)
	UpdateVenueStatus(id int64, status int64) error

	CreatePlace(req *VenuePlaceInfo) error
	UpdatePlace(req *VenuePlaceInfo) error
	PlaceList(req *VenuePlaceListReq) (list []*VenuePlaceInfo, total int, err error)
	UpdatePlaceStatus(id int64, status int64) error
}

type VenueListReq struct {
	PageSize int64  `json:"page_size"`
	Page     int64  `json:"page"`
	Name     string `json:"name"`
}
type VenueInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	AddressDetail string `json:"address_detail"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Mobile        string `json:"mobile"`
	Pic           string `json:"pic"`
	Information   string `json:"information"`
	Status        int64  `json:"status"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

type VenuePlaceListReq struct {
	PageSize int64  `json:"page_size"`
	Page     int64  `json:"page"`
	Name     string `json:"name"`
}
type VenuePlaceInfo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	VenueId   int64  `json:"venue_id"`
	Status    int64  `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
