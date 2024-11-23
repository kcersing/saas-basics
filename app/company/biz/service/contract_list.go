package service

import (
	"context"
	contract "rpc_gen/kitex_gen/company/contract"
)

type ContractListService struct {
	ctx context.Context
} // NewContractListService new ContractListService
func NewContractListService(ctx context.Context) *ContractListService {
	return &ContractListService{ctx: ctx}
}

// Run create note info
func (s *ContractListService) Run(req *contract.ContractListReq) (resp *contract.ContractListResp, err error) {
	// Finish your business logic.

	return
}
