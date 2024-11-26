namespace go sys

include "../base/base.thrift"
include "../base/data.thrift"


struct ListReq {
    1: optional string name (api.raw = "name")
    2: optional i64 dictionaryId (api.raw = "dictionaryId" )
    3: optional string type (api.raw = "type" )
    4: optional string mobile (api.raw = "mobile" )
    5: optional i64 product (api.raw = "product" )
    6: optional i64 venue (api.raw = "venue" )
}

service SysService {
    // 商品列表
    base.NilResponse ProductList(1: ListReq req) (api.post = "/api/sys/product/list")

    // 属性列表
    base.NilResponse PropertyList(1: ListReq req) (api.post = "/api/sys/property/list")

    // 属性类型
    base.NilResponse PropertyType(1: ListReq req) (api.post = "/api/sys/property/type")

    // 场馆列表
    base.NilResponse VenueList(1: ListReq req) (api.post = "/api/sys/venue/list")

    // 会员列表
    base.NilResponse MemberList(1: ListReq req) (api.post = "/api/sys/member/list")

    // 合同列表
    base.NilResponse ContractList(1: ListReq req) (api.post = "/api/sys/contract/list")

    // 员工列表
    base.NilResponse StaffList(1: ListReq req) (api.post = "/api/sys/staff/list")

    // 场地列表
    base.NilResponse PlaceList(1: ListReq req) (api.post = "/api/sys/place/list")

    base.NilResponse RoleList(1: ListReq req) (api.post = "/api/sys/role/list")


}



