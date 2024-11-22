package do

type Product interface {
	CreateProperty(req PropertyInfo) error
	UpdateProperty(req PropertyInfo) error
	DeleteProperty(id int64) error
	UpdatePropertyStatus(ID int64, status int64) error
	PropertyList(req ProductListReq) (resp []*PropertyInfo, total int, err error)

	Create(req CreateOrUpdateProduct) error
	Update(req ProductInfo) error
	Delete(id int64) error
	List(req ProductListReq) (resp []*ProductInfo, total int, err error)
	UpdateStatus(ID int64, status int64) error
	Info(id int64) (info *ProductInfo, err error)
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
	VenueId    []int64         `json:"venueId"`
	Venues     string          `json:"venues"`
}
type PropertyVenue struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateOrUpdateProduct struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Pic            string  `json:"pic"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	Stock          int64   `json:"stock"`
	CreateID       int64   `json:"create_id"`
	CardProperty   int64   `json:"cardProperty"`
	ClassProperty  []int64 `json:"classProperty"`
	CourseProperty []int64 `json:"courseProperty"`
}

type ProductInfo struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Pic         string  `json:"pic"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Status      int64   `json:"status"`
	CreateID    int64   `json:"create_id"`
	CreateName  string  `json:"create_name"`

	CardProperty   []PropertyInfo `json:"cardProperty"`
	ClassProperty  []PropertyInfo `json:"classProperty"`
	CourseProperty []PropertyInfo `json:"courseProperty"`
	Property       []PropertyInfo `json:"property"`
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
