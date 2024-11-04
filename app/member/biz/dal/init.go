package dal

import (
	"member/biz/dal/cache"
	"member/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
	cache.Init()
}
