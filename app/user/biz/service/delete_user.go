package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
