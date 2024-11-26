namespace go system.auth

include "base/base.thrift"

// system service
service SystemService {

  // Create role information | 创建角色
  base.NilResponse CreateRole(1: RoleInfo req) (api.post = "/service/role/create")

  // Update role information | 更新角色
  base.NilResponse UpdateRole(1: RoleInfo req) (api.post = "/service/role/update")

  // Delete role information | 删除角色信息
  base.NilResponse DeleteRole(1: base.IDReq req) (api.post = "/service/role/del")

  // Get role information | 获取角色信息
  RoleInfo RoleByID(1: base.IDReq req) (api.get = "/service/role")

  // 创建菜单权限
  base.NilResponse CreateMenuAuth(1: MenuAuthInfoReq req) (api.post = "/service/auth/menu/create")

  // 更新菜单权限
  base.NilResponse UpdateMenuAuth(1: MenuAuthInfoReq req) (api.post = "/service/auth/menu/update")


  // Get role list | 获取角色列表
  RoleListResp RoleList(1: base.PageInfoReq req) (api.get = "/service/role/list")

  // Set role status | 设置角色状态, 启用1/禁用0
  base.NilResponse UpdateRoleStatus(1: base.StatusCodeReq req) (api.post = "/service/role/status")

  // 创建API权限
  base.NilResponse CreateAuth(1: CreateOrUpdateApiAuthReq req) (api.post = "/service/auth/api/create")

  // 更新API权限
  base.NilResponse UpdateAuth(1: CreateOrUpdateApiAuthReq req) (api.post = "/service/auth/api/update")

  // 获取角色api权限列表
 list<ApiAuthInfo> ApiAuth(1: base.IDReq req) (api.post = "/service/auth/api/role")

  // Get logs list | 获取日志列表
  LogsListResp GetLogsList(1: LogsListReq req) (api.post = "/service/logs/list")

  // Delete logs | 删除日志信息
  base.NilResponse DeleteLogs(1: base.Ids req) (api.post = "/service/logs/deleteAll")
}

// 菜单授权请求数据
struct MenuAuthInfoReq {
  1: i64 role_id (api.raw = "roleId")
    2: list<i64> menu_ids (api.raw = "menuIds")
}
struct RoleListResp {
    1: base.BaseResp resp
    2: optional list<RoleInfo> extra
}
struct LogsListResp {
    1: base.BaseResp resp
    2: optional list<LogsInfo> extra
}


//创建或更新角色信息参数
struct RoleInfo {
    1:  i64 ID (api.raw = "Id")
    2:  string name (api.raw = "name")
    3:  string value (api.raw = "value")
    4:  string defaultRouter (api.raw = "defaultRouter")
    5:  i64 status (api.raw = "status")
    6:  string remark  (api.raw = "remark")
    7:  i64 orderNo (api.raw = "orderNo")
    8:  string createdAt (api.raw = "createdAt")
    9:  string updatedAt (api.raw = "updatedAt")
    10: list<i64> Menus (api.raw = "menus")
    11: list<i64> Apis (api.raw = "apis")
}
//日志列表请求数据
struct LogsListReq {
    1:  i64 page (api.raw = "page")
    2:  i64 pageSize (api.raw = "pageSize")
    3:  string type (api.raw = "type")
    4:  string method (api.raw = "method")
    5:  string api (api.raw = "api")
    6:  string success (api.raw = "success")
    7:  string operators (api.raw = "operators")
}

// authorization message
// API授权数据
struct ApiAuthInfo {
    1:  string path (api.raw = "path")
    2:  string method (api.raw = "method")
}

// 创建或更新API授权信息
struct CreateOrUpdateApiAuthReq {
    1:  i64 role_id (api.raw = "roleId")
//    2:  ApiAuthInfo data (api.raw = "api_auth_info")
    2:  list<i64> apis (api.raw = "apis")
}



//bool: 布尔值，对应Java中的boolean，
//byte: 有符号字节，对应Java中的byte，对应MySQL的tinyint
//i16: 16位有符号整型，对应Java中的short，对应MySQL的smallint
//i32: 32位有符号整型，对应Java中的int，对应MySQL的int
//i64: 64位有符号整型，对应Java中的long，对应MySQL的bigint
//double: 64位浮点型，对应Java中的double
//string: 字符串，对应Java中的String
//binary: Blob 类型，对应Java中的byte[]
//list<t1>:一系列的t1类型组成的有序表，元素可以重复
//set<t1>:一系列的t1类型组成的无序集合，元素不可以重复
//map<t1,t2>:以t1为key,t2为value的键值对，t1不可以重复

struct LogsInfo {
	1: string Type
	2: string Method
	3: string Api
	4: bool Success
	5: string ReqContent
	6: string RespContent
	7: string Ip
	8: string UserAgent
	9: string Operator
	10: i64 Time
	11: string CreatedAt
	12: string UpdatedAt
}

