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


service SysService {
    /**商品列表*/
    base.NilResponse ProductList(1:SysProductListReq req) (api.post = "/service/sys/product/list")
    /**场馆列表*/
    base.NilResponse VenueList(1: base.ListReq req) (api.post = "/service/sys/venue/list")
    /**会员列表*/
    base.NilResponse MemberList(1: base.ListReq req) (api.post = "/service/sys/member/list")
    /**合同列表*/
    base.NilResponse ContractList(1: base.ListReq req) (api.post = "/service/sys/contract/list")
    /**员工列表*/
    base.NilResponse  StaffList(1: base.ListReq req) (api.post = "/service/sys/staff/list")
    /**场地列表*/
    base.NilResponse PlaceList(1: base.ListReq req) (api.post = "/service/sys/place/list")
    /**角色列表*/
    base.NilResponse RoleList(1: base.ListReq req) (api.post = "/service/sys/role/list")


}



