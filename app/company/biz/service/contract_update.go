package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	contract "rpc_gen/kitex_gen/company/contract"
)

type ContractUpdateService struct {
	ctx context.Context
} // NewContractUpdateService new ContractUpdateService
func NewContractUpdateService(ctx context.Context) *ContractUpdateService {
	return &ContractUpdateService{ctx: ctx}
}

// Run create note info
func (s *ContractUpdateService) Run(req *contract.CreateOrUpdateContractReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
