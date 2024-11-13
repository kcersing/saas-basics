package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type UpdateMenuAuthService struct {
	ctx context.Context
} // NewUpdateMenuAuthService new UpdateMenuAuthService
func NewUpdateMenuAuthService(ctx context.Context) *UpdateMenuAuthService {
	return &UpdateMenuAuthService{ctx: ctx}
}

// Run create note info
func (s *UpdateMenuAuthService) Run(req *auth.MenuAuthInfoReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
