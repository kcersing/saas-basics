package enums

// ReturnCommunityParticipantConditionValues 获取比赛人员状态
func ReturnCommunityParticipantConditionValues(key int64) (values string) {
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

//field.Int64("condition").
//			Default(1).
//			Optional().
//			Comment("状态[1:未报名;2:报名中;3:活动未开始;4:活动中;5:已结束]"),
