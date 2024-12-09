package do

import (
	"saas/idl_gen/model/order"
)

type Order interface {
	Info(id int64) (info *order.OrderInfo, err error)
	Update(req order.UpdateOrderReq) error
	List() (resp []*order.OrderInfo, total int64, err error)
	UpdateStatus(ID int64, status int64) error
}
