namespace go member_product

include "../base/base.thrift"


service MemberProductService {
    base.NilResponse MemberProductInfo(1: base.IDReq req) (api.post = "/service/member/product-info")

    base.NilResponse MemberProductList(1: MemberProductListReq req) (api.post = "/service/member/product-list")

//    base.NilResponse MemberProductSearch(1: MemberProductSearchReq req) (api.post = "/service/member/search-product")

}
struct MemberProductInfo {
    1: optional i64 id  (api.raw = "id")
    2: optional string name  (api.raw = "name")
    3: optional double price  (api.raw = "price")
    4: optional double fee  (api.raw = "fee")
    6: optional i64 status  (api.raw = "status")
    7: optional i64 duration  (api.raw = "duration")
    8: optional i64 length  (api.raw = "length")
    9: optional string type  (api.raw = "type")
    10: optional i64 deadline (api.raw = "deadline")
    12: optional i64 count  (api.raw = "count")
    13: optional i64 countSurplus  (api.raw = "countSurplus")
    14: optional string validityAt   (api.raw = "validityAt")
    15: optional string cancelAt  (api.raw = "cancelAt")
    16: optional string createdAt   (api.raw = "updatedAt")
    17: optional string updatedAt  (api.raw = "updatedAt")
}

struct MemberProductListReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string name ="" (api.raw = "name")
    4: optional list<i64> status=0 (api.raw = "status")
    5: optional string type="" (api.raw = "type") // 类型
}


