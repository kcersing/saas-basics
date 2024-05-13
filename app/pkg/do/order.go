package do

import "time"

type Order interface {
	Create(req CreateOrder) error
	Update(req OrderInfo) error
	Delete(id int64) error
	List(req OrderListReq) (resp []*OrderInfo, total int, err error)
	UpdateStatus(ID int64, status int64) error
	Info(ID int64) (roleInfo *OrderInfo, err error)
}

type CreateOrder struct {
	VenueId    int64   `json:"venue_id"`
	MemberId   int64   `json:"member_id"`
	CreateId   int64   `json:"create_id"`
	Total      float64 `json:"total"`
	Sales      []int64 `json:"sales"`
	ProductId  int64   `json:"product_id"`
	Quantity   int64   `json:"quantity"`
	ContractId int64   `json:"contract_id"`
	AssignAt   string  `json:"assign_at"`
}

type OrderInfo struct {
	ID           int64     `json:"id"`
	OrderSn      string    `json:"order_sn"`
	Status       int64     `json:"status"`
	VenueId      int64     `json:"venue_id"`
	VenueName    string    `json:"venue_name"`
	MemberId     int64     `json:"member_id"`
	MemberName   string    `json:"member_name"`
	Source       int64     `json:"source"`
	SourceName   string    `json:"source_name"`
	Device       int64     `json:"device"`
	DeviceName   string    `json:"device_name"`
	CreateId     int64     `json:"create_id"`
	CreateName   int64     `json:"create_name"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CompletionAt time.Time `json:"completionAt"`
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
}
type OrderAmount struct {
}
