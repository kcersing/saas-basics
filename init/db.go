package init

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/config"
	"saas/pkg/db"
	"saas/pkg/db/ent"
	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	hlog.Info(config.GlobalServerConfig)
	onceClient.Do(func() {
		DB = db.InitDB(config.GlobalServerConfig.PostgreSQLInfo.Host, config.GlobalServerConfig.IsProd)
	})
}
