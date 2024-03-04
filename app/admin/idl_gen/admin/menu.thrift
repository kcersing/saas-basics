namespace go admin.menu

include "../base/base.thrift"
include "../base/data.thrift"


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
    1:  string id (api.raw = "id" )
    2:  string name (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:  string parent_id (api.raw = "parent_id")
    4:  string level (api.raw = "level")
    5:  string path (api.raw = "path")
    6:  string redirect (api.raw = "redirect")
    7:  string component (api.raw = "component")
    8:  string orderNo (api.raw = "orderNo")
    9:  string disabled (api.raw = "disabled")
    10:  string menuType (api.raw = "menuType")
    11:  Meta meta (api.raw = "meta")
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
  base.NilResponse MenuLists(1: base.PageInfoReq req) (api.get = "/api/admin/menu/list")

  //创建菜单额外参数
  base.NilResponse CreateMenuParam(1: CreateOrUpdateMenuParamReq req) (api.post = "/api/admin/menu/param/create")

  //更新菜单额外参数
  base.NilResponse UpdateMenuParam(1: CreateOrUpdateMenuParamReq req) (api.post = "/api/admin/menu/param/update")

  //删除菜单额外参数
  base.NilResponse DeleteMenuParam(1: base.IDReq req) (api.post = "/api/admin/menu/param")

  //获取某个菜单的额外参数列表
  base.NilResponse MenuParamListByMenuID(1: base.IDReq req) (api.post = "/api/admin/menu/param/list")

}
