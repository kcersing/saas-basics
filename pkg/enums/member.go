package enums

func ReturnMemberConditionValues(key int64) (values string) {
	switch key {
	case 1:
		values = "潜在"
	case 2:
		values = "正式"
	default:
		values = "状态异常"
	}
	return
}
func ReturnMemberGenderValues(key int64) (values string) {
	switch key {
	case 1:
		values = "女性"
	case 2:
		values = "男性"
	case 3:
		values = "未知"
	default:
		values = "异常"
	}
	return
}
func ReturnMemberGenderKey(values string) (key int64) {
	switch values {
	case "女性":
		key = 1
	case "男性":
		key = 2
	case "未知":
		key = 3
	default:
		key = 0
	}
	return
}
