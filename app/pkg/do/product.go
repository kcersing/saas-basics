package do

type Product interface {
	CreateProperty(req PropertyInfo) error
	UpdateProperty(req PropertyInfo) error
	DeleteProperty(id int64) error
	UpdatePropertyStatus(ID int64, status int64) error
	PropertyList(req ProductListReq) (resp []*PropertyInfo, total int, err error)

	Create(req ProductInfo) error
	Update(req ProductInfo) error
	Delete(id int64) error
	List(req ProductListReq) (resp []*ProductInfo, total int, err error)
	UpdateStatus(ID int64, status int64) error
	InfoByID(ID int64) (roleInfo *ProductInfo, err error)
}

type PropertyInfo struct {
	ID         int64           `json:"id"`
	Name       string          `json:"name"`
	Price      float64         `json:"price"`
	Duration   int64           `json:"duration"`
	Length     int64           `json:"length"`
	Count      int64           `json:"count"`
	Type       string          `json:"type"`
	Data       string          `json:"data"`
	Status     int64           `json:"status"`
	CreateId   int64           `json:"create_id"`
	CreateName string          `json:"create_name"`
	Venue      []PropertyVenue `json:"venue"`
	Venues     string          `json:"venues"`
	VenueId    []int64         `json:"venueId"`
}
type PropertyVenue struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type ProductInfo struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Pic         string         `json:"pic"`
	Description string         `json:"description"`
	Property    []PropertyInfo `json:"property"`
	Price       float64        `json:"price"`
	Stock       int64          `json:"stock"`
	Status      int64          `json:"status"`
	CreateID    int64          `json:"create_id"`
	CreateName  string         `json:"create_name"`
}
type ProductListReq struct {
	Page        int64    `json:"page"`
	PageSize    int64    `json:"pageSize"`
	Name        string   `json:"name"`
	Status      []int64  `json:"status"`
	CreatedTime []string `json:"createdTime"`
	Venue       []int64  `json:"venue"`
	Type        string   `json:"type"`
}
