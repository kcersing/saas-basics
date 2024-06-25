package do

import (
	"time"
)

type Order interface {
	Create(req CreateOrder) (string, error)
	Update(req OrderInfo) error
	List(req OrderListReq) (resp []*OrderInfo, total int, err error)
	UpdateStatus(id int64, status int64) error
	Info(id int64) (info *OrderInfo, err error)
	GetBySnOrder(sn string) (info *OrderInfo, err error)

	UnifyPay(req UnifyPayReq) error
	QRPay(req QRPayReq) (string, error)
}

type CreateOrder struct {
	AssignAt       string         `json:"assign_at,omitempty"`
	CardProperty   []PropertyItem `json:"cardProperty,omitempty"`
	CourseProperty []PropertyItem `json:"courseProperty,omitempty"`
	ClassProperty  []PropertyItem `json:"classProperty,omitempty"`
	Venue          int64          `json:"venue,omitempty"`
	Member         int64          `json:"member,omitempty"`
	Product        int64          `json:"product,omitempty"`
	NatureType     int64          `json:"natureType,omitempty"`
	Total          float64        `json:"total,omitempty"`
	Staffs         []StaffItem    `json:"staffs,omitempty"`
	Contract       []int64        `json:"contract,omitempty"`
	SignImg        string         `json:"signImg,omitempty"`
}

type PropertyItem struct {
	Property int64 `json:"property,omitempty"`
	Quantity int64 `json:"quantity,omitempty"`
}
type StaffItem struct {
	Id    int64 `json:"id,omitempty"`
	Ratio int64 `json:"ratio,omitempty"`
}

type OrderInfo struct {
	ID           int64        `json:"id"`
	OrderSn      string       `json:"order_sn"`
	Status       int64        `json:"status"`
	VenueId      int64        `json:"venue_id"`
	VenueName    string       `json:"venue_name"`
	MemberId     int64        `json:"member_id"`
	MemberName   string       `json:"member_name"`
	Source       int64        `json:"source"`
	SourceName   string       `json:"source_name"`
	Device       int64        `json:"device"`
	DeviceName   string       `json:"device_name"`
	CreateId     int64        `json:"create_id"`
	CreateName   string       `json:"create_name"`
	CreatedAt    string       `json:"createdAt"`
	UpdatedAt    string       `json:"updatedAt"`
	CompletionAt string       `json:"completionAt"`
	OrderPay     []OrderPay   `json:"order_pay"`
	OrderSales   []OrderSales `json:"order_sales"`
	OrderItem    OrderItem    `json:"order_item"`
	OrderAmount  OrderAmount  `json:"order_amount"`
}
type OrderListReq struct {
	ID           int64     `json:"id"`
	Status       int64     `json:"status"`
	VenueId      int64     `json:"venue_id"`
	MemberId     int64     `json:"member_id"`
	CreateId     int64     `json:"create_id"`
	CreateName   string    `json:"create_name"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	OrderSn      string    `json:"order_sn"`
	UserId       int64     `json:"user_id"`
	Source       int64     `json:"source"`
	Device       int64     `json:"device"`
	CompletionAt time.Time `json:"completionAt"`
	Page         int64     `json:"page"`
	PageSize     int64     `json:"pageSize"`
}
type OrderSales struct {
	ID        int64     `json:"id"`
	OrderId   int64     `json:"order_id"`
	SalesId   int64     `json:"sales_id"`
	SalesName string    `json:"sales_name"`
	Ratio     int64     `json:"ratio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type OrderPay struct {
	ID         int64     `json:"id"`
	OrderId    int64     `json:"order_id"`
	PaySn      string    `json:"pay_sn"`
	Remission  float64   `json:"remission"`
	Pay        float64   `json:"pay"`
	CreateId   int64     `json:"create_id"`
	CreateName string    `json:"create_name"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
type OrderItem struct {
	OrderId                int64  `json:"order_id"`
	ProductId              int64  `json:"product_id"`
	ProductName            string `json:"product_name"`
	RelatedUserProductId   int64  `json:"related_user_product_id"`
	RelatedUserProductName string `json:"related_user_product_name"`
	Data                   string `json:"data"`
}
type OrderAmount struct {
	OrderId   int64   `json:"order_id"`
	Total     float64 `json:"total"`
	Remission float64 `json:"remission"`
	Actual    float64 `json:"actual"`
	Residue   float64 `json:"residue"`
}
type UnifyPayReq struct {
	OrderSn   string  `json:"orderSn"`
	Note      string  `json:"note"`
	Payment   float64 `json:"payment"`
	Remission float64 `json:"remission"`
}
type QRPayReq struct {
	OrderSn string `json:"orderSn"`
	PayType string `json:"payType"`
}

// OrderStatus 订单状态
type OrderStatus int

const (
	OrderStatusUnpaid = iota
	OrderStatusPartial
	OrderStatusCompleted
	OrderStatusPending
	OrderStatusRefunded
	OrderStatusReversed
)

var OrderStatusNames = map[OrderStatus]string{
	OrderStatusUnpaid:    "未付款",
	OrderStatusPartial:   "部分付款",
	OrderStatusCompleted: "已完成",
	OrderStatusPending:   "退款中",
	OrderStatusRefunded:  "已退款",
	OrderStatusReversed:  "已取消",
}

func (s OrderStatus) String() string {
	return OrderStatusNames[s]
}
