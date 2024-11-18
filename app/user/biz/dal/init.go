package dal

import (
	"user/biz/dal/cache"
	"user/biz/dal/casbin"
	"user/biz/dal/mysql"
	"user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	casbin.Init()
	cache.Init()
}
