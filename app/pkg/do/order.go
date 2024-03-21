package do

import "time"

type Order interface {
	Create(req OrderInfo) error
	Update(req OrderInfo) error
	Delete(id int64) error
	List(req OrderListReq) (resp []*OrderInfo, total int, err error)
	UpdateStatus(ID int64, status int64) error
	Info(ID int64) (roleInfo *OrderInfo, err error)
}
type OrderInfo struct {
	ID           int64     `json:"id,omitempty"`
	OrderSn      string    `json:"order_sn,omitempty"`
	Status       int64     `json:"status,omitempty"`
	VenueId      int64     `json:"venue_id,omitempty"`
	VenueName    string    `json:"venue_name,omitempty"`
	MemberId     int64     `json:"member_id,omitempty"`
	MemberName   string    `json:"member_name,omitempty"`
	Source       int64     `json:"source,omitempty"`
	SourceName   string    `json:"source_name,omitempty"`
	Device       int64     `json:"device,omitempty"`
	DeviceName   string    `json:"device_name,omitempty"`
	CreateId     int64     `json:"create_id,omitempty"`
	CreateName   int64     `json:"create_name,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	CompletionAt time.Time `json:"completionAt,omitempty"`
}
type OrderListReq struct {
	ID           int64     `json:"id,omitempty"`
	Status       int64     `json:"status,omitempty"`
	VenueId      int64     `json:"venue_id,omitempty"`
	MemberId     int64     `json:"member_id,omitempty"`
	CreateId     int64     `json:"create_id,omitempty"`
	CreateName   string    `json:"create_name,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	OrderSn      string    `json:"order_sn,omitempty"`
	UserId       int64     `json:"user_id,omitempty"`
	Source       int64     `json:"source,omitempty"`
	Device       int64     `json:"device,omitempty"`
	CompletionAt time.Time `json:"completionAt,omitempty"`
	Page         int64     `json:"page,omitempty"`
	PageSize     int64     `json:"pageSize,omitempty"`
}
type OrderSales struct {
	ID        int64     `json:"id,omitempty"`
	OrderId   int64     `json:"order_id,omitempty"`
	SalesId   int64     `json:"sales_id,omitempty"`
	SalesName string    `json:"sales_name,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
type OrderPay struct {
	ID         int64     `json:"id,omitempty"`
	OrderId    int64     `json:"order_id,omitempty"`
	PaySn      string    `json:"pay_sn,omitempty"`
	Remission  float64   `json:"remission,omitempty"`
	Pay        float64   `json:"pay,omitempty"`
	CreateId   int64     `json:"create_id,omitempty"`
	CreateName string    `json:"create_name,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}
type OrderItem struct {
}
type OrderAmount struct {
}
