package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type MenuAuthService struct {
	ctx context.Context
} // NewMenuAuthService new MenuAuthService
func NewMenuAuthService(ctx context.Context) *MenuAuthService {
	return &MenuAuthService{ctx: ctx}
}

// Run create note info
func (s *MenuAuthService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
