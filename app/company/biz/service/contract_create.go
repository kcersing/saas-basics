package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	contract "rpc_gen/kitex_gen/company/contract"
)

type ContractCreateService struct {
	ctx context.Context
} // NewContractCreateService new ContractCreateService
func NewContractCreateService(ctx context.Context) *ContractCreateService {
	return &ContractCreateService{ctx: ctx}
}

// Run create note info
func (s *ContractCreateService) Run(req *contract.CreateOrUpdateContractReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
