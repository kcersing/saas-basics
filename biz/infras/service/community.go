package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/community"
)

type Community struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (c Community) CreateCommunity(req community.CommunityInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) UpdateCommunity(req community.CommunityInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) DeleteCommunity(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) CommunityList(req community.CommunityListReq) (resp []*community.CommunityInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Community) CommunityInfo(id int64) (resp *community.CommunityInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Community) UpdateCommunityStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) UpdateCommunityShow(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Community) DelCommunity(ID int64) error {
	//TODO implement me
	panic("implement me")
}

func NewCommunity(ctx context.Context, c *app.RequestContext) do.Community {
	return &Community{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
