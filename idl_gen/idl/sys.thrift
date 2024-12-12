namespace go sys

include "../base/base.thrift"


service SysService {
    // 商品列表
    base.NilResponse ProductList(1: base.ListReq req) (api.post = "/api/sys/product/list")

    // 场馆列表
    base.NilResponse VenueList(1: base.ListReq req) (api.post = "/api/sys/venue/list")

    // 会员列表
    base.NilResponse MemberList(1: base.ListReq req) (api.post = "/api/sys/member/list")

    // 合同列表
    base.NilResponse ContractList(1: base.ListReq req) (api.post = "/api/sys/contract/list")

    // 员工列表
    base.NilResponse  StaffList(1: base.ListReq req) (api.post = "/service/sys/staff/list")

    // 场地列表
    base.NilResponse PlaceList(1: base.ListReq req) (api.post = "/api/sys/place/list")

    base.NilResponse RoleList(1: base.ListReq req) (api.post = "/api/sys/role/list")


}



