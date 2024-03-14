package do

type Order interface {
	Create(req OrderInfo) error
	Update(req OrderInfo) error
	Delete(id int64) error
	List(req OrderListReq) (resp []*OrderInfo, total int, err error)
	UpdateStatus(ID int64, status int8) error
	Info(ID int64) (roleInfo *OrderInfo, err error)
}
type OrderInfo struct {
}
type OrderListReq struct {
}
