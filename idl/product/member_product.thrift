namespace go member

include "../base/base.thrift"

struct MemberProductListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "memberId")
    4:  optional string name (api.raw = "name")
    5:  optional i64 venueId (api.raw = "venueId")
    6:  optional i64 status (api.raw = "status")
   
}
struct MemberPropertyListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "memberId")
    4:  optional string type (api.raw = "type")
    5:  optional string name (api.raw = "name")
    6:  optional i64 venueId (api.raw = "venueId")
    7:  optional i64 status (api.raw = "status")
    8:  optional i64 memberProductId (api.raw = "memberProductId")
}

struct MemberProductSearchReq{
    1:  optional list<i64> members (api.raw = "members")
}
struct MemberPropertySearchReq{
    1:  optional list<i64> memberProductIds (api.raw = "memberProducts")
}

service ProductService {

  base.NilResponse MemberProductList(1: MemberProductListReq req) (api.post = "/service/member/product-list")

  base.NilResponse MemberPropertyList(1: MemberPropertyListReq req) (api.post = "/service/member/property-list")

  base.NilResponse MemberProductDetail(1: base.IDReq req) (api.post = "/service/member/product-detail")

  base.NilResponse MemberPropertyDetail(1: base.IDReq req) (api.post = "/service/member/property-detail")

  base.NilResponse MemberPropertyUpdate(1: MemberPropertyListReq req) (api.post = "/service/member/property-update")

  base.NilResponse MemberProductSearch(1: MemberProductSearchReq req) (api.post = "/service/member/search-product")

  base.NilResponse MemberPropertySearch(1: MemberPropertySearchReq req) (api.post = "/service/member/search-property")

}
