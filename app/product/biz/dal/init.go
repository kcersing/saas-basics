package dal

import (
	"product/biz/dal/cache"
	"product/biz/dal/casbin"
	"product/biz/dal/mysql"
	"product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	casbin.Init()
	cache.Init()
}
