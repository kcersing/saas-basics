namespace go admin.member

include "../base/base.thrift"
include "../base/data.thrift"


// The profile request | 更新个人数据
struct ProfileReq {
    1:  optional string nickname (api.raw = "nickname")
    2:  optional string avatar (api.raw = "avatar")
    3:  optional string mobile (api.raw = "mobile")
    4:  optional string email (api.raw = "email")
}

// Create or update user information request | 创建或更新用户信息
struct CreateOrUpdateMemberReq {
    1:  i64 id (api.raw = "id")
    2:  optional string avatar (api.raw = "avatar")
    4:  optional string mobile (api.raw = "mobile")
    5:  optional string email (api.raw = "email")
    6:  optional i64 status (api.raw = "status")
    7:  optional string username (api.raw = "username")
    8:  optional string nickname (api.raw = "nickname")
    9:  optional string password (api.raw = "password")
}

// Get user list request | 获取用户列表请求参数
struct MemberListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 limit (api.raw = "limit")
    3:  optional string username (api.raw = "username")
    4:  optional string nickname (api.raw = "nickname")
    5:  optional string email (api.raw = "email")
    6:  optional string mobile (api.raw = "mobile")
}


service MemberService {

  // 新增用户
  base.NilResponse CreateMember(1: CreateOrUpdateMemberReq req) (api.post = "/api/admin/member/create")

  // 更新用户
  base.NilResponse UpdateMember(1: CreateOrUpdateMemberReq req) (api.post = "/api/admin/member/update")

  // 获取用户基本信息
  base.NilResponse MemberInfo(1: base.Empty req) (api.get = "/api/admin/member/info")

  // 获取用户列表
  base.NilResponse MemberList(1: MemberListReq req) (api.post = "/api/admin/member/list")

  // 更新用户个人信息
  base.NilResponse UpdateProfile(1: ProfileReq req) (api.post = "/api/admin/member/profile-update")

  // 获取用户个人信息
  base.NilResponse MemberProfile(1: base.Empty req) (api.get = "/api/admin/member/profile")

  // 更新用户状态
  base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/member/status")

}