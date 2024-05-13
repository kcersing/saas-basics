package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type EntryLogs struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (e EntryLogs) Create(logsReq *do.EntryLogsInfo) error {
	//TODO implement me
	panic("implement me")
}

func (e EntryLogs) List(req *do.EntryLogsListReq) (list []*do.EntryLogsInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func NewEntryLogs(ctx context.Context, c *app.RequestContext) do.EntryLogs {
	return &EntryLogs{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
