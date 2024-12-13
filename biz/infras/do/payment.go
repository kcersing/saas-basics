package do

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/refund/response"
	"saas/idl_gen/model/payment"
)

type Payment interface {
	UnifyPay(req payment.PayReq) (interface{}, error)
	QRPay(req payment.PayReq) (interface{}, error)
	RefundOrder(req payment.RefundOrderReq) (*response.ResponseRefund, error)

	Notify(req payment.NotifyReq) error
}
