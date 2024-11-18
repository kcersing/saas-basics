package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

type ChangePasswordService struct {
	ctx context.Context
} // NewChangePasswordService new ChangePasswordService
func NewChangePasswordService(ctx context.Context) *ChangePasswordService {
	return &ChangePasswordService{ctx: ctx}
}

// Run create note info
func (s *ChangePasswordService) Run(req *user.ChangePasswordReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
