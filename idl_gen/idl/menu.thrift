namespace go menu

include "../base/base.thrift"

// menu service
service MenuService {

  // 获取角色菜单权限列表
 base.NilResponse MenuAuth(1: base.IDReq req) (api.post = "/service/auth/menu/role")
  //获取角色菜单列表
 base.NilResponse MenuRole() (api.post = "/service/menu/role")
  // 创建或API
  base.NilResponse CreateApi(1: ApiInfo req) (api.post = "/service/api/create")

  // 更新API
  base.NilResponse UpdateApi(1: ApiInfo req) (api.post = "/service/api/update")

  // 删除API信息
  base.NilResponse DeleteApi(1: base.IDReq req) (api.post = "/service/api")

  // 获取API列表
  base.NilResponse ApiList(1: ApiPageReq req) (api.post = "/service/api/list")

  base.NilResponse ApiTree(1: ApiPageReq req) (api.post = "/service/api/tree")

  // 创建菜单
  base.NilResponse CreateMenu(1: CreateOrUpdateMenuReq req) (api.post = "/service/menu/create")

  //更新菜单
  base.NilResponse UpdateMenu(1: CreateOrUpdateMenuReq req) (api.post = "/service/menu/update")

  //删除菜单信息
  base.NilResponse DeleteMenu(1: base.IDReq req) (api.post = "/service/menu")


  //获取菜单列表
  base.NilResponse MenuLists(1: base.PageInfoReq req) (api.post = "/service/menu/list")

  base.NilResponse MenuTree(1: base.PageInfoReq req) (api.post = "/service/menu/tree")

//  //创建菜单额外参数
//  base.NilResponse CreateMenuParam(1: CreateOrUpdateMenuParamReq req) (api.post = "/service/menu/param/create")
//
//  //更新菜单额外参数
//  base.NilResponse UpdateMenuParam(1: CreateOrUpdateMenuParamReq req) (api.post = "/service/menu/param/update")
//
//  //删除菜单额外参数
//  base.NilResponse DeleteMenuParam(1: base.IDReq req) (api.post = "/service/menu/param")
//
//  //获取某个菜单的额外参数列表
//  base.NilResponse MenuParamListByMenuID(1: base.IDReq req) (api.post = "/service/menu/param/list")

}

// API信息
struct ApiInfo {
    1:  i64 id (api.raw = "id")
    2:  string createdAt (api.raw = "createdAt")
    3:  string updatedAt (api.raw = "updatedAt")
    4:  string path (api.raw = "path")
    5:  string description (api.raw = "description")
    6:  string group (api.raw = "group")
    7:  string method (api.raw = "method")
}

// API列表请求数据
struct ApiPageReq {
    1:  i64 page (api.raw = "page")
    2:  i64 pageSize (api.raw = "pageSize")
    3:  string path (api.raw = "path")
    4:  string description (api.raw = "description")
    5:  string method (api.raw = "method")
    6:  string group (api.raw = "group")
}

//菜单的meta数据
struct Meta {
    1:  string title (api.raw = "title" )
    2:  string icon (api.raw = "icon" )
    3:  string activeMenu (api.raw = "activeMenu" )
    4:  bool affix (api.raw = "affix" )
    5:  bool noCache (api.raw = "noCache" )
}

// 创建或更新菜单信息参数
struct CreateOrUpdateMenuReq {
    1:  i64 id (api.raw = "id" )
    2:  string name (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:  i64 parentId (api.raw = "parentId")
    4:  i64 level (api.raw = "level")
    5:  string path (api.raw = "path")
    6:  string redirect (api.raw = "redirect")
    7:  string component (api.raw = "component")
    8:  i64 menuType (api.raw = "menuType")
    9:  bool hidden (api.raw = "hidden")
    10:  i64 sort (api.raw = "sort")
    12:  i64 status (api.raw = "status")
    13:  string url (api.raw = "url")

    //meta
       14:  string title (api.raw = "title" )
       15:  string icon (api.raw = "icon" )
       16:  string activeMenu (api.raw = "activeMenu" )
       17:  bool affix (api.raw = "affix" )
       18:  bool noCache (api.raw = "noCache" )


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
	    1:  i64 id (api.raw = "id" )
        2:  string name (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
        3:  i64 parentId (api.raw = "parentId")
        4:  i64 level (api.raw = "level")
        5:  string path (api.raw = "path")
        6:  string redirect (api.raw = "redirect")
        7:  string component (api.raw = "component")
        8:  i64 menuType (api.raw = "menuType")
        9:  bool hidden (api.raw = "hidden")
        10:  i64 sort (api.raw = "sort")
        11:  Meta meta (api.raw = "meta")
        12:  i64 status (api.raw = "status")
        13:  string url (api.raw = "url")
	    14: list<MenuInfo> Children  (api.raw = "children")

}
