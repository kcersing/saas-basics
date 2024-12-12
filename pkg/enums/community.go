package enums

func ReturnCommunityParticipantStatusValues(key int64) (values string) {
	switch key {
	case 1:
		values = "未报名"
	case 2:
		values = "报名中"
	case 3:
		values = "活动未开始"
	case 4:
		values = "活动中"
	default:
		values = "已结束"
	}
	return
}
