namespace go admin.role

include "../base/base.thrift"
include "../base/data.thrift"

//创建或更新角色信息参数
struct RoleInfo {
    1:  string ID (api.raw = "ID")
    2:  string name (api.raw = "name")
    3:  string value (api.raw = "value")
    4:  string defaultRouter (api.raw = "defaultRouter")
    5:  string status (api.raw = "status")
    6:  string remark  (api.raw = "remark")
    7:  string orderNo (api.raw = "orderNo")
    8:  string createdAt (api.raw = "createdAt")
    9:  string updatedAt (api.raw = "updatedAt")
}

// role service
service RoleService {

  // Create role information | 创建角色
  base.NilResponse CreateRole(1: RoleInfo req) (api.post = "/api/admin/role/create")

  // Update role information | 更新角色
  base.NilResponse UpdateRole(1: RoleInfo req) (api.post = "/api/admin/role/update")

  // Delete role information | 删除角色信息
  base.NilResponse DeleteRole(1: base.IDReq req) (api.post = "/api/admin/role")

  // Get role information | 获取角色信息
  base.NilResponse RoleByID(1: base.IDReq req) (api.get = "/api/admin/role")

  // Get role list | 获取角色列表
  base.NilResponse RoleList(1: base.PageInfoReq req) (api.get = "/api/admin/role/list")

  // Set role status | 设置角色状态, 启用1/禁用0
  base.NilResponse UpdateRoleStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/role/status")
}

// authorization message
// API授权数据
struct ApiAuthorityInfo {
    1:  string path (api.raw = "path")
    2:  string method (api.raw = "method")
}

// 创建或更新API授权信息
struct CreateOrUpdateApiAuthorityReq {
    1:  string role_id (api.raw = "role_id")
    2:  ApiAuthorityInfo apiAuthorityInfo (api.raw = "api_authority_info")
}

// 菜单授权请求数据
struct MenuAuthorityInfoReq {
    1:  string role_id (api.raw = "role_id")
    2:  string menu_ids (api.raw = "menu_ids")
}

// authorization service
service authorityService {

  // 创建API权限
  base.NilResponse CreateAuthority(1: CreateOrUpdateApiAuthorityReq req) (api.post = "/api/admin/authority/api/create")

  // 更新API权限
  base.NilResponse UpdateApiAuthority(1: CreateOrUpdateApiAuthorityReq req) (api.post = "/api/admin/authority/api/update")

  // 获取角色api权限列表
  base.NilResponse ApiAuthority(1: base.IDReq req) (api.post = "/api/admin/authority/api/role")

  // 创建菜单权限
  base.NilResponse CreateMenuAuthority(1: MenuAuthorityInfoReq req) (api.post = "/api/admin/authority/menu/create")

  // 更新菜单权限
  base.NilResponse UpdateMenuAuthority(1: MenuAuthorityInfoReq req) (api.post = "/api/admin/authority/menu/update")

  // 获取角色菜单权限列表
  base.NilResponse MenuAuthority(1: base.IDReq req) (api.post = "/api/admin/authority/menu/role")

}

// api message
// API信息
struct ApiInfo {
    1:  string id (api.raw = "id")
    2:  string createdAt (api.raw = "createdAt")
    3:  string updatedAt (api.raw = "updatedAt")
    4:  string path (api.raw = "path")
    5:  string description (api.raw = "description")
    6:  string group (api.raw = "group")
    7:  string method (api.raw = "method")
}

// API列表请求数据
struct ApiPageReq {
    1:  string page (api.raw = "page")
    2:  string limit (api.raw = "limit")
    3:  string path (api.raw = "path")
    4:  string description (api.raw = "description")
    5:  string method (api.raw = "method")
    6:  string group (api.raw = "group")
}

// api service
service ApisService {
  // 创建或API
  base.NilResponse CreateApi(1: ApiInfo req) (api.post = "/api/admin/api/create")

  // 更新API
  base.NilResponse UpdateApi(1: ApiInfo req) (api.post = "/api/admin/api/update")

  // 删除API信息
  base.NilResponse DeleteApi(1: base.IDReq req) (api.post = "/api/admin/api")

  // 获取API列表
  base.NilResponse ApiList(1: ApiPageReq req) (api.post = "/api/admin/api/list")
}