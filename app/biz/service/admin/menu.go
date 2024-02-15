package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/app/biz/do"
	"saas/app/biz/model/admin/menu"
	"saas/app/biz/model/base"
	"saas/app/config"
	"saas/pkg/db"
	"saas/pkg/db/ent"
	menu2 "saas/pkg/db/ent/menu"
	"strconv"
)

type Menu struct {
	ctx  context.Context
	c    *app.RequestContext
	salt string
	db   *ent.Client
}

func (m Menu) CreateMenu(req menu.CreateOrUpdateMenuReq) error {
	//TODO implement me
	parentID, _ := strconv.Atoi(req.ParentID)
	_, err := m.db.Menu.
		Create().
		SetParentID(parentID).
		SetRouteName(req.RouteName).
		SetRoutePath(req.RoutePath).
		SetStatus(req.Status).
		SetMenuName(req.Name).
		SetMenuType(req.MenuType).
		SetIconType(req.IconType).
		SetIcon(req.Icon).
		SetI18nKey(req.I18nKey).
		SetLevel(req.Level).
		Save(m.ctx)
	if err != nil {
		return err
	}
	return err
}

func (m Menu) UpdateMenu(req menu.CreateOrUpdateMenuReq) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) DeleteMenu(id string) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) MenuByRole(id string) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) MenuList(req base.PageInfoReq) []*ent.Menu {
	//TODO implement me

	all, err := m.db.Menu.
		Query().
		Order(ent.Asc(menu2.FieldLevel)).
		All(m.ctx)
	if err != nil {

	}

	hlog.Info(all)
	hlog.Info(err)
	return all
}

func (m Menu) CreateMenuParam(req menu.CreateOrUpdateMenuParamReq) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) UpdateMenuParam(req menu.CreateOrUpdateMenuParamReq) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) DeleteMenuParam(id string) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) MenuParamListByMenuID(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewMenu(ctx context.Context, c *app.RequestContext) do.Menu {
	return &Menu{
		ctx:  ctx,
		c:    c,
		salt: config.GlobalServerConfig.MysqlInfo.Salt,
		db:   db.InitDB(config.GlobalServerConfig.MysqlInfo.Host),
	}
}
