package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type UpdateApiAuthService struct {
	ctx context.Context
} // NewUpdateApiAuthService new UpdateApiAuthService
func NewUpdateApiAuthService(ctx context.Context) *UpdateApiAuthService {
	return &UpdateApiAuthService{ctx: ctx}
}

// Run create note info
func (s *UpdateApiAuthService) Run(req *auth.CreateOrUpdateApiAuthReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
