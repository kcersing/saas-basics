namespace go sys

include "../base/base.thrift"

struct SysProductListReq {
    1: optional string name="" (api.raw = "name")
    /**类型 card:卡 course:课 coursePackage:课包 Lessons:团课 */
    3: optional string type="" (api.raw = "type" )
    6: optional i64 venueId=0 (api.raw = "venueId" )
    /**次级类型courseOne一对一私教课 courseMore一对多私教课 cardTerm期限卡 cardSub次卡 lessons团课 coursePackage私教课包*/
    7: optional string subType="" (api.raw = "subType" )
}
struct SysStaffListReq {
    1: optional string name="" (api.raw = "name")
    /**职能*/
    3: optional list<string> functions="" (api.raw = "functions")
    6: optional i64 venueId=0 (api.raw = "venueId" )
    /**标签*/
    5: optional list<i64> tagId=0 (api.raw = "tagId" )
}

struct SysPlaceListReq {
    1: optional string name="" (api.raw = "name")
    /**类型 1场馆 2场地*/
    2: optional i64 type=0 (api.raw = "type" )
    /**对应数据字典值*/
    3: optional i64 classify=0 (api.raw = "classify" )
    6: optional i64 venueId=0 (api.raw = "venueId" )
    5: optional list<i64> productId=0 (api.raw = "productId" )
}

struct SysMemberListReq {
    1: optional string name="" (api.raw = "name")
    /**状态[1:潜在会员;2:正式会员]*/
    2: optional i64 condition=0 (api.raw = "condition" )
    4: optional string mobile="" (api.raw = "mobile" )
    6: optional i64 venueId=0 (api.raw = "venueId" )
    /**产品id 查询拥有此项产品的会员*/
    7: optional list<i64> productId=0 (api.raw = "productId" )
}

service SysService {
    /**商品列表*/
    base.NilResponse ProductList(1:SysProductListReq req) (api.post = "/service/sys/product/list")
    /**场馆列表*/
    base.NilResponse VenueList(1: base.ListReq req) (api.post = "/service/sys/venue/list")
    /**会员列表*/
    base.NilResponse MemberList(1: SysMemberListReq req) (api.post = "/service/sys/member/list")
    /**合同列表*/
    base.NilResponse ContractList(1: base.ListReq req) (api.post = "/service/sys/contract/list")
    /**员工列表*/
    base.NilResponse StaffList(1: SysStaffListReq req) (api.post = "/service/sys/staff/list")
    /**场地列表*/
    base.NilResponse PlaceList(1: SysPlaceListReq req) (api.post = "/service/sys/place/list")
    /**角色列表*/
    base.NilResponse RoleList(1: base.ListReq req) (api.post = "/service/sys/role/list")


}



