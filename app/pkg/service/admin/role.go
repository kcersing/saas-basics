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

type Role struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (r Role) Create(req do.RoleInfo) error {
	//TODO implement me
	panic("implement me")
}

func (r Role) Update(req do.RoleInfo) error {
	//TODO implement me
	panic("implement me")
}

func (r Role) Delete(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (r Role) RoleInfoByID(ID uint64) (roleInfo *do.RoleInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Role) List(req *do.RoleListReq) (roleInfoList []*do.RoleInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Role) UpdateStatus(ID uint64, status uint8) error {
	//TODO implement me
	panic("implement me")
}

func NewRole(ctx context.Context, c *app.RequestContext) do.Role {
	return &Role{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
