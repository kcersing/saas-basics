namespace go system.menu

include "../base/base.thrift"



//菜单的meta数据
struct Meta {
    1:  string title (api.raw = "title" )
    2:  string icon (api.raw = "icon" )
    3:  string hideMenu (api.raw = "hideMenu" )
    4:  string hideBreadcrumb (api.raw = "hideBreadcrumb" )
    5:  string currentActiveMenu (api.raw = "currentActiveMenu" )
    6:  string ignoreKeepAlive (api.raw = "ignoreKeepAlive" )
    7:  string hideTab (api.raw = "hideTab" )
    8:  string frameSrc (api.raw = "frameSrc" )
    9:  string carryParam (api.raw = "carryParam" )
    10:  string hideChildrenInMenu (api.raw = "hideChildrenInMenu" )
    11:  string affix (api.raw = "affix" )
    12:  string dynamicLevel (api.raw = "dynamicLevel" )
    13:  string realPath (api.raw = "realPath" )
}


// 创建或更新菜单信息参数
struct CreateOrUpdateMenuReq {
    1:  i64 id (api.raw = "id" )
    2:  string name (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:  i64 parent_id (api.raw = "parentID")
//    4:  i64 level (api.raw = "level")
    5:  string path (api.raw = "path")
//    6:  string redirect (api.raw = "redirect")
//    7:  string component (api.raw = "component")
    8:  i64 orderNo (api.raw = "orderNo")
    9:  bool disabled (api.raw = "disabled")
//    10:  string menuType (api.raw = "menuType")
//    11:  Meta meta (api.raw = "meta")
}

//更新菜单额外参数
struct CreateOrUpdateMenuParamReq{
    1:  string ID (api.raw = "ID")
    2:  string menuID (api.raw = "menuID")
    3:  string type (api.raw = "type")
    4:  string key (api.raw = "key")
    5:  string value (api.raw = "value")
}

struct MenuInfoTree {
    1: MenuInfo MenuInfo;
    2: string CreatedAt;
    3: string UpdatedAt;
    4: list<MenuInfoTree> Children;
    5: bool Ignore;
}

struct MenuInfo {
	1: i64 ID
	2: i64 ParentID
	3: string Path
	4: string Name
	5: string Key
	6: list<MenuInfo> Children
	7: i64 OrderNo
	8: i64 Disabled
	9: bool Ignore
}

struct Tree {
	1:string Title
    2:string Value
	3:string Key
	4:string Method
	5:list<Tree> Children
}


// menu service
service MenuService {

  // 创建菜单
  base.NilResponse CreateMenu(1: CreateOrUpdateMenuReq req) (api.post = "/api/admin/menu/create")

  //更新菜单
  base.NilResponse UpdateMenu(1: CreateOrUpdateMenuReq req) (api.post = "/api/admin/menu/update")

  //删除菜单信息
  base.NilResponse DeleteMenu(1: base.IDReq req) (api.post = "/api/admin/menu")

  //获取角色菜单列表
  list<MenuInfoTree> MenuByRole(1: base.IDReq req) (api.post = "/api/admin/menu/role")

  //获取菜单列表
  base.NilResponse MenuLists(1: base.PageInfoReq req) (api.post = "/api/admin/menu/list")

  list<Tree> MenuTree(1: base.PageInfoReq req) (api.post = "/api/admin/menu/tree")

  //创建菜单额外参数
  base.NilResponse CreateMenuParam(1: CreateOrUpdateMenuParamReq req) (api.post = "/api/admin/menu/param/create")

  //更新菜单额外参数
  base.NilResponse UpdateMenuParam(1: CreateOrUpdateMenuParamReq req) (api.post = "/api/admin/menu/param/update")

  //删除菜单额外参数
  base.NilResponse DeleteMenuParam(1: base.IDReq req) (api.post = "/api/admin/menu/param")

  //获取某个菜单的额外参数列表
  base.NilResponse MenuParamListByMenuID(1: base.IDReq req) (api.post = "/api/admin/menu/param/list")

}