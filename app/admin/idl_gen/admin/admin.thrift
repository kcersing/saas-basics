namespace go admin.admin

include "../base/base.thrift"
include "../base/data.thrift"


// 验证码数据
struct CaptchaInfoResp {

}

// 登录数据
struct AdminLoginRequest {
    1:  string username (api.raw = "username" api.vd = "len($) > 0 && len($) < 33>")
    2:  string password (api.raw = "password" api.vd = "len($) > 0 && len($) < 33>")
}
// 更新密码数据
struct AdminChangePasswordRequest {
    1:  string old_password (api.raw = "old_password" api.vd = "len($) > 0 && len($) < 33>")
    2:  string new_password (api.raw = "new_password" api.vd = "len($) > 0 && len($) < 33>")
}
//添加用户数据
struct AddUserRequest {
    1:  string username (api.raw = "username" api.vd = "len($) > 0 && len($) < 33>")
    2:  string mobile (api.raw = "mobile")
    3:  string gender (api.raw = "gender")
    4:  string age (api.raw = "age")
    5:  string password (api.raw = "password")
}
//删除用户数据
struct DeleteUserRequest {
    1:  string account_id (api.raw = "account_id")
}

service AdminService {
    //检查系统状态
    base.NilResponse HealthCheck(1: base.Empty req) (api.get = "/api/health")
    //获取验证码
    base.NilResponse Captcha(1: CaptchaInfoResp req) (api.post = "/api/captcha")
    //登录
    base.NilResponse AdminLogin(1: AdminLoginRequest req) (api.post = "/admin/login")
    //更新密码
    base.NilResponse AdminChangePassword(1: AdminChangePasswordRequest req) (api.post = "/admin/password")
    //添加用户
    base.NilResponse AdminAddUser(1: AddUserRequest req) (api.post = "/admin/add-user")
    //删除用户
    base.NilResponse AdminDeleteUser(1: DeleteUserRequest req) (api.post = "/admin/delete-user")

}