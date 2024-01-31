namespace go admin

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