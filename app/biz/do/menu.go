package do

import (
	"saas/app/biz/model/admin/menu"
	"saas/app/biz/model/base"
	"saas/pkg/db/ent"
)

type Menu interface {
	CreateMenu(req menu.CreateOrUpdateMenuReq) error
	UpdateMenu(req menu.CreateOrUpdateMenuReq) error
	DeleteMenu(id string) error
	MenuByRole(id string) error
	MenuList(req base.PageInfoReq) []*ent.Menu
	CreateMenuParam(req menu.CreateOrUpdateMenuParamReq) error
	UpdateMenuParam(req menu.CreateOrUpdateMenuParamReq) error
	DeleteMenuParam(id string) error
	MenuParamListByMenuID(id string) error
}
