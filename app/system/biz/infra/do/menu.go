package do

import (
	"rpc_gen/kitex_gen/base"
	"rpc_gen/kitex_gen/system/menu"
)

type Menu interface {
	Create(menuReq *menu.CreateOrUpdateMenuReq) error
	Update(menuReq *menu.CreateOrUpdateMenuReq) error
	Delete(id int64) error
	MenuRole(roleID int64) (list []*menu.MenuInfoTree, err error)
	List(req *base.PageInfoReq) (list []*menu.MenuInfoTree, total int, err error)

	MenuTree(req *base.PageInfoReq) (list []*base.Tree, err error)

	CreateMenuParam(req *MenuParam) error
	UpdateMenuParam(req *MenuParam) error
	DeleteMenuParam(menuParamID int64) error
	MenuParamListByMenuID(menuID int64) (list []MenuParam, total int64, err error)
}

type MenuInfoTree struct {
	MenuInfo  `json:"menu_info"`
	CreatedAt string          `json:"createdAt,omitempty"`
	UpdatedAt string          `json:"updatedAt,omitempty"`
	Children  []*MenuInfoTree `json:"children,omitempty"`
	Ignore    bool            `json:"ignore,omitempty"`
}
type MenuInfo struct {
	ID       int64       `json:"id,omitempty"`
	ParentID int64       `json:"parentId,omitempty"`
	Path     string      `json:"path,omitempty"`
	Name     string      `json:"name,omitempty"`
	Key      string      `json:"key,omitempty"`
	Children []*MenuInfo `json:"children,omitempty"`
	OrderNo  int32       `json:"orderNo,omitempty"`
	Disabled int32       `json:"disabled,omitempty"`
	Ignore   bool        `json:"ignore,omitempty"`
	//Level     int32       `json:"level,omitempty"`
	//Redirect  string      `json:"redirect,omitempty"`
	//Component string      `json:"component,omitempty"`
	//Meta      *MenuMeta   `json:"meta,omitempty"`
	//MenuType  int32       `json:"menuType,omitempty"`
}

type MenuMeta struct {
	Title              string `json:"title,omitempty"`
	Icon               string `json:"icon,omitempty"`
	HideMenu           bool   `json:"hideMenu,omitempty"`
	HideBreadcrumb     bool   `json:"hideBreadcrumb,omitempty"`
	CurrentActiveMenu  string `json:"currentActiveMenu,omitempty"`
	IgnoreKeepAlive    bool   `json:"ignoreKeepAlive,omitempty"`
	HideTab            bool   `json:"hideTab,omitempty"`
	FrameSrc           string `json:"frameSrc,omitempty"`
	CarryParam         bool   `json:"carryParam,omitempty"`
	HideChildrenInMenu bool   `json:"hideChildrenInMenu,omitempty"`
	Affix              bool   `json:"affix,omitempty"`
	DynamicLevel       int32  `json:"dynamicLevel,omitempty"`
	RealPath           string `json:"realPath,omitempty"`
}

type MenuListReq struct {
	Page     int64 `json:"page,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

// MenuParam is the menu parameter structure.data stored at the table `sys_menu_params`
type MenuParam struct {
	ID        int64  `json:"id,omitempty"`
	MenuID    int64  `json:"menuId,omitempty"`
	Type      string `json:"type,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
