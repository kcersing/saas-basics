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
	Level     uint32
	ParentID  uint64
	Path      string
	Name      string
	Redirect  string
	Component string
	OrderNo   uint32
	Disabled  bool
	Meta      *MenuMeta
	ID        uint64
	MenuType  uint32
}

type MenuMeta struct {
	Title              string
	Icon               string
	HideMenu           bool
	HideBreadcrumb     bool
	CurrentActiveMenu  string
	IgnoreKeepAlive    bool
	HideTab            bool
	FrameSrc           string
	CarryParam         bool
	HideChildrenInMenu bool
	Affix              bool
	DynamicLevel       uint32
	RealPath           string
}

type MenuListReq struct {
	Page     uint64
	PageSize uint64
}

type MenuInfoTree struct {
	MenuInfo
	CreatedAt string
	UpdatedAt string
	Children  []*MenuInfoTree
}

// MenuParam is the menu parameter structure.data stored at the table `sys_menu_params`
type MenuParam struct {
	ID        uint64
	MenuID    uint64
	Type      string
	Key       string
	Value     string
	CreatedAt string
	UpdatedAt string
}
