namespace go system.token

include "../base/base.thrift"



// Token信息
struct TokenInfo {
    1:  i64 id (api.raw = "id")
    2:  string createdAt (api.raw = "createdAt")
    3:  string updatedAt (api.raw = "updatedAt")
    4:  i64 user_id (api.raw = "user_id")
    5:  string username (api.raw = "username")
    6:  string token (api.raw = "token")
    7:  string source (api.raw = "source")
    8:  string expiredAt (api.raw = "expiredAt")
}

// token列表请求参数
struct TokenListReq {
    1:  i64 page (api.raw = "page")
    2:  i64 limit (api.raw = "limit")
    3:  string username (api.raw = "username")
    4:  i64 user_id (api.raw = "user_id")
}

service TokenService{
  // 更新Token
  base.NilResponse UpdateToken(1: TokenInfo req) (api.post = "/api/admin/token/update")

  // 删除token信息
  base.NilResponse DeleteToken(1: base.IDReq req) (api.post = "/api/admin/token")

  // 获取token列表
  base.NilResponse TokenList(1: TokenListReq req) (api.post = "/api/admin/token/list")

}