package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
)

type Common struct {
	ctx   context.Context
	c     *app.RequestContext
	db    *ent.Client
	cache *ristretto.Cache
}

func NewCommon(ctx context.Context, c *app.RequestContext) do.Common {
	return &Common{
		ctx:   ctx,
		c:     c,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
