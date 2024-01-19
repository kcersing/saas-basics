include "../base/common.thrift"
include "../base/user.thrift"

namespace go cwg.user

struct AdminLoginRequest {
    1:  string username
    2:  string password
}

struct AdminLoginResponse {
    1:  common.BaseResponse base_resp
    2:  string token
}

service UserService {
    AdminLoginResponse AdminLogin(1: AdminLoginRequest req)
}