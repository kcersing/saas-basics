package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type CreateMenuAuthorityService struct {
	ctx context.Context
} // NewCreateMenuAuthorityService new CreateMenuAuthorityService
func NewCreateMenuAuthorityService(ctx context.Context) *CreateMenuAuthorityService {
	return &CreateMenuAuthorityService{ctx: ctx}
}

// Run create note info
func (s *CreateMenuAuthorityService) Run(req *auth.MenuAuthorityInfoReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
