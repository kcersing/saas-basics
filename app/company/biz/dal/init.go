package dal

import (
	"company/biz/dal/cache"
	"company/biz/dal/casbin"
	"company/biz/dal/mysql"
	"company/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	casbin.Init()
	cache.Init()
}
