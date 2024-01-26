namespace go user

include "../base/common.thrift"
include "../base/user.thrift"

struct LoginRequest {
    1:  string code (api.raw = "code")
}

struct AdminLoginRequest {
    1:  string username (api.raw = "username" api.vd = "len($) > 0 && len($) < 33>")
    2:  string password (api.raw = "password" api.vd = "len($) > 0 && len($) < 33>")
}

struct AdminChangePasswordRequest {
    1:  string old_password (api.raw = "old_password" api.vd = "len($) > 0 && len($) < 33>")
    2:  string new_password (api.raw = "new_password" api.vd = "len($) > 0 && len($) < 33>")
}

struct AddUserRequest {
    1:  string account_id (api.raw = "account_id")
    2:  string username (api.raw = "username" api.vd = "len($) > 0 && len($) < 33>")
    3:  string phone_number (api.raw = "phone_number")
    4:  string avatar_blob_id (api.raw = "avatar_blob_id")
    5:  string open_id (api.raw = "open_id")
}

struct DeleteUserRequest {
    1:  string account_id (api.raw = "account_id")
}

struct GetUserInfoRequest {}

service UserService {
    // user
    common.NilResponse Login(1: LoginRequest req) (api.post = "/login/user")
    common.NilResponse GetUserInfo(1: GetUserInfoRequest req) (api.get = "/user/info")


    // admin
    common.NilResponse AdminLogin(1: AdminLoginRequest req) (api.post = "/admin/login")
    common.NilResponse AdminChangePassword(1: AdminChangePasswordRequest req) (api.post = "/admin/password")
    common.NilResponse AdminAddUser(1: AddUserRequest req) (api.post = "/admin/user")
    common.NilResponse AdminDeleteUser(1: DeleteUserRequest req) (api.delete = "/admin/user")
}
