namespace go frontend.member

include "../base/base.thrift"
include "../base/data.thrift"


// Create or update user information request | 创建或更新用户信息
struct CreateOrUpdateMemberReq {
    1:  optional i64 id (api.raw = "id")
    2:  optional string avatar (api.raw = "avatar")
    4:  optional string mobile (api.raw = "mobile")
    5:  optional string email (api.raw = "email")
    6:  optional i64 status (api.raw = "status")
    7:  optional string name (api.raw = "name")
    8:  optional string gender (api.raw = "gender")
    9:  optional string wecom (api.raw = "wecom")
    10: optional i64 createId (api.raw = "createId")
    11: optional string birthday (api.raw = "birthday")
}

// Get user list request | 获取用户列表请求参数
struct MemberListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional string name (api.raw = "name")
    4:  optional string mobile (api.raw = "mobile")
}

struct MemberSearchReq {
    1:  string value (api.raw = "value")
    2:  string option (api.raw = "option")
}

struct MemberProductListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "member_id")
    4:  optional string name (api.raw = "name")
    5:  optional i64 venueId (api.raw = "venue_id")
    6:  optional i64 status (api.raw = "status")

}
struct MemberPropertyListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "member_id")
    4:  optional string type (api.raw = "type")
    5:  optional string name (api.raw = "name")
    6:  optional i64 venueId (api.raw = "venue_id")
    7:  optional i64 status (api.raw = "status")
    8:  optional i64 memberProductId (api.raw = "member_product_id")
}

struct MemberContractListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "member_id")
    4:  optional i64 venueId (api.raw = "venue_id")
    5:  optional i64 contractId (api.raw = "contract_id")
}
struct MemberProductSearchReq{
    1:  optional list<i64> members (api.raw = "members")
}
struct MemberPropertySearchReq{
    1:  optional list<i64> memberProductIds (api.raw = "member_products")
}

service MemberService {

  // 新增用户
  base.NilResponse CreateMember(1: CreateOrUpdateMemberReq req) (api.post = "/api/admin/member/create")

  // 更新用户
  base.NilResponse UpdateMember(1: CreateOrUpdateMemberReq req) (api.post = "/api/admin/member/update")

  // 获取用户基本信息
  base.NilResponse MemberInfo(1: base.IDReq req) (api.post = "/api/admin/member/info")

  // 获取用户列表
  base.NilResponse MemberList(1: MemberListReq req) (api.post = "/api/admin/member/list")

  // 更新用户状态
  base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/member/status")

  // 更新用户状态
  base.NilResponse MemberSearch(1: MemberSearchReq req) (api.post = "/api/admin/member/search")

  base.NilResponse MemberProductList(1: MemberProductListReq req) (api.post = "/api/admin/member/product-list")

  base.NilResponse MemberPropertyList(1: MemberPropertyListReq req) (api.post = "/api/admin/member/property-list")

  base.NilResponse MemberProductDetail(1: base.IDReq req) (api.post = "/api/admin/member/product-detail")

  base.NilResponse MemberPropertyDetail(1: base.IDReq req) (api.post = "/api/admin/member/property-detail")

  base.NilResponse MemberPropertyUpdate(1: MemberPropertyListReq req) (api.post = "/api/admin/member/property-update")

  base.NilResponse MemberContractList(1: MemberPropertyListReq req) (api.post = "/api/admin/member/contract-list")

  base.NilResponse MemberProductSearch(1: MemberProductSearchReq req) (api.post = "/api/admin/member/search-product")

  base.NilResponse MemberPropertySearch(1: MemberPropertySearchReq req) (api.post = "/api/admin/member/search-property")

}
