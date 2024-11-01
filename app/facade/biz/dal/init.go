package dal

import (
	"facade/biz/dal/mysql"
	"facade/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
