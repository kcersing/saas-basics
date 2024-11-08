package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type UpdateApiAuthorityService struct {
	ctx context.Context
} // NewUpdateApiAuthorityService new UpdateApiAuthorityService
func NewUpdateApiAuthorityService(ctx context.Context) *UpdateApiAuthorityService {
	return &UpdateApiAuthorityService{ctx: ctx}
}

// Run create note info
func (s *UpdateApiAuthorityService) Run(req *auth.CreateOrUpdateApiAuthorityReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
