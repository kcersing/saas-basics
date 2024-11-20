namespace go admin.contract

include "../base/base.thrift"

struct ContractListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional string status (api.raw = "status")
}
struct CreateOrUpdateContractReq {
    1:  optional i64 id (api.raw = "id")
    2:  optional string name (api.raw = "name")
    3:  optional i64 status (api.raw = "status")
    4:  optional string content (api.raw = "content")
}
service CompanyService {
   base.NilResponse ContractList(1: ContractListReq req) (api.post = "/api/admin/contract/list")
   base.NilResponse ContractCreate(1: CreateOrUpdateContractReq req) (api.post = "/api/admin/contract/create")
   base.NilResponse ContractUpdate(1: CreateOrUpdateContractReq req) (api.post = "/api/admin/contract/update")
   base.NilResponse ContractUpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/contract/status")
   base.NilResponse ContractByID(1: base.IDReq req) (api.post = "/api/admin/contract")
}