package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/role"
	"saas/pkg/db/ent/user"
	"strconv"
	"time"
)

type Role struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (r Role) Create(req do.RoleInfo) error {
	roleEnt, err := r.db.Role.Create().
		SetName(req.Name).
		SetValue(req.Value).
		SetDefaultRouter(req.DefaultRouter).
		SetStatus(uint8(req.Status)).
		SetRemark(req.Remark).
		SetOrderNo(req.OrderNo).
		Save(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Role failed")
		return err
	}

	// set roleEnt to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(roleEnt.ID)), roleEnt, 1, 24*time.Hour)
	return nil
}

func (r Role) Update(req do.RoleInfo) error {
	roleEnt, err := r.db.Role.UpdateOneID(req.ID).
		SetName(req.Name).
		SetValue(req.Value).
		SetDefaultRouter(req.DefaultRouter).
		SetStatus(uint8(req.Status)).
		SetRemark(req.Remark).
		SetOrderNo(req.OrderNo).
		SetUpdatedAt(time.Now()).
		Save(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Role failed")
		return err
	}

	// set roleEnt to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(roleEnt.ID)), roleEnt, 1, 24*time.Hour)
	return nil
}

func (r Role) Delete(id uint64) error {
	// whether role is used by user
	exist, err := r.db.User.Query().Where(user.RoleIDEQ(id)).Exist(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "query user - role failed")
		return err
	}
	if exist {
		return errors.New("role is used by user")
	}
	// delete role from db
	err = r.db.Role.DeleteOneID(id).Exec(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "delete Role failed")
		return err
	}
	// delete role from cache
	r.cache.Del("roleData" + strconv.Itoa(int(id)))
	return nil
}

func (r Role) RoleInfoByID(ID uint64) (roleInfo *do.RoleInfo, err error) {
	roleInterface, ok := r.cache.Get("roleData" + strconv.Itoa(int(ID)))
	if ok {
		if l, ok := roleInterface.(*ent.Role); ok {
			return &do.RoleInfo{
				ID:            l.ID,
				Name:          l.Name,
				Value:         l.Value,
				DefaultRouter: l.DefaultRouter,
				Status:        uint64(l.Status),
				Remark:        l.Remark,
				OrderNo:       l.OrderNo,
				CreatedAt:     l.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:     l.UpdatedAt.Format("2006-01-02 15:04:05"),
			}, nil
		}
	}
	// get role from db
	roleEnt, err := r.db.Role.Query().Where(role.IDEQ(ID)).Only(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Role failed")
		return nil, err
	}
	// set role to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(ID)), roleEnt, 1, 24*time.Hour)
	// convert to RoleInfo
	roleInfo = &do.RoleInfo{
		ID:            roleEnt.ID,
		Name:          roleEnt.Name,
		Value:         roleEnt.Value,
		DefaultRouter: roleEnt.DefaultRouter,
		Status:        uint64(roleEnt.Status),
		Remark:        roleEnt.Remark,
		OrderNo:       roleEnt.OrderNo,
		CreatedAt:     roleEnt.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     roleEnt.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return
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
