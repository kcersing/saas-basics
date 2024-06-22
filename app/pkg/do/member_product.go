package do

type MemberProduct interface {
	ProductList(req MemberProductListReq) (resp []*MemberProductInfo, total int, err error)
	ProductDetail(ID int64) (roleInfo *MemberProductInfo, err error)
	ProductUpdate(req MemberProductInfo) error
	ProductUpdateStatus(ID int64, status int64) error

	PropertyList(req MemberPropertyListReq) (resp []*MemberPropertyInfo, total int, err error)
	PropertyDetail(ID int64) (roleInfo *MemberPropertyInfo, err error)
	PropertyUpdate(req MemberPropertyInfo) error
	PropertyUpdateStatus(ID int64, status int64) error

	EntryList(req MemberEntryListReq) (resp []*MemberEntryInfo, total int, err error)
}
type MemberProductInfo struct {
	ID          int64          `json:"id"`
	MemberId    int64          `json:"member_id"`
	Name        string         `json:"name"`
	Pic         string         `json:"pic"`
	Description string         `json:"description"`
	Property    []PropertyInfo `json:"property"`
	Price       float64        `json:"price"`
	Stock       int64          `json:"stock"`
	Status      int64          `json:"status"`
}
type MemberPropertyListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	Status   int64  `json:"status"`
}

type MemberEntryListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	Status   int64  `json:"status"`
}

type MemberProductListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	Status   int64  `json:"status"`
}

type MemberPropertyInfo struct {
	ID        int64   `json:"id"`
	MemberId  int64   `json:"member_id"`
	ProductId int64   `json:"productId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Duration  int64   `json:"duration"`
	Length    int64   `json:"length"`
	Count     int64   `json:"count"`
	Type      string  `json:"type"`
	Data      string  `json:"data"`
	Status    int64   `json:"status"`
}

type MemberEntryInfo struct {
	ID int64 `json:"id"`
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
