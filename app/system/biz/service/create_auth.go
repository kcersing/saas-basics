package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type CreateAuthService struct {
	ctx context.Context
} // NewCreateAuthService new CreateAuthService
func NewCreateAuthService(ctx context.Context) *CreateAuthService {
	return &CreateAuthService{ctx: ctx}
}

// Run create note info
func (s *CreateAuthService) Run(req *auth.CreateOrUpdateApiAuthReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
