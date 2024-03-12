package admin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/role"
)

type Auth struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	Cbs   *casbin.Enforcer
	db    *ent.Client
	cache *ristretto.Cache
}

func (a Auth) UpdateApiAuth(roleIDStr string, infos []*do.ApiAuthInfo) error {
	// clear old policies
	var oldPolicies [][]string
	oldPolicies = a.Cbs.GetFilteredPolicy(0, roleIDStr)
	if len(oldPolicies) != 0 {
		removeResult, err := a.Cbs.RemoveFilteredPolicy(0, roleIDStr)
		if err != nil {
			return err
		}
		if !removeResult {
			return errors.New("casbin policies remove failed")
		}
	}
	// add new policies
	var policies [][]string
	for _, v := range infos {
		policies = append(policies, []string{roleIDStr, v.Path, v.Method})
	}
	addResult, err := a.Cbs.AddPolicies(policies)
	if err != nil {
		return err
	}
	if !addResult {
		return errors.New("casbin policies add failed")
	}
	return nil
}

func (a Auth) ApiAuth(roleIDStr string) (infos []*do.ApiAuthInfo, err error) {

	policies := a.Cbs.GetFilteredPolicy(0, roleIDStr)

	hlog.Info("111111111111111111111111111111111")
	hlog.Info(policies)
	hlog.Info("111111111111111111111111111111111")
	for _, v := range policies {
		infos = append(infos, &do.ApiAuthInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return
}

func (a Auth) UpdateMenuAuth(roleID int64, menuIDs []int64) error {
	tx, err := a.db.Tx(a.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction err")
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				hlog.Error("UpdateMenuAuthority err:", err, "rollback err:", rollbackErr)
			}
		}
	}()

	err = tx.Role.UpdateOneID(roleID).ClearMenus().Exec(a.ctx)
	if err != nil {
		return errors.Wrap(err, "delete role's menu failed, error")
	}

	err = tx.Role.UpdateOneID(roleID).AddMenuIDs(menuIDs...).Exec(a.ctx)
	if err != nil {
		return errors.Wrap(err, "add role's menu failed, error")
	}

	return tx.Commit()
}

func (a Auth) MenuAuth(roleID int64) (menuIDs []int64, err error) {
	menus, err := a.db.Role.Query().Where(role.IDEQ(roleID)).QueryMenus().All(a.ctx)
	for _, v := range menus {
		if v.ID != 1 {
			menuIDs = append(menuIDs, v.ID)
		}
	}
	return
}

func NewAuth(ctx context.Context, c *app.RequestContext) do.Auth {
	return &Auth{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
		Cbs:   infras.CasbinEnforcer(),
	}
}
