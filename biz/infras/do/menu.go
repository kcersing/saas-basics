package do

import (
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/menu"
)

type Menu interface {
	Create(menuReq *menu.CreateOrUpdateMenuReq) error
	Update(menuReq *menu.CreateOrUpdateMenuReq) error
	Delete(id int64) error
	MenuRole(roleID int64) (list []*menu.MenuInfo, err error)
	List(req *base.PageInfoReq) (list []*menu.MenuInfo, total int, err error)

	MenuInfo(id int64) (info *menu.MenuInfo, err error)

	MenuTree(req *base.PageInfoReq) (list []*base.Tree, err error)

	CreateMenuParam(req *menu.MenuParam) error
	UpdateMenuParam(req *menu.MenuParam) error
	DeleteMenuParam(menuParamID int64) error
	MenuParamListByMenuID(menuID int64) (list []menu.MenuParam, total int64, err error)
}
