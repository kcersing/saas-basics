package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	contract "rpc_gen/kitex_gen/company/contract"
)

type ContractByIDService struct {
	ctx context.Context
} // NewContractByIDService new ContractByIDService
func NewContractByIDService(ctx context.Context) *ContractByIDService {
	return &ContractByIDService{ctx: ctx}
}

// Run create note info
func (s *ContractByIDService) Run(req *base.IDReq) (resp *contract.ContractInfo, err error) {
	// Finish your business logic.

	return
}
