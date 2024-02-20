package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/app/admin/config"
	infras "saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	menu2 "saas/pkg/db/ent/menu"
)

type Menu struct {
	ctx  context.Context
	c    *app.RequestContext
	salt string
	db   *ent.Client
}

func (m Menu) Create(menuReq *do.MenuInfo) error {
	//// get menu level
	//if req.ParentID == 0 {
	//	// it is a first level menu
	//	req.ParentID = 1
	//	req.Level = 1
	//} else {
	//	// it is a children level menu
	//	// get parent menu level
	//	parent, err := m.db.Menu.Query().Where(menu2.IDEQ(req.ParentID)).First(m.ctx)
	//	if err != nil {
	//		return errors.Wrap(err, "query menu failed")
	//	}
	//	// set menu level
	//	req.Level = parent.MenuLevel + 1
	//}
	//
	//// create menu
	//err := m.db.Menu.Create().
	//	SetMenuLevel(req.Level).
	//	SetMenuType(req.MenuType).
	//	SetParentID(req.ParentID).
	//	SetPath(req.Path).
	//	SetName(req.Name).
	//	SetRedirect(req.Redirect).
	//	SetComponent(req.Component).
	//	SetOrderNo(req.OrderNo).
	//	SetDisabled(req.Disabled).
	//	// meta
	//	SetTitle(req.Meta.Title).
	//	SetIcon(req.Meta.Icon).
	//	SetHideMenu(req.Meta.HideMenu).
	//	SetHideBreadcrumb(req.Meta.HideBreadcrumb).
	//	SetCurrentActiveMenu(req.Meta.CurrentActiveMenu).
	//	SetIgnoreKeepAlive(req.Meta.IgnoreKeepAlive).
	//	SetHideTab(req.Meta.HideTab).
	//	SetFrameSrc(req.Meta.FrameSrc).
	//	SetCarryParam(req.Meta.CarryParam).
	//	SetHideChildrenInMenu(req.Meta.HideChildrenInMenu).
	//	SetAffix(req.Meta.Affix).
	//	SetDynamicLevel(req.Meta.DynamicLevel).
	//	SetRealPath(req.Meta.RealPath).
	//	Exec(m.ctx)
	//
	//if err != nil {
	//	return errors.Wrap(err, "create menu failed")
	//}
	panic("implement me")
}

func (m Menu) Update(menuReq *do.MenuInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) Delete(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) ListByRole(roleID uint64) (list []*do.MenuInfoTree, total uint64, err error) {
	//TODO implement me
	panic("implement me")
}

func (m Menu) List(req *do.MenuListReq) (list []*do.MenuInfoTree, total int, err error) {

	all, err := m.db.Menu.
		Query().
		Order(ent.Asc(menu2.FieldMenuLevel)).
		WithChildren().
		All(m.ctx)
	if err != nil {

	}

	hlog.Info(all)
	hlog.Info(err)
	panic("implement me")
}

func (m Menu) CreateMenuParam(req *do.MenuParam) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) UpdateMenuParam(req *do.MenuParam) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) DeleteMenuParam(menuParamID uint64) error {
	//TODO implement me
	panic("implement me")
}

func (m Menu) MenuParamListByMenuID(menuID uint64) (list []do.MenuParam, total uint64, err error) {
	//TODO implement me
	panic("implement me")
}

func NewMenu(ctx context.Context, c *app.RequestContext) do.Menu {
	return &Menu{
		ctx:  ctx,
		c:    c,
		salt: config.GlobalServerConfig.MysqlInfo.Salt,
		db:   infras.DB,
	}
}
