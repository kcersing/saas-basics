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
	PageSize int64  `json:"page_size,omitempty"`
	Page     int64  `json:"page,omitempty"`
	Name     string `json:"name,omitempty"`
}
type VenueInfo struct {
	ID            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Address       string `json:"address,omitempty"`
	AddressDetail string `json:"address_detail,omitempty"`
	Latitude      string `json:"latitude,omitempty"`
	Longitude     string `json:"longitude,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
	Pic           string `json:"pic,omitempty"`
	Information   string `json:"information,omitempty"`
	Status        int64  `json:"status,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
}

type VenuePlaceListReq struct {
	PageSize int64  `json:"page_size,omitempty"`
	Page     int64  `json:"page,omitempty"`
	Name     string `json:"name,omitempty"`
}
type VenuePlaceInfo struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Pic       string `json:"pic,omitempty"`
	VenueId   int64  `json:"venue_id,omitempty"`
	Status    int64  `json:"status,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
