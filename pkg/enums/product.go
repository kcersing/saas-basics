package enums

// 声明每个枚举项的索引值
// 产品类型
const (
	Card          = "card"
	Course        = "course"
	CoursePackage = "coursePackage"
	Contest       = "contest"
	Bootcamp      = "bootcamp"
	Sms           = "Sms"
	CourseOne     = "courseOne"
	CourseMore    = "courseMore"
	CardTerm      = "cardTerm"
	CardSub       = "cardSub"
	Lessons       = "lessons"
)

// ReturnProductTypeValues 获取产品类型
func ReturnProductTypeValues(key string) (values string) {
	switch key {
	case Card:
		values = "卡"
	case Course:
		values = "私教课"
	case Lessons:
		values = "团课"
	case CoursePackage:
		values = "私教课包"
	case Contest:
		values = "赛事"
	case Bootcamp:
		values = "训练营"

	case CourseOne:
		values = "一对一私教课"
	case CourseMore:
		values = "一对多私教课"
	case CardTerm:
		values = "期限卡"
	case CardSub:
		values = "次卡"
	case Sms:
		values = "短信"

	default:
		values = "类型异常"
	}
	return
}
