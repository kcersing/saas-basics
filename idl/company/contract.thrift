namespace go company.contract

include "../base/base.thrift"

struct ContractListReq {
    1: optional i64 page (api.raw = "page")
    2: optional i64 pageSize (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional string status (api.raw = "status")
}
struct CreateOrUpdateContractReq {
    1:  optional i64 id (api.raw = "id")
    2:  optional string name (api.raw = "name")
    3:  optional i64 status (api.raw = "status")
    4:  optional string content (api.raw = "content")
}

struct ContractInfo{
	1:i64 Id
	 2:string Name
	 3:i64 Status
	 4:string Content
	 5:string CreatedAt
	 6:string UpdatedAt
}
struct ContractListResp {
    1: base.BaseResp resp
    2: optional list<ContractInfo> extra
}
service CompanyService {
   ContractListResp ContractList(1: ContractListReq req) (api.post = "/service/contract/list")
   base.NilResponse ContractCreate(1: CreateOrUpdateContractReq req) (api.post = "/service/contract/create")
   base.NilResponse ContractUpdate(1: CreateOrUpdateContractReq req) (api.post = "/service/contract/update")
   base.NilResponse ContractUpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/contract/status")
   ContractInfo ContractByID(1: base.IDReq req) (api.post = "/service/contract")
}