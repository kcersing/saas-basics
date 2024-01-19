namespace go user

include "../base/common.thrift"
include "../base/user.thrift"


struct AdminLoginRequest {
    1:  string username (api.raw = "username" api.vd = "len($) > 0 && len($) < 33>")
    2:  string password (api.raw = "password" api.vd = "len($) > 0 && len($) < 33>")
}