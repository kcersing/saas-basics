package enums

// ReturnParticipantStatusValues 获取比赛人员状态
func ReturnParticipantStatusValues(key int64) (values string) {
	switch key {
	case 1:
		values = "报名中"
	case 2:
		values = "报名完成"
	case 3:
		values = "取消报名"
	case 4:
		values = "报名失效"
	case 5:
		values = "报名成功"
	default:
		values = "状态异常"
	}
	return
}
