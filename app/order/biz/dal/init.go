package dal

import (
	"order/biz/dal/cache"
	"order/biz/dal/casbin"
	"order/biz/dal/mysql"
	"order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	casbin.Init()
	cache.Init()
}
