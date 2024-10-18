package admin

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	infras "saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	menu2 "saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/role"
	"time"
)

type Menu struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m Menu) Create(menuReq *do.MenuInfo) error {
	// get menu level
	if menuReq.ParentID == 0 {
		// it is a first level menu
		menuReq.ParentID = 1
		//menuReq.Level = 1
	}
	//else {
	//// it is a children level menu
	//// get parent menu level
	//parent, err := m.db.Menu.Query().Where(menu2.IDEQ(menuReq.ParentID)).First(m.ctx)
	//if err != nil {
	//	return errors.Wrap(err, "query menu failed")
	//}
	//// set menu level
	//menuReq.Level = parent.MenuLevel + 1
	//}

	// create menu
	err := m.db.Menu.Create().
		SetParentID(menuReq.ParentID).
		SetPath(menuReq.Path).
		SetName(menuReq.Name).
		SetOrderNo(menuReq.OrderNo).
		SetDisabled(menuReq.Disabled).

		//SetMenuLevel(menuReq.Level).
		//SetMenuType(menuReq.MenuType).
		//SetRedirect(menuReq.Redirect).
		//SetComponent(menuReq.Component).
		// meta
		//SetTitle(menuReq.Meta.Title).
		//SetIcon(menuReq.Meta.Icon).
		//SetHideMenu(menuReq.Meta.HideMenu).
		//SetHideBreadcrumb(menuReq.Meta.HideBreadcrumb).
		//SetCurrentActiveMenu(menuReq.Meta.CurrentActiveMenu).
		//SetIgnoreKeepAlive(menuReq.Meta.IgnoreKeepAlive).
		//SetHideTab(menuReq.Meta.HideTab).
		//SetFrameSrc(menuReq.Meta.FrameSrc).
		//SetCarryParam(menuReq.Meta.CarryParam).
		//SetHideChildrenInMenu(menuReq.Meta.HideChildrenInMenu).
		//SetAffix(menuReq.Meta.Affix).
		//SetDynamicLevel(menuReq.Meta.DynamicLevel).
		//SetRealPath(menuReq.Meta.RealPath).
		Exec(m.ctx)

	if err != nil {
		return errors.Wrap(err, "create menu failed")
	}
	return nil
}

func (m Menu) Update(menuReq *do.MenuInfo) error {
	// get menu level
	if menuReq.ParentID == 0 {
		// it is a first level menu
		menuReq.ParentID = 1
		//menuReq.Level = 1
	}
	//else {
	//// it is a children level menu
	//// get parent menu level
	//parent, err := m.db.Menu.Query().Where(menu2.IDEQ(menuReq.ParentID)).First(m.ctx)
	//if err != nil {
	//	return errors.Wrap(err, "query menu failed")
	//}
	//// set menu level
	//menuReq.Level = parent.MenuLevel + 1
	//}

	// update menu
	err := m.db.Menu.UpdateOneID(menuReq.ID).
		SetParentID(menuReq.ParentID).
		SetPath(menuReq.Path).
		SetName(menuReq.Name).
		SetOrderNo(menuReq.OrderNo).

		//SetMenuLevel(menuReq.Level).
		//SetMenuType(menuReq.MenuType).
		//SetRedirect(menuReq.Redirect).
		//SetComponent(menuReq.Component).
		SetDisabled(menuReq.Disabled).
		// meta
		//SetTitle(menuReq.Meta.Title).
		//SetIcon(menuReq.Meta.Icon).
		//SetHideMenu(menuReq.Meta.HideMenu).
		//SetHideBreadcrumb(menuReq.Meta.HideBreadcrumb).
		//SetCurrentActiveMenu(menuReq.Meta.CurrentActiveMenu).
		//SetIgnoreKeepAlive(menuReq.Meta.IgnoreKeepAlive).
		//SetHideTab(menuReq.Meta.HideTab).
		//SetFrameSrc(menuReq.Meta.FrameSrc).
		//SetCarryParam(menuReq.Meta.CarryParam).
		//SetHideChildrenInMenu(menuReq.Meta.HideChildrenInMenu).
		//SetAffix(menuReq.Meta.Affix).
		//SetDynamicLevel(menuReq.Meta.DynamicLevel).
		//SetRealPath(menuReq.Meta.RealPath).
		Exec(m.ctx)
	if err != nil {
		return errors.Wrap(err, "update menu failed")
	}

	return nil
}

func (m Menu) Delete(id int64) error {

	// find out the menu whether it has children
	// if it has children, it can not be deleted
	exist, err := m.db.Menu.Query().Where(menu2.ParentIDEQ(id)).Exist(m.ctx)
	if err != nil {
		return errors.Wrap(err, "query menu failed")
	}
	if exist {
		return errors.New("menu has children, can not be deleted")
	}

	// delete menu
	err = m.db.Menu.DeleteOneID(id).Exec(m.ctx)
	if err != nil {
		return errors.Wrap(err, "delete menu failed")
	}
	return nil
}

func (m Menu) ListByRole(roleID int64) (list []*do.MenuInfoTree, total int64, err error) {

	menus, err := m.db.Role.
		Query().
		Where(role.IDEQ(roleID)).
		QueryMenus().
		Where(menu2.DisabledEQ(0)).
		//WithChildren().
		Order(ent.Asc(menu2.FieldOrderNo)).
		All(m.ctx)

	if err != nil {
		return nil, 0, errors.Wrap(err, "query m by role failed")
	}

	list = findMenuChildren(menus, 1)
	total = int64(len(list))
	return
}

func (m Menu) List(req *do.MenuListReq) (list []*do.MenuInfoTree, total int, err error) {
	// query menu list
	menus, err := m.db.Menu.Query().Order(ent.Asc(menu2.FieldOrderNo)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query menu list failed")
	}
	list = findMenuChildren(menus, 1)
	total, _ = m.db.Menu.Query().Count(m.ctx)
	return
}
func (m Menu) MenuTree(req *do.MenuListReq) (list []*do.Tree, total int, err error) {

	inter, exist := m.cache.Get("MenuTree")
	if exist {
		if v, ok := inter.([]*do.Tree); ok {
			return v, len(v), nil
		}
	}
	menus, err := m.db.Menu.Query().Order(ent.Asc(menu2.FieldOrderNo)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query menu list failed")
	}

	list = findMenuTreeChildren(menus, 1)

	total, _ = m.db.Menu.Query().Count(m.ctx)

	m.cache.SetWithTTL("MenuTree", &list, 1, 30*time.Hour)
	return
}

func (m Menu) CreateMenuParam(req *do.MenuParam) error {
	// check menu whether exist
	exist, err := m.db.Menu.Query().Where(menu2.IDEQ(req.MenuID)).Exist(m.ctx)
	if err != nil {
		return errors.Wrap(err, "query menu failed")
	}
	if !exist {
		return errors.New(fmt.Sprintf("menu not exist, menu id: %d", req.MenuID))
	}

	// create menu param
	err = m.db.MenuParam.Create().
		SetMenusID(req.MenuID).
		SetType(req.Type).
		SetKey(req.Key).
		SetValue(req.Value).
		Exec(m.ctx)
	if err != nil {
		return errors.Wrap(err, "create menu param failed")
	}
	return nil

}

func (m Menu) UpdateMenuParam(req *do.MenuParam) error {
	// check menu whether exist
	exist, err := m.db.Menu.Query().Where(menu2.IDEQ(req.MenuID)).Exist(m.ctx)
	if err != nil {
		return errors.Wrap(err, "query menu failed")
	}
	if !exist {
		return errors.New(fmt.Sprintf("menu not exist, menu id: %d", req.MenuID))
	}

	// update menu param
	err = m.db.MenuParam.UpdateOneID(req.ID).
		SetMenusID(req.MenuID).
		SetType(req.Type).
		SetKey(req.Key).
		SetValue(req.Value).
		Exec(m.ctx)
	if err != nil {
		return errors.Wrap(err, "update menu param failed")
	}
	return nil
}

func (m Menu) DeleteMenuParam(menuParamID int64) error {
	// delete menu param
	err := m.db.MenuParam.DeleteOneID(menuParamID).Exec(m.ctx)
	if err != nil {
		return errors.Wrap(err, "delete menu param failed")
	}
	return nil
}

func (m Menu) MenuParamListByMenuID(menuID int64) (list []do.MenuParam, total int64, err error) {
	// query menu param list
	params, err := m.db.Menu.Query().Where(menu2.IDEQ(menuID)).QueryParams().All(m.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query menu param list failed")
	}

	// convert to MenuParam
	for _, v := range params {
		var p do.MenuParam
		p.ID = v.ID
		p.Type = v.Type
		p.Key = v.Key
		p.Value = v.Value
		p.CreatedAt = v.CreatedAt.Format(time.DateTime)
		p.UpdatedAt = v.UpdatedAt.Format(time.DateTime)
		list = append(list, p)
	}

	total = int64(len(list))
	return
}

func NewMenu(ctx context.Context, c *app.RequestContext) do.Menu {
	return &Menu{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}

func findMenuChildren(data []*ent.Menu, parentID int64) []*do.MenuInfoTree {
	if data == nil {
		return nil
	}
	var result []*do.MenuInfoTree
	for _, v := range data {
		// discard the parent menu, only find the children menu

		if v.ParentID == parentID && v.ID != parentID {
			var m = new(do.MenuInfoTree)
			m.ID = v.ID
			m.Name = v.Name
			m.Key = v.Path
			m.OrderNo = v.OrderNo
			m.Ignore = v.Ignore
			//m.CreatedAt = v.CreatedAt.Format(time.DateTime)
			//m.UpdatedAt = v.UpdatedAt.Format(time.DateTime)
			//m.MenuType = v.MenuType
			//m.Level = v.MenuLevel

			//	Title:              v.Name,
			//m.ParentID = v.ParentID
			//m.Path = v.Path
			//m.Redirect = v.Redirect
			//m.Component = v.Component

			//m.Meta = &do.MenuMeta{
			//	Icon:               v.Icon,
			//	HideMenu:           v.HideMenu,
			//	HideBreadcrumb:     v.HideBreadcrumb,
			//	CurrentActiveMenu:  v.CurrentActiveMenu,
			//	IgnoreKeepAlive:    v.IgnoreKeepAlive,
			//	HideTab:            v.HideTab,
			//	FrameSrc:           v.FrameSrc,
			//	CarryParam:         v.CarryParam,
			//	HideChildrenInMenu: v.HideChildrenInMenu,
			//	Affix:              v.Affix,
			//	DynamicLevel:       v.DynamicLevel,
			//	RealPath:           v.RealPath,
			//}

			m.Children = findMenuChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}

func findMenuTreeChildren(data []*ent.Menu, parentID int64) []*do.Tree {
	if data == nil {
		return nil
	}
	var result []*do.Tree
	for _, v := range data {
		if v.ParentID == parentID && v.ID != parentID {
			var m = new(do.Tree)
			m.Title = v.Name
			m.Value = v.ID
			m.Key = v.ID
			m.Children = findMenuTreeChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}
