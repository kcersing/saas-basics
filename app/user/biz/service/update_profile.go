package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

type UpdateProfileService struct {
	ctx context.Context
} // NewUpdateProfileService new UpdateProfileService
func NewUpdateProfileService(ctx context.Context) *UpdateProfileService {
	return &UpdateProfileService{ctx: ctx}
}

// Run create note info
func (s *UpdateProfileService) Run(req *user.ProfileReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
