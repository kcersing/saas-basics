package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type CreateAuthorityService struct {
	ctx context.Context
} // NewCreateAuthorityService new CreateAuthorityService
func NewCreateAuthorityService(ctx context.Context) *CreateAuthorityService {
	return &CreateAuthorityService{ctx: ctx}
}

// Run create note info
func (s *CreateAuthorityService) Run(req *auth.CreateOrUpdateApiAuthorityReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
