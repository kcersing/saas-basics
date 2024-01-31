namespace go user

include "../base/base.thrift"
include "../base/data.thrift"

struct LoginRequest {
    1:  string code (api.raw = "code")
}

struct GetUserInfoRequest {}

service UserService {
    // 用户登录
    base.NilResponse Login(1: LoginRequest req) (api.post = "/login/user")
    //用户信息
    base.NilResponse GetUserInfo(1: GetUserInfoRequest req) (api.get = "/user/info")

}
