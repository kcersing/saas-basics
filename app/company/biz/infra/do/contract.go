package do

import "rpc_gen/kitex_gen/company/contract"

type Contract interface {
	Create(req *contract.ContractInfo) error
	Update(req *contract.ContractInfo) error
	UpdateStatus(ID int64, status int64) error
	Info(id int64) (info *contract.ContractInfo, err error)
	List(req *contract.ContractListReq) (list []*contract.ContractInfo, total int, err error)
}
