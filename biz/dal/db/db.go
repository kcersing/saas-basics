package db

import (
	"saas/config"
	"saas/pkg/db"
	"saas/pkg/db/ent"
	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	onceClient.Do(func() {
		DB = db.InitDB(config.GlobalServerConfig.MySQLInfo.Host, config.GlobalServerConfig.IsProd)
	})
}
