package enums

// 声明每个枚举项的索引值
// 产品类型
const (
	Card          string = "card"
	Course               = "course"
	Class                = "class"
	CoursePackage        = "coursePackage"
	Contest              = "contest"
	Bootcamp             = "bootcamp"
)

// ReturnProductTypeValues 获取产品类型
func ReturnProductTypeValues(key string) (values string) {
	switch key {
	case Card:
		values = "卡"
	case Course:
		values = "私教课"
	case Class:
		values = "团课"
	case CoursePackage:
		values = "私教课包"
	case Contest:
		values = "赛事"
	case Bootcamp:
		values = "训练营"
	default:
		values = "类型异常"
	}
	return
}

// 私教课类型
const (
	CourseOne  = "courseOne"
	CourseMore = "courseMore"
)

// ReturnCourseTypeValues 获取课程类型
func ReturnCourseTypeValues(key string) (values string) {
	switch key {
	case CourseOne:
		values = "一对一私教课"
	case CourseMore:
		values = "一对多私教课"
	default:
		values = "类型异常"
	}
	return
}
