package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type UserProfileService struct {
	ctx context.Context
} // NewUserProfileService new UserProfileService
func NewUserProfileService(ctx context.Context) *UserProfileService {
	return &UserProfileService{ctx: ctx}
}

// Run create note info
func (s *UserProfileService) Run(req *base.Empty) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
