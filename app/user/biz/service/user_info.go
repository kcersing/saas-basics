package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type UserInfoService struct {
	ctx context.Context
} // NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

// Run create note info
func (s *UserInfoService) Run(req *base.Empty) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
