package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type UpdateMenuAuthorityService struct {
	ctx context.Context
} // NewUpdateMenuAuthorityService new UpdateMenuAuthorityService
func NewUpdateMenuAuthorityService(ctx context.Context) *UpdateMenuAuthorityService {
	return &UpdateMenuAuthorityService{ctx: ctx}
}

// Run create note info
func (s *UpdateMenuAuthorityService) Run(req *auth.MenuAuthorityInfoReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
