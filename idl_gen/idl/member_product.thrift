namespace go member.memberProduct

include "../base/base.thrift"

service MemberProductService {
    base.NilResponse MemberProductInfo(1: base.IDReq req) (api.post = "/service/member/product-info")
    base.NilResponse MemberProductList(1: MemberProductListReq req) (api.post = "/service/member/product-list")

//    base.NilResponse MemberProductSearch(1: MemberProductSearchReq req) (api.post = "/service/member/search-product")

}

struct MemberProductInfo {
    1: optional i64 id = 0  (api.raw = "id")
    2: optional string name = ""  (api.raw = "name")
    3: optional double price = 0  (api.raw = "price")
    4: optional double fee = 0 (api.raw = "fee")
    6: optional i64 status = 0  (api.raw = "status")
    7: optional i64 duration = 0 (api.raw = "duration")
    8: optional i64 length = 0 (api.raw = "length")
    9: optional string type = "" (api.raw = "type")
    10: optional i64 deadline = 0 (api.raw = "deadline")
    12: optional i64 number = 0 (api.raw = "number")
    13: optional i64 numberSurplus = 0 (api.raw = "numberSurplus")
    14: optional string validityAt = ""   (api.raw = "validityAt")
    15: optional string cancelAt = ""  (api.raw = "cancelAt")
    16: optional string createdAt = ""   (api.raw = "updatedAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")

    18: optional i64 orderId = 0  (api.raw = "orderId")
    19: optional i64 venueId = 0  (api.raw = "venueId")
    20: optional i64 productId = 0 (api.raw = "productId")
    21: optional i64 memberId = 0 (api.raw = "memberId")
    22: optional string subType = "" (api.raw = "subType")
    23: optional string sn = ""  (api.raw = "sn")

    /**课程-数组*/
    24: optional list<MemberProductCourses> courses = {}  (api.raw = "courses")
    /**团课*/
    25: optional list<MemberProductCourses> lessons = {} (api.raw = "lessons")


     /**课程 1不限2指定*/
     27: optional i64 isCourse =0 (api.raw = "isCourse")
}

struct MemberProductCourses {
    1: optional i64 id =0 (api.raw = "id")
    2: optional string type = "" (api.raw = "type")
    3: optional string name=""  (api.raw = "name")
    5: optional i64 memberProductId = 0  (api.raw = "member_product_id")
    6: optional i64 coursesId =0 (api.raw = "courses_id")


}

struct MemberProductListReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string name ="" (api.raw = "name")
    4: optional list<i64> status = {} (api.raw = "status")
    5: optional string type = "" (api.raw = "type") // 类型
    6: optional i64 venueId = 0 (api.raw = "venueId")
    7: optional string subType = "" (api.raw = "subType")
    8: optional i64 memberId = 0 (api.raw = "memberId")
}