package enums

import "fmt"

// PropertyType 自定义类型
type PropertyType string

// 声明每个枚举项的索引值
const (
	Card   PropertyType = "card"
	Course PropertyType = "course"
	Class  PropertyType = "class"
)

func GetAllValues() map[string]string {

	return map[string]string{
		"card":   "卡",
		"course": "私教课",
		"class":  "团课",
	}
}

func Test() {
	values := GetAllValues()

	fmt.Println(Card)
	fmt.Println(Course)
	fmt.Println(Class)
	fmt.Println(values)
}
