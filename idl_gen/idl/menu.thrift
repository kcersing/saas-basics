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
  base.NilResponse MenuLists(1: MenuListReq req) (api.post = "/service/menu/list")

  base.NilResponse MenuTree(1: MenuListReq req) (api.post = "/service/menu/tree")


  // 获取用户基本信息
  base.NilResponse MenuInfo(1: base.IDReq req) (api.post = "/service/menu/info")

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
struct MenuListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    19:optional string type="" (api.raw = "type")
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
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  string path = "" (api.raw = "path")
    4:  string description = "" (api.raw = "description")
    5:  string method  = ""(api.raw = "method")
    6:  string group  = ""(api.raw = "group")
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
    1:optional i64 id =0(api.raw = "id" )
    2:optional string name="" (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:optional i64 parentId=0 (api.raw = "parentId")
    4:optional i64 level=0 (api.raw = "level")
    5:optional string path="" (api.raw = "path")
    6:optional string redirect="" (api.raw = "redirect")
    7:optional string component="" (api.raw = "component")
    8:optional i64 menuType=0 (api.raw = "menuType")
    9:optional bool hidden=true (api.raw = "hidden")
    10:optional i64 sort=0 (api.raw = "sort")
    12:optional i64 status=1 (api.raw = "status")
    13:optional string url="" (api.raw = "url")

    //meta
       14:optional string title="" (api.raw = "title" )
       15:optional string icon="" (api.raw = "icon" )
       16:optional string activeMenu="" (api.raw = "activeMenu" )
       17:optional bool affix=true (api.raw = "affix" )
       18:optional bool noCache=true (api.raw = "noCache" )

   19:optional string type="" (api.raw = "type")
}

//更新菜单额外参数
struct CreateOrUpdateMenuParamReq{
    1:  string id (api.raw = "ID")
    2:  string menuID (api.raw = "menuID")
    3:  string type (api.raw = "type")
    4:  string key (api.raw = "key")
    5:  string value (api.raw = "value")
}

struct MenuInfoTree {
    1: MenuInfo menuInfo;
    2: string createdAt;
    3: string updatedAt;
    4: list<MenuInfoTree> children;
    5: bool ignore;
}

struct MenuInfo {
	    1: i64 id (api.raw = "id" )
        2: string name (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
        3: i64 parentId (api.raw = "parentId")
        4: i64 level (api.raw = "level")
        5: string path (api.raw = "path")
        6: string redirect (api.raw = "redirect")
        7: string component (api.raw = "component")
        8: i64 menuType (api.raw = "menuType")
        9: bool hidden (api.raw = "hidden")
        10: i64 sort (api.raw = "sort")
        11: Meta meta (api.raw = "meta")
        12: i64 status (api.raw = "status")
        13: string url (api.raw = "url")
	    14: list<MenuInfo> children  (api.raw = "children")
	    15: string createdAt (api.raw = "createdAt")
        16: string updatedAt (api.raw = "updatedAt")
        17:  string title (api.raw = "title" )
        19:optional string type="" (api.raw = "type")
}

// MenuParam is the menu parameter structure.data stored at the table `sys_menu_params`
struct MenuParam  {
	1: i64 id       (api.raw = "id")
	2: i64 menuId    (api.raw = "menuId")
    3: string type   (api.raw = "type")
    4: string key    (api.raw = "key")
    5: string value   (api.raw = "value")
    6: string createdAt (api.raw = "createdAt")
    7: string updatedAt (api.raw = "updatedAt")
}
