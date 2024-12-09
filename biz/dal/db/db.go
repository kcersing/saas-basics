package db

import (
	"saas/biz/dal/db/ent"
	"saas/config"

	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	onceClient.Do(func() {
		DB = InItDB(config.GlobalServerConfig.MySQLInfo.Host, config.GlobalServerConfig.IsProd)
	})
}
