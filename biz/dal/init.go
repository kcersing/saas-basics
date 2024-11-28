package dal

import (
	"saas/biz/dal/cache"
	"saas/biz/dal/casbin"
	"saas/biz/dal/config"
	"saas/biz/dal/data"
	"saas/biz/dal/db"
	"saas/biz/dal/logger"
)

func Init() {
	config.InitConfig()
	cache.InitCache()
	db.InitDB()
	data.NewInitDatabase()
	casbin.InitCasbin()
	logger.InitLogger()
}
