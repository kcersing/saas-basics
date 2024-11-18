namespace go system.sys

include "../base/base.thrift"

service SystemService {
    // 商品列表
    base.SysListResp ProductList(1: base.ListReq req) (api.post = "/api/sys/product/list")

    // 属性列表
    base.SysListResp PropertyList(1: base.ListReq req) (api.post = "/api/sys/property/list")

    // 属性类型
    base.SysListResp PropertyType(1: base.ListReq req) (api.post = "/api/sys/property/type")

    // 场馆列表
    base.SysListResp VenueList(1: base.ListReq req) (api.post = "/api/sys/venue/list")

    // 会员列表
    base.SysListResp MemberList(1: base.ListReq req) (api.post = "/api/sys/member/list")

    // 合同列表
    base.SysListResp ContractList(1: base.ListReq req) (api.post = "/api/sys/contract/list")

    // 员工列表
    base.SysListResp StaffList(1: base.ListReq req) (api.post = "/api/sys/staff/list")

    // 场地列表
    base.SysListResp PlaceList(1: base.ListReq req) (api.post = "/api/sys/place/list")

    base.SysListResp RoleList(1: base.ListReq req) (api.post = "/api/sys/role/list")


}



