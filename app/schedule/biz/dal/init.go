package dal

import (
	"schedule/biz/dal/cache"
	"schedule/biz/dal/casbin"
	"schedule/biz/dal/mysql"
	"schedule/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	casbin.Init()
	cache.Init()
}
