namespace go admin.menu

include "../base/base.thrift"
include "../base/data.thrift"


// 创建或更新菜单信息参数
struct CreateOrUpdateMenuReq {
    1:  string id (api.raw = "id" )
    2:  string name (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:  string menu_type (api.raw = "menu_type")
    4:  string parent_id (api.raw = "parent_id")
    6:  string level (api.raw = "level")
    7:  string route_path (api.raw = "route_path")
    8:  string route_name (api.raw = "route_name")
    9:  string component (api.raw = "component")
    10:  string Status (api.raw = "status")
    11:  string icon_type (api.raw = "icon_type")
    12:  string icon (api.raw = "icon")
    13:  string i18n_key (api.raw = "i18n_key")
    14:  string redirect (api.raw = "redirect")
}

//更新菜单额外参数
struct CreateOrUpdateMenuParamReq{
    1:  string ID (api.raw = "ID")
    2:  string menuID (api.raw = "menuID")
    3:  string type (api.raw = "type")
    4:  string key (api.raw = "key")
    5:  string value (api.raw = "value")
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
  base.NilResponse MenuByRole(1: base.Empty req) (api.get = "/api/admin/menu/role")

  //获取菜单列表
  base.NilResponse MenuLists(1: base.PageInfoReq req) (api.get = "/api/admin/menu/lists")

  //创建菜单额外参数
  base.NilResponse CreateMenuParam(1: CreateOrUpdateMenuParamReq req) (api.post = "/api/admin/menu/param/create")

  //更新菜单额外参数
  base.NilResponse UpdateMenuParam(1: CreateOrUpdateMenuParamReq req) (api.post = "/api/admin/menu/param/update")

  //删除菜单额外参数
  base.NilResponse DeleteMenuParam(1: base.IDReq req) (api.post = "/api/admin/menu/param")

  //获取某个菜单的额外参数列表
  base.NilResponse MenuParamListByMenuID(1: base.IDReq req) (api.post = "/api/admin/menu/param/list")

}
