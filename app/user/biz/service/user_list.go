package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

type UserListService struct {
	ctx context.Context
} // NewUserListService new UserListService
func NewUserListService(ctx context.Context) *UserListService {
	return &UserListService{ctx: ctx}
}

// Run create note info
func (s *UserListService) Run(req *user.UserListReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
