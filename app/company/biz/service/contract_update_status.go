package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type ContractUpdateStatusService struct {
	ctx context.Context
} // NewContractUpdateStatusService new ContractUpdateStatusService
func NewContractUpdateStatusService(ctx context.Context) *ContractUpdateStatusService {
	return &ContractUpdateStatusService{ctx: ctx}
}

// Run create note info
func (s *ContractUpdateStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
