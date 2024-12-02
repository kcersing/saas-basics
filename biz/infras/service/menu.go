package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/menu"
	"saas/pkg/db/ent"
	menu2 "saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/role"
	"strconv"
	"time"
)

type Menu struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m Menu) MenuInfo(id int64) (info *menu.MenuInfo, err error) {
	//TODO implement me
	info = new(menu.MenuInfo)

	inter, exist := m.cache.Get("menuInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := inter.(*menu.MenuInfo); ok {
			return u, nil
		}
	}
	menuEnt, err := m.db.Menu.Query().Where(menu2.IDEQ(id)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member failed")
		return info, err
	}

	var meta = new(menu.Meta)

	info.Meta = meta
	info.ParentId = menuEnt.ParentID
	info.Path = menuEnt.Path
	info.ID = menuEnt.ID
	info.Name = menuEnt.Name
	info.Level = menuEnt.Level
	info.MenuType = menuEnt.MenuType
	info.Redirect = menuEnt.Redirect
	info.Component = menuEnt.Component
	info.Hidden = menuEnt.Hidden
	info.URL = menuEnt.URL
	info.Status = menuEnt.Status
	info.Sort = menuEnt.Sort

	info.CreatedAt = menuEnt.CreatedAt.Format(time.DateTime)
	info.UpdatedAt = menuEnt.UpdatedAt.Format(time.DateTime)
	info.Meta = &menu.Meta{
		Icon:       menuEnt.Icon,
		Title:      menuEnt.Title,
		ActiveMenu: menuEnt.ActiveMenu,
		NoCache:    menuEnt.NoCache,
		Affix:      menuEnt.Affix,
	}
	m.cache.SetWithTTL("menuInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
	return

}

func NewMenu(ctx context.Context, c *app.RequestContext) do.Menu {
	return &Menu{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
func (m Menu) Create(menuReq *menu.CreateOrUpdateMenuReq) error {
	// get menu level
	if menuReq.ParentId == 0 {
		// it is a first level menu
		menuReq.ParentId = 1
		menuReq.Level = 1
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
		SetParentID(menuReq.ParentId).
		SetPath(menuReq.Path).
		SetName(menuReq.Name).
		SetLevel(menuReq.Level).
		SetMenuType(menuReq.MenuType).
		SetRedirect(menuReq.Redirect).
		SetComponent(menuReq.Component).
		SetHidden(menuReq.Hidden).
		SetURL(menuReq.URL).
		SetStatus(menuReq.Status).
		SetSort(menuReq.Sort).
		// meta
		SetTitle(menuReq.Title).
		SetIcon(menuReq.Icon).
		SetActiveMenu(menuReq.ActiveMenu).
		SetAffix(menuReq.Affix).
		SetNoCache(menuReq.NoCache).
		Exec(m.ctx)

	if err != nil {
		return errors.Wrap(err, "create menu failed")
	}
	return nil
}

func (m Menu) Update(menuReq *menu.CreateOrUpdateMenuReq) error {
	// get menu level
	if menuReq.ParentId == 0 {
		// it is a first level menu
		menuReq.ParentId = 1
		menuReq.Level = 1
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
		SetParentID(menuReq.ParentId).
		SetPath(menuReq.Path).
		SetName(menuReq.Name).
		SetLevel(menuReq.Level).
		SetMenuType(menuReq.MenuType).
		SetRedirect(menuReq.Redirect).
		SetComponent(menuReq.Component).
		SetHidden(menuReq.Hidden).
		SetURL(menuReq.URL).
		SetStatus(menuReq.Status).
		// meta
		SetTitle(menuReq.Title).
		SetIcon(menuReq.Icon).
		SetActiveMenu(menuReq.ActiveMenu).
		SetAffix(menuReq.Affix).
		SetNoCache(menuReq.NoCache).
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

func (m Menu) MenuRole(roleID int64) (list []*menu.MenuInfo, err error) {

	menus, err := m.db.Role.
		Query().
		Where(role.IDEQ(roleID)).
		QueryMenus().
		Where(menu2.DisabledEQ(0)).
		//WithChildren().
		Order(ent.Asc(menu2.FieldSort)).
		All(m.ctx)

	if err != nil {
		return nil, errors.Wrap(err, "query m by role failed")
	}

	list = findMenuChildren(menus, 1)

	return
}

func (m Menu) List(req *base.PageInfoReq) (list []*menu.MenuInfo, total int, err error) {
	// query menu list
	menus, err := m.db.Menu.Query().Order(ent.Asc(menu2.FieldSort)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query menu list failed")
	}
	list = findMenuChildren(menus, 1)
	total, _ = m.db.Menu.Query().Count(m.ctx)
	return
}
func (m Menu) MenuTree(req *base.PageInfoReq) (list []*base.Tree, err error) {

	inter, exist := m.cache.Get("MenuTree")
	if exist {
		if v, ok := inter.([]*base.Tree); ok {
			return v, nil
		}
	}
	menus, err := m.db.Menu.Query().Order(ent.Asc(menu2.FieldSort)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query menu list failed")
	}

	list = findMenuTreeChildren(menus, 1)

	m.cache.SetWithTTL("MenuTree", &list, 1, 30*time.Hour)
	return
}

func (m Menu) CreateMenuParam(req *menu.MenuParam) error {
	// check menu whether exist
	exist, err := m.db.Menu.Query().Where(menu2.IDEQ(req.MenuId)).Exist(m.ctx)
	if err != nil {
		return errors.Wrap(err, "query menu failed")
	}
	if !exist {
		return errors.New(fmt.Sprintf("menu not exist, menu id: %d", req.MenuId))
	}

	// create menu param
	err = m.db.MenuParam.Create().
		SetMenusID(req.MenuId).
		SetType(req.Type).
		SetKey(req.Key).
		SetValue(req.Value).
		Exec(m.ctx)
	if err != nil {
		return errors.Wrap(err, "create menu param failed")
	}
	return nil

}

func (m Menu) UpdateMenuParam(req *menu.MenuParam) error {
	// check menu whether exist
	exist, err := m.db.Menu.Query().Where(menu2.IDEQ(req.MenuId)).Exist(m.ctx)
	if err != nil {
		return errors.Wrap(err, "query menu failed")
	}
	if !exist {
		return errors.New(fmt.Sprintf("menu not exist, menu id: %d", req.MenuId))
	}

	// update menu param
	err = m.db.MenuParam.UpdateOneID(req.ID).
		SetMenusID(req.MenuId).
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

func (m Menu) MenuParamListByMenuID(menuID int64) (list []menu.MenuParam, total int64, err error) {
	// query menu param list
	params, err := m.db.Menu.Query().Where(menu2.IDEQ(menuID)).QueryParams().All(m.ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "query menu param list failed")
	}

	// convert to MenuParam
	for _, v := range params {
		var p menu.MenuParam
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

func findMenuChildren(data []*ent.Menu, parentID int64) []*menu.MenuInfo {
	if data == nil {
		return nil
	}
	var result []*menu.MenuInfo
	for _, v := range data {
		// discard the parent menu, only find the children menu

		if v.ParentID == parentID && v.ID != parentID {
			var m = new(menu.MenuInfo)
			var meta = new(menu.Meta)

			m.Meta = meta
			m.ParentId = v.ParentID
			m.Path = v.Path
			m.ID = v.ID
			m.Name = v.Name
			m.Level = v.Level
			m.MenuType = v.MenuType
			m.Redirect = v.Redirect
			m.Component = v.Component
			m.Hidden = v.Hidden
			m.URL = v.URL
			m.Status = v.Status
			m.Sort = v.Sort

			m.CreatedAt = v.CreatedAt.Format(time.DateTime)
			m.UpdatedAt = v.UpdatedAt.Format(time.DateTime)
			m.Meta = &menu.Meta{
				Icon:       v.Icon,
				Title:      v.Title,
				ActiveMenu: v.ActiveMenu,
				NoCache:    v.NoCache,
				Affix:      v.Affix,
			}

			m.Children = findMenuChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}

func findMenuTreeChildren(data []*ent.Menu, parentID int64) []*base.Tree {
	if data == nil {
		return nil
	}
	var result []*base.Tree
	for _, v := range data {
		if v.ParentID == parentID && v.ID != parentID {
			var m = new(base.Tree)
			m.Title = v.Title
			m.Value = strconv.FormatInt(v.ID, 10)
			m.Key = strconv.FormatInt(v.ID, 10)
			m.Children = findMenuTreeChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}
