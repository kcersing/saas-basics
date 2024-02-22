namespace go admin.user

include "../base/base.thrift"
include "../base/data.thrift"

// user message

// login request | 登录参数
struct LoginReq {
    1:  string username (api.raw = "username")
    2:  string password (api.raw = "password")
    3:  string captcha_id (api.raw = "captcha_id")
    4:  string captcha (api.raw = "captcha")
}

// The profile request | 更新个人数据
struct ProfileReq {
    1:  string nickname (api.raw = "nickname")
    2:  string avatar (api.raw = "avatar")
    3:  string mobile (api.raw = "mobile")
    4:  string email (api.raw = "email")
}

// register request | 注册参数
struct RegisterReq {
    1:  string username (api.raw = "username")
    2:  string password (api.raw = "password")
    3:  string captcha_id (api.raw = "captcha_id")
    4:  string captcha (api.raw = "captcha")
    5:  string email (api.raw = "email")
}

// change user's password request | 修改密码请求参数
struct ChangePasswordReq {
    1:  string user_id (api.raw = "user_id")
    2:  string old_password (api.raw = "old_password")
    3:  string new_password (api.raw = "new_password")
}

// Create or update user information request | 创建或更新用户信息
struct CreateOrUpdateUserReq {
    1:  string id (api.raw = "id")
    2:  string avatar (api.raw = "avatar")
    3:  string role_id (api.raw = "role_id")
    4:  string mobile (api.raw = "mobile")
    5:  string email (api.raw = "email")
    6:  string status (api.raw = "status")
    7:  string username (api.raw = "username")
    8:  string nickname (api.raw = "nickname")
    9:  string password (api.raw = "password")
}

// Get user list request | 获取用户列表请求参数
struct UserListReq {
    1:  string page (api.raw = "page")
    2:  string limit (api.raw = "limit")
    3:  string username (api.raw = "username")
    4:  string nickname (api.raw = "nickname")
    5:  string email (api.raw = "email")
    6:  string mobile (api.raw = "mobile")
    7:  string role_id (api.raw = "role_id")
}


service UserService {
  // 注册
//  base.NilResponse Login(1: LoginReq req) (api.post = "/api/login")

  // 注册
  base.NilResponse Register(1: RegisterReq req) (api.post = "/api/register")

  // 获取用户权限码
  base.NilResponse UserPermCode(1: base.Empty req) (api.post = "/api/admin/user/perm")

  // 修改密码
  base.NilResponse ChangePassword(1: ChangePasswordReq req) (api.post = "/api/admin/user/change-password")

  // 新增用户
  base.NilResponse CreateUser(1: CreateOrUpdateUserReq req) (api.post = "/api/admin/user/create")

  // 更新用户
  base.NilResponse UpdateUser(1: CreateOrUpdateUserReq req) (api.post = "/api/admin/user/update")

  // 获取用户基本信息
  base.NilResponse UserInfo(1: base.Empty req) (api.get = "/api/admin/user/info")

  // 获取用户列表
  base.NilResponse UserList(1: UserListReq req) (api.post = "/api/admin/user/list")

  // 删除用户信息
  base.NilResponse DeleteUser(1: base.IDReq req) (api.post = "/api/admin/user")

  // 更新用户个人信息
  base.NilResponse UpdateProfile(1: ProfileReq req) (api.post = "/api/admin/user/profile-update")

  // 获取用户个人信息
  base.NilResponse UserProfile(1: base.Empty req) (api.get = "/api/admin/user/profile")

  // 更新用户状态
  base.NilResponse UpdateUserStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/user/status")

}
