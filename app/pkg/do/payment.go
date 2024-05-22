package do

type Payment interface {
	UnifyPay(req UnifyPayReq) error
	QRPay(req QRPayReq)
	QueryOrder(OrderSn string)
	CloseOrder(OrderSn string)
}
type UnifyPayReq struct {
	OrderSn         string `json:"order_sn"`
	TotalAmount     int64  `json:"total_amount"`
	Subject         string `json:"subject"`
	PayWay          string `json:"pay_way"`
	AppId           string `json:"app_id"`
	SubOpenId       string `json:"sub_open_id"`
	NotifyUrl       string `json:"notify_url"`
	ClientIp        string `json:"client_ip"`
	OrderExpiration int32  `json:"order_expiration"`
}
type QRPayReq struct {
	OrderSn     string `json:"order_sn"`
	TotalAmount int64  `json:"total_amount"`
	Subject     string `json:"subject"`
	AuthCode    string `json:"auth_code"`
	NotifyUrl   string `json:"notify_url"`
	ClientIp    string `json:"client_ip"`
}
