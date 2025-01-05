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
	"saas/idl_gen/model/member"
)

type MemberProduct struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m MemberProduct) UpdateMemberProduct(req member.CreateOrUpdateMemberReq) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) MemberProductInfo(id int64) (info *member.MemberInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberProduct) entMemberInfo(req *ent.MemberProduct) (info *member.MemberInfo, err error) {

	return info, err
}

func (m MemberProduct) MemberProductList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func NewMemberProduct(ctx context.Context, c *app.RequestContext) do.MemberProduct {
	return &MemberProduct{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
