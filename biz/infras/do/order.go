package do

import (
	"saas/biz/dal/db/ent"
	"saas/idl_gen/model/order"
)

type Order interface {
	Info(id int64) (info *order.OrderInfo, err error)
	Update(req *order.UpdateOrderReq) error
	List(req *order.ListOrderReq) (resp []*order.OrderInfo, total int, err error)
	UpdateStatus(ID int64, status int64) error
	OrderListExport(req *order.ListOrderReq) (string, error)

	CreateParticipantOrder(req CreateParticipantOrderReq) (orderOne *order.OrderInfo, err error)
	CreateSmsOrder(req CreateSmsOrderReq) (orderOne *order.OrderInfo, err error)

	Buy(req *order.BuyReq) (orderOne *order.OrderInfo, err error)
	FinishOrder(c FinishOrder) (err error)

	OrderAllCount(req *order.OrderAllCountReq) (resp []*order.OrderCountInfo, actual float64, total int, err error)
}
type CreateParticipantOrderReq struct {
	Member *ent.Member
	Device string
	NodeId int64
	Fee    float64
}
type CreateSmsOrderReq struct {
	VenueId int64
	Number  int64
	Fee     float64
}
type FinishOrder struct {
	Sn            string
	Fee           int64
	TransactionId string
	Transaction   []byte
}
