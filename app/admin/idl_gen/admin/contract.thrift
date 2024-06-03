namespace go admin.admin

include "../base/base.thrift"
include "../base/data.thrift"


struct ContractListReq {
    1:  i64 page (api.raw = "page")
    2:  i64 limit (api.raw = "limit")
    3:  string name (api.raw = "name")
    4:  string status (api.raw = "status")
}
struct CreateOrUpdateContractReq {
    1:  string name (api.raw = "name")
    2:  string status (api.raw = "status")
    3:  string content (api.raw = "content")
}
service ContractService {
   base.NilResponse ContractList(1: ContractListReq req) (api.post = "/api/admin/contract/list")
   base.NilResponse ContractCreate(1: CreateOrUpdateContractReq req) (api.post = "/api/admin/contract/create")
   base.NilResponse ContractUpdate(1: CreateOrUpdateContractReq req) (api.post = "/api/admin/contract/update")
   base.NilResponse ContractUpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/contract/status")
   base.NilResponse ContractByID(1: base.IDReq req) (api.post = "/api/admin/contract")
}