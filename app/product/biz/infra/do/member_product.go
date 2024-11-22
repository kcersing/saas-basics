package do

type MemberProduct interface {
	ProductList(req MemberProductListReq) (resp []*MemberProductInfo, total int, err error)
	ProductDetail(ID int64) (info *MemberProductInfo, err error)
	ProductUpdate(req MemberProductInfo) error
	ProductUpdateStatus(ID int64, status int64) error

	PropertyList(req MemberPropertyListReq) (resp []*MemberPropertyInfo, total int, err error)
	PropertyDetail(ID int64) (info *MemberPropertyInfo, err error)
	PropertyUpdate(req MemberPropertyInfo) error
	PropertyUpdateStatus(ID int64, status int64) error

	ContractList(req MemberContractListReq) (resp []*MemberContractInfo, total int, err error)
}
type MemberContractListReq struct {
	MemberId   int64 `json:"member_id"`
	VenueId    int64 `json:"venue_id"`
	ContractId int64 `json:"contract_id"`
	Page       int64 `json:"page"`
	PageSize   int64 `json:"pageSize"`
}
type MemberContractInfo struct {
	Name              string `json:"name"`
	MemberId          int64  `json:"member_id"`
	MemberName        int64  `json:"member_name"`
	VenueId           int64  `json:"venue_id"`
	VenueName         string `json:"venue_name"`
	MemberProductId   int64  `json:"member_product_id"`
	MemberProductName string `json:"member_product_name"`
	ContractId        int64  `json:"contract_id"`
	ContractName      int64  `json:"contract_name"`
	Sign              string `json:"sign"`
	SignImg           string `json:"sign_img"`
	Content           string `json:"content"`
}

type MemberProductInfo struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Status     int64   `json:"status"`
	StatusName string  `json:"status_name"`
	CreatedAt  string  `json:"created_at"`
	Sn         string  `json:"sn"`
	MemberId   int64   `json:"member_id"`
	OrderId    int64   `json:"order_id"`
	VenueId    int64   `json:"venue_id"`

	VenueName string `json:"venue_name"`
	ProductId int64  `json:"product_id"`

	Property []MemberPropertyInfo `json:"property"`
}

type MemberProductListReq struct {
	Name     string `json:"name"`
	MemberId int64  `json:"member_id"`
	VenueId  int64  `json:"venue_id"`
	Status   int64  `json:"status"`

	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type MemberPropertyListReq struct {
	Name            string `json:"name"`
	MemberId        int64  `json:"member_id"`
	MemberProductId int64  `json:"member_product_id"`
	Type            string `json:"type"`
	Status          int64  `json:"status"`
	VenueId         int64  `json:"venue_id"`

	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type MemberPropertyInfo struct {
	ID              int64           `json:"id"`
	Sn              string          `json:"sn"`
	MemberId        int64           `json:"member_id"`
	MemberProductId int64           `json:"member_product_id"`
	PropertyId      int64           `json:"property_id"`
	Name            string          `json:"name"`
	Price           float64         `json:"price"`
	Duration        int64           `json:"duration"`
	Length          int64           `json:"length"`
	Count           int64           `json:"count"`
	CountSurplus    int64           `json:"count_surplus"`
	Type            string          `json:"type"`
	Status          int64           `json:"status"`
	CreatedAt       string          `json:"created_at"`
	ValidityAt      string          `json:"validity_at"`
	CancelAt        string          `json:"cancel_at"`
	Venue           []PropertyVenue `json:"venue"`
	VenueId         []int64         `json:"venueId"`
	Venues          string          `json:"venues"`
}

type MemberEntryInfo struct {
	ID               int64  `json:"id"`
	CreatedAt        string `json:"created_at"`
	MemberPropertyId int64  `json:"member_property_id"`
	EntryTime        string `json:"entry_time"`
	LeavingTime      string `json:"leaving_time"`
	MemberId         int64  `json:"member_id"`
	MemberProductId  int64  `json:"member_product_id"`
	UserId           int64  `json:"user_id"`
	VenueId          int64  `json:"venue_id"`
}

// MPStatus 会员产品状态
type MPStatus int

const (
	MPStatusUnfinished = iota
	MPStatusNotActivated
	MPStatusActivated
	MPStatusExpire
	MPStatusExhaust
	MPStatusUpgrade
	MPStatusFreeze
)

var MPStatusNames = map[MPStatus]string{
	MPStatusUnfinished:   "未完成",
	MPStatusNotActivated: "未激活",
	MPStatusActivated:    "已激活",
	MPStatusExpire:       "已到期",
	MPStatusExhaust:      "已完结",
	MPStatusUpgrade:      "已升级",
	MPStatusFreeze:       "已冻结",
}

func (s MPStatus) String() string {
	return MPStatusNames[s]
}
