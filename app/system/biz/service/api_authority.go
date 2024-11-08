package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type ApiAuthorityService struct {
	ctx context.Context
} // NewApiAuthorityService new ApiAuthorityService
func NewApiAuthorityService(ctx context.Context) *ApiAuthorityService {
	return &ApiAuthorityService{ctx: ctx}
}

// Run create note info
func (s *ApiAuthorityService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
