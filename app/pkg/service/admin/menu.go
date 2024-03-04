package admin

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	infras "saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	menu2 "saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/role"
)

type Menu struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
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
	menus, err := m.db.Role.Query().Where(role.IDEQ(roleID)).
		QueryMenus().
		//WithChildren().
		Order(ent.Asc(menu2.FieldOrderNo)).All(m.ctx)

	if err != nil {
		return nil, 0, errors.Wrap(err, "query m by role failed")
	}

	list = findMenuChildren(menus, 1)

	total = uint64(len(list))
	return

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

func (m Menu) DeleteMenuParam(menuParamID uint64) error {
	// delete menu param
	err := m.db.MenuParam.DeleteOneID(menuParamID).Exec(m.ctx)
	if err != nil {
		return errors.Wrap(err, "delete menu param failed")
	}
	return nil
}

func (m Menu) MenuParamListByMenuID(menuID uint64) (list []do.MenuParam, total uint64, err error) {
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
		p.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
		p.UpdatedAt = v.UpdatedAt.Format("2006-01-02 15:04:05")
		list = append(list, p)
	}

	total = uint64(len(list))
	return
}

func NewMenu(ctx context.Context, c *app.RequestContext) do.Menu {
	return &Menu{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}

func findMenuChildren(data []*ent.Menu, parentID uint64) []*do.MenuInfoTree {
	if data == nil {
		return nil
	}
	var result []*do.MenuInfoTree
	for _, v := range data {
		// discard the parent menu, only find the children menu

		if v.ParentID == parentID && v.ID != parentID {
			var m = new(do.MenuInfoTree)
			m.ID = v.ID
			m.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
			m.UpdatedAt = v.UpdatedAt.Format("2006-01-02 15:04:05")
			m.MenuType = v.MenuType
			m.Level = v.MenuLevel
			m.ParentID = v.ParentID
			m.Path = v.Path
			m.Name = v.Name
			m.Redirect = v.Redirect
			m.Component = v.Component
			m.OrderNo = v.OrderNo
			m.Meta = &do.MenuMeta{
				Title:              v.Title,
				Icon:               v.Icon,
				HideMenu:           v.HideMenu,
				HideBreadcrumb:     v.HideBreadcrumb,
				CurrentActiveMenu:  v.CurrentActiveMenu,
				IgnoreKeepAlive:    v.IgnoreKeepAlive,
				HideTab:            v.HideTab,
				FrameSrc:           v.FrameSrc,
				CarryParam:         v.CarryParam,
				HideChildrenInMenu: v.HideChildrenInMenu,
				Affix:              v.Affix,
				DynamicLevel:       v.DynamicLevel,
				RealPath:           v.RealPath,
			}

			m.Children = findMenuChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}
