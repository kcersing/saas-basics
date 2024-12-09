package enums

const (
	Buy = "buy"
)

func ReturnOrderNatureValues(key string) (values string) {
	switch key {
	case Buy:
		values = "购买"
		values = "业务异常"
	}
	return
}
