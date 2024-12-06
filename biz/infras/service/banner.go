package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/banner"
	"saas/pkg/db/ent"
)

type Banner struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (b Banner) Create(req banner.BannerInfo) error {
	//TODO implement me
	panic("implement me")
}

func (b Banner) Update(req banner.BannerInfo) error {
	//TODO implement me
	panic("implement me")
}

func (b Banner) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (b Banner) List() (resp []*banner.BannerInfo, total int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (b Banner) UpdateShow(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func NewBanner(ctx context.Context, c *app.RequestContext) do.Banner {
	return &Banner{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
