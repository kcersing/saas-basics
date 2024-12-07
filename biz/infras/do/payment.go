package do

import "saas/idl_gen/model/payment"

type Payment interface {
	WXPay(payment.WXPayReq) error
	WXNotify(payment.WXNotifyReq) error
}
