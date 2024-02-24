package do

type Menu interface {
	Create(menuReq *MenuInfo) error
	Update(menuReq *MenuInfo) error
	Delete(id uint64) error
	ListByRole(roleID uint64) (list []*MenuInfoTree, total uint64, err error)
	List(req *MenuListReq) (list []*MenuInfoTree, total int, err error)
	CreateMenuParam(req *MenuParam) error
	UpdateMenuParam(req *MenuParam) error
	DeleteMenuParam(menuParamID uint64) error
	MenuParamListByMenuID(menuID uint64) (list []MenuParam, total uint64, err error)
}

type MenuInfo struct {
	Level     uint32      `json:"level,omitempty"`
	ParentID  uint64      `json:"parentId,omitempty"`
	Path      string      `json:"path,omitempty"`
	Name      string      `json:"name,omitempty"`
	Redirect  string      `json:"redirect,omitempty"`
	Component string      `json:"component,omitempty"`
	OrderNo   uint32      `json:"orderNo,omitempty"`
	Disabled  bool        `json:"disabled,omitempty"`
	Meta      *MenuMeta   `json:"meta,omitempty"`
	ID        uint64      `json:"id,omitempty"`
	MenuType  uint32      `json:"menuType,omitempty"`
	Children  []*MenuInfo `json:"children,omitempty"`
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
	DynamicLevel       uint32 `json:"dynamicLevel,omitempty"`
	RealPath           string `json:"realPath,omitempty"`
}

type MenuListReq struct {
	Page     uint64 `json:"page,omitempty"`
	PageSize uint64 `json:"pageSize,omitempty"`
}

type MenuInfoTree struct {
	MenuInfo  `json:"menu_info"`
	CreatedAt string          `json:"createdAt,omitempty"`
	UpdatedAt string          `json:"updatedAt,omitempty"`
	Children  []*MenuInfoTree `json:"children,omitempty"`
}

// MenuParam is the menu parameter structure.data stored at the table `sys_menu_params`
type MenuParam struct {
	ID        uint64 `json:"id,omitempty"`
	MenuID    uint64 `json:"menuId,omitempty"`
	Type      string `json:"type,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
