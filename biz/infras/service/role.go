package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/base"

	"saas/idl_gen/model/auth"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/role"
	entuser "saas/pkg/db/ent/user"
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

func NewRole(ctx context.Context, c *app.RequestContext) do.Role {
	return &Role{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (r Role) Create(req *auth.RoleInfo) error {
	roleEnt, err := r.db.Role.Create().
		SetName(*req.Name).
		SetValue(*req.Value).
		SetDefaultRouter(*req.DefaultRouter).
		SetStatus(*req.Status).
		SetRemark(*req.Remark).
		SetOrderNo(*req.OrderNo).
		Save(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Role failed")
		return err
	}

	// set roleEnt to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(roleEnt.ID)), roleEnt, 1, 1*time.Hour)
	return nil
}

func (r Role) Update(req *auth.RoleInfo) error {
	roleEnt, err := r.db.Role.UpdateOneID(*req.Id).
		SetName(*req.Name).
		SetValue(*req.Value).
		SetDefaultRouter(*req.DefaultRouter).
		SetStatus(*req.Status).
		SetRemark(*req.Remark).
		SetOrderNo(*req.OrderNo).
		SetUpdatedAt(time.Now()).
		Save(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Role failed")
		return err
	}

	// set roleEnt to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(roleEnt.ID)), roleEnt, 1, 1*time.Hour)
	return nil
}

func (r Role) Delete(id int64) error {
	// whether role is used by user
	exist, err := r.db.User.Query().Where(entuser.RoleIDEQ(id)).Exist(r.ctx)
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
func entRoleInfo(entRole ent.Role) *auth.RoleInfo {
	createdAt := entRole.CreatedAt.Format(time.DateTime)
	updatedAt := entRole.UpdatedAt.Format(time.DateTime)
	return &auth.RoleInfo{
		Id:            &entRole.ID,
		Name:          &entRole.Name,
		Value:         &entRole.Value,
		DefaultRouter: &entRole.DefaultRouter,
		Status:        &entRole.Status,
		Remark:        &entRole.Remark,
		OrderNo:       &entRole.OrderNo,
		CreatedAt:     &createdAt,
		UpdatedAt:     &updatedAt,
	}
}
func (r Role) RoleInfoByID(ID int64) (roleInfo *auth.RoleInfo, err error) {
	roleInterface, ok := r.cache.Get("roleData" + strconv.Itoa(int(ID)))
	if ok {
		if l, ok := roleInterface.(*ent.Role); ok {
			return entRoleInfo(*l), nil
		}
	}
	// get role from db
	roleEnt, err := r.db.Role.Query().Where(role.IDEQ(ID)).Only(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Role failed")
		return nil, err
	}
	// set role to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(ID)), roleEnt, 1, 1*time.Hour)
	// convert to RoleInfo
	return entRoleInfo(*roleEnt), nil
}

func (r Role) List(req *base.PageInfoReq) (roleInfoList []*auth.RoleInfo, total int64, err error) {

	roleEntList, err := r.db.Role.Query().
		Where(role.Delete(0)).
		Order(ent.Asc(role.FieldOrderNo)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "get RoleList failed")
		return nil, 0, err
	}
	// convert to List
	for _, roleEnt := range roleEntList {
		menuArr, _ := roleEnt.QueryMenus().GroupBy(menu.FieldID).Ints(r.ctx)
		var mArr []int64
		for v := range menuArr {
			mArr = append(mArr, cast.ToInt64(v))
		}

		var rArr []int64
		for v2 := range roleEnt.Apis {
			rArr = append(rArr, cast.ToInt64(v2))
		}
		re := entRoleInfo(*roleEnt)
		re.Menus = mArr
		re.Apis = rArr
		roleInfoList = append(roleInfoList, re)
	}
	t, _ := r.db.Role.Query().Count(r.ctx)
	total = int64(t)
	return
}

func (r Role) UpdateStatus(ID int64, status int64) error {

	roleEnt, err := r.db.Role.UpdateOneID(ID).SetStatus(status).Save(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Role status failed")
		return err
	}
	// set role to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(ID)), roleEnt, 1, 1*time.Hour)

	return nil
}
