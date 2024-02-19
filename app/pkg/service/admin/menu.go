package admin

//
//import (
//	"context"
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/common/hlog"
//	"github.com/pkg/errors"
//	"saas/app/biz/do"
//	"saas/app/biz/model/admin/menu"
//	"saas/app/biz/model/base"
//	"saas/app/config"
//	"saas/pkg/db"
//	"saas/pkg/db/ent"
//	menu2 "saas/pkg/db/ent/menu"
//)
//
//type Menu struct {
//	ctx  context.Context
//	c    *app.RequestContext
//	salt string
//	db   *ent.Client
//}
//
//func (m Menu) CreateMenu(req menu.CreateOrUpdateMenuReq) error {
//	// get menu level
//	if req.ParentID == 0 {
//		// it is a first level menu
//		req.ParentID = 1
//		req.Level = 1
//	} else {
//		// it is a children level menu
//		// get parent menu level
//		parent, err := m.db.Menu.Query().Where(menu2.IDEQ(req.ParentID)).First(m.ctx)
//		if err != nil {
//			return errors.Wrap(err, "query menu failed")
//		}
//		// set menu level
//		req.Level = parent.MenuLevel + 1
//	}
//
//	// create menu
//	err := m.db.Menu.Create().
//		SetMenuLevel(req.Level).
//		SetMenuType(req.MenuType).
//		SetParentID(req.ParentID).
//		SetPath(req.Path).
//		SetName(req.Name).
//		SetRedirect(req.Redirect).
//		SetComponent(req.Component).
//		SetOrderNo(req.OrderNo).
//		SetDisabled(req.Disabled).
//		// meta
//		SetTitle(req.Meta.Title).
//		SetIcon(req.Meta.Icon).
//		SetHideMenu(req.Meta.HideMenu).
//		SetHideBreadcrumb(req.Meta.HideBreadcrumb).
//		SetCurrentActiveMenu(req.Meta.CurrentActiveMenu).
//		SetIgnoreKeepAlive(req.Meta.IgnoreKeepAlive).
//		SetHideTab(req.Meta.HideTab).
//		SetFrameSrc(req.Meta.FrameSrc).
//		SetCarryParam(req.Meta.CarryParam).
//		SetHideChildrenInMenu(req.Meta.HideChildrenInMenu).
//		SetAffix(req.Meta.Affix).
//		SetDynamicLevel(req.Meta.DynamicLevel).
//		SetRealPath(req.Meta.RealPath).
//		Exec(m.ctx)
//
//	if err != nil {
//		return errors.Wrap(err, "create menu failed")
//	}
//	return nil
//
//}
//
//func (m Menu) UpdateMenu(req menu.CreateOrUpdateMenuReq) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m Menu) DeleteMenu(id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m Menu) MenuByRole(id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m Menu) MenuList(req base.PageInfoReq) []*ent.Menu {
//	//TODO implement me
//
//	all, err := m.db.Menu.
//		Query().
//		Order(ent.Asc(menu2.FieldMenuLevel)).WithChildren().
//		All(m.ctx)
//	if err != nil {
//
//	}
//
//	hlog.Info(all)
//	hlog.Info(err)
//	return all
//}
//
//func (m Menu) CreateMenuParam(req menu.CreateOrUpdateMenuParamReq) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m Menu) UpdateMenuParam(req menu.CreateOrUpdateMenuParamReq) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m Menu) DeleteMenuParam(id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m Menu) MenuParamListByMenuID(id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func NewMenu(ctx context.Context, c *app.RequestContext) do.Menu {
//	return &Menu{
//		ctx:  ctx,
//		c:    c,
//		salt: config.GlobalServerConfig.MysqlInfo.Salt,
//		db:   db.InitDB(config.GlobalServerConfig.MysqlInfo.Host),
//	}
//}
