package do

import "saas/idl_gen/model/contract"

type Contract interface {
	Create(req *contract.CreateOrUpdateContractReq) error
	Update(req *contract.CreateOrUpdateContractReq) error
	UpdateStatus(ID int64, status int64) error
	Info(id int64) (info *contract.ContractInfo, err error)
	List(req *contract.ContractListReq) (list []*contract.ContractInfo, total int, err error)
}
