package service

import (
	"context"
	base "system/kitex_gen/base"
	auth "system/kitex_gen/system/auth"
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
