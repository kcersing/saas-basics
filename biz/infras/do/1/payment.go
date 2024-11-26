package do

type Payment interface {
	UnifyPay(req UnifyPayReq) error
	QRPay(req QRPayReq) (string, error)
}
