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

	CreateParticipantOrder(req CreateParticipantOrderReq) (orderOne *order.OrderInfo, err error)
}
type CreateParticipantOrderReq struct {
	Member    *ent.Member
	Device    string
	ContestId int64
}
