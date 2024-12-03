namespace go contract

include "../base/base.thrift"

struct ContractListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string name = "" (api.raw = "name")
    4: optional string status = ""(api.raw = "status")
}
struct CreateOrUpdateContractReq {
    1:  optional i64 id (api.raw = "id")
    2:  optional string name (api.raw = "name")
    3:  optional i64 status (api.raw = "status")
    4:  optional string content (api.raw = "content")
}

struct ContractInfo{
	1:i64 Id (api.raw = "id")
	 2:string Name (api.raw = "name")
	 3:i64 Status (api.raw = "status")
	 4:string Content (api.raw = "content")
	 5:string CreatedAt (api.raw = "createdAt")
	 6:string UpdatedAt (api.raw = "updatedAt")
}

service ContractService {
   base.NilResponse ContractList(1: ContractListReq req) (api.post = "/service/contract/list")
   base.NilResponse ContractCreate(1: CreateOrUpdateContractReq req) (api.post = "/service/contract/create")
   base.NilResponse ContractUpdate(1: CreateOrUpdateContractReq req) (api.post = "/service/contract/update")
   base.NilResponse ContractUpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/contract/status")
   base.NilResponse ContractByID(1: base.IDReq req) (api.post = "/service/contract")
}