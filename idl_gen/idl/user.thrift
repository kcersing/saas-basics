namespace go user

include "../base/base.thrift"

struct UserInfo {
	1:i64 Id (api.raw = "id")
	2:i64 Status (api.raw = "status")
	3:string Username (api.raw = "username")
	4:string Password (api.raw = "password")
	5:string Name (api.raw = "name")
	6:i64 RoleId (api.raw = "roleId")
	7:string Mobile (api.raw = "mobile")
	8:string Avatar (api.raw = "avatar")
	9:string CreatedAt (api.raw = "createdAt")
	10:string UpdatedAt  (api.raw = "updatedAt")
	11:string RoleName (api.raw = "roleName")
	14:string RoleValue (api.raw = "roleValue")
	12:string Functions (api.raw = "functions")
	13:string Gender (api.raw = "gender")
//	14:i64 Age (api.raw = "age")
//	15:string Birthday (api.raw = "birthday")
    16:string Detail (api.raw = "detail")
    17:optional i64 jobTime (api.raw = "jobTime")
	254:i64 DefaultVenueId (api.raw = "defaultVenueId")
}

// login request | 登录参数
struct LoginReq {
    1:  string username (api.raw = "username")
    2:  string password (api.raw = "password")
    3:  string captchaId (api.raw = "captchaId")
    4:  string captcha (api.raw = "captcha")
}
struct LoginResp{
    1:  i64 userId (api.raw = "userId")
    2:  string username (api.raw = "username")
    3:  string roleName (api.raw = "roleName")
    4:  string roleValue (api.raw = "roleValue")
    5:  i64 roleId (api.raw = "roleId")
}

// register request | 注册参数
struct RegisterReq {
    1:  optional string username (api.raw = "username")
    2:  optional string password (api.raw = "password")
    3:  optional string captchaId (api.raw = "captchaId")
    4:  optional string captcha (api.raw = "captcha")
    5:  optional string email (api.raw = "email")
}

// change user's password request | 修改密码请求参数
struct ChangePasswordReq {
    1: i64 userId (api.raw = "userId")
    2: string newPassword (api.raw = "newPassword")
}

// Create or update user information request | 创建或更新用户信息
struct CreateOrUpdateUserReq {
    1:  optional i64 id (api.raw = "id")
    2:  optional string avatar (api.raw = "avatar")
    4:  optional string mobile (api.raw = "mobile")
    6:  optional i64 status (api.raw = "status")
    7:  optional string name (api.raw = "name")
    8:  optional string gender (api.raw = "gender")
    9:  optional i64 roleId (api.raw = "roleId")
    10: optional i64 createId (api.raw = "createId")
    12: optional string password (api.raw = "password")
    13:  optional string username (api.raw = "username")
    14:  optional string functions (api.raw = "functions")
    15:  optional string detail (api.raw = "detail")
    16:  optional i64 jobTime (api.raw = "jobTime")
}


// Get user list request | 获取用户列表请求参数
struct UserListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    4:  optional string name (api.raw = "name")
    5:  optional i64 jobTime (api.raw = "jobTime")
    6:  optional string mobile (api.raw = "mobile")
    7:  optional i64 roleId (api.raw = "roleId")

}

struct SetUserRole{
    1:  i64 userId (api.raw = "userId")
    2:  optional i64 roleId (api.raw = "roleId")
}

struct SetDefaultVenueReq{
    1: i64 userId (api.raw = "userId")
    2: optional i64 venueId (api.raw = "venueId")
}
struct Functions{
    1: optional i64 a (api.raw = "a")
    2: optional i64 b (api.raw = "b")
    3: optional i64 c (api.raw = "c")
    4: optional i64 d (api.raw = "d")
    5: optional i64 e (api.raw = "e")
}
service UserService {
  // 登录
//  base.NilResponse Login(1: LoginReq req) (api.post = "/service/login")

  // 注册
  //base.NilResponse Register(1: RegisterReq req) (api.post = "/service/register")

  // 修改密码
//  base.NilResponse ChangePassword(1: ChangePasswordReq req) (api.post = "/service/user/change-password")

  // 新增用户
  base.NilResponse CreateUser(1: CreateOrUpdateUserReq req) (api.post = "/service/user/create")

  // 更新用户
  base.NilResponse UpdateUser(1: CreateOrUpdateUserReq req) (api.post = "/service/user/update")

  // 获取用户基本信息
  base.NilResponse UserInfo(1: base.Empty req) (api.get = "/service/user/info")

  // 获取用户列表
  base.NilResponse UserList(1: UserListReq req) (api.post = "/service/user/list")

  // 删除用户信息
  base.NilResponse DeleteUser(1: base.IDReq req) (api.post = "/service/user")

  // 更新用户状态
//  base.NilResponse UpdateUserStatus(1: base.StatusCodeReq req) (api.post = "/service/user/status")

  // 设置用户角色
//  base.NilResponse SetUserRole(1: SetUserRole req) (api.post = "/service/user/set-role")

  // 设置用户默认场馆
//  base.NilResponse SetDefaultVenue(1: SetDefaultVenueReq req) (api.post = "/service/user/set-default-venue")

}
