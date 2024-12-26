package enums

const (
	Buy = "buy"
)

func ReturnOrderNatureValues(key string) (values string) {
	switch key {
	case Buy:
		values = "购买"
	default:
		values = "业务异常"
	}
	return
}
func ReturnOrderStatusValues(key int64) (values string) {
	switch key {
	case 1:
		values = "待支付"
	case 2:
		values = "订单取消"
	case 3:
		values = "订单失效"
	case 4:
		values = "订单退款完成"
	case 5:
		values = "订单已完成"
	default:
		values = "状态异常"
	}
	return
}
func ReturnOrderDeviceValues(key string) (values string) {
	switch key {
	case "pc":
		values = "电脑端"
	case "wxc":
		values = "微信小程序"

	default:
		values = "状态异常"
	}
	return
}
func ReturnOrderPayWayValues(key string) (values string) {
	switch key {
	case "pc":
		values = "电脑支付"
	case "wxc":
		values = "微信小程序支付"

	default:
		values = "状态异常"
	}
	return
}
