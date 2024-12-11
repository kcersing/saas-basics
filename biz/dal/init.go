package dal

import (
	"saas/biz/dal/cache"
	"saas/biz/dal/casbin"
	"saas/biz/dal/config"
	"saas/biz/dal/data"
	"saas/biz/dal/db"
	"saas/biz/dal/logger"
	"saas/biz/dal/minio"
	"saas/biz/dal/wechat"
)

func Init() {
	config.InitConfig()
	cache.InitCache()
	db.InitDB()
	data.NewInitDatabase()
	casbin.InitCasbin()
	logger.InitLogger()

	data.NewInitDatabase().InitDatabaseUser()
	data.NewInitDatabase().InitDatabaseDict()
	data.NewInitDatabase().InsertDatabaseMenuData()
	data.NewInitDatabase().InitDatabaseApi()
	wechat.InitWXPaymentApp()

	go minio.Init()
}
