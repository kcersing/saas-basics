package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

type CreateUserService struct {
	ctx context.Context
} // NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// Run create note info
func (s *CreateUserService) Run(req *user.CreateOrUpdateUserReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
