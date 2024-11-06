package dal

import (
	"system/biz/dal/mysql"
	"system/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
