include "../base/common.thrift"
include "../base/user.thrift"

namespace go cwg.user

struct LoginRequest {
    1:  string code
}

struct LoginResponse {
    1:  common.BaseResponse base_resp
    2:  string token
}

struct AdminLoginRequest {
    1:  string username
    2:  string password
}

struct AdminLoginResponse {
    1:  common.BaseResponse base_resp
    2:  string token
}

struct ChangeAdminPasswordRequest {
    1:  string account_id
    2:  string old_password
    3:  string new_password
}

struct ChangeAdminPasswordResponse {
    1:  common.BaseResponse base_resp
}

struct GetUserInfoResponse {
    1:  common.BaseResponse base_resp
    2:  user.UserInfo user_info
}
struct GetUserRequest {
    1:  string accont_id
}

struct AddUserRequest {
    1:  string account_id
    2:  string username
    3:  string phone_number
    4:  string avatar_blob_id
    5:  string open_id
}

struct AddUserResponse {
    1:  common.BaseResponse base_resp
}

struct DeleteUserRequest {
    1:  string account_id
}

struct DeleteUserResponse {
    1:  common.BaseResponse base_resp
}

service UserService {
    LoginResponse Login(1: LoginRequest req)
    AdminLoginResponse AdminLogin(1: AdminLoginRequest req)
    ChangeAdminPasswordResponse ChangeAdminPassword(1: ChangeAdminPasswordRequest req)
    GetUserInfoResponse GetUser(1: GetUserRequest req)
    AddUserResponse AddUser(1: AddUserRequest req)
    DeleteUserResponse DeleteUser(1: DeleteUserRequest req)
}