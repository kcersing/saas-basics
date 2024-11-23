namespace go company.sys

include "../base/base.thrift"

service CompanyService {

    // 场馆列表
    base.SysListResp VenueList(1: base.ListReq req) (api.post = "/service/sys/venue/list")

    // 合同列表
    base.SysListResp ContractList(1: base.ListReq req) (api.post = "/service/sys/contract/list")

    // 场地列表
    base.SysListResp PlaceList(1: base.ListReq req) (api.post = "/service/sys/place/list")


}



