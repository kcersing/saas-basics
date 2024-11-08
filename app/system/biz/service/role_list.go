package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type RoleListService struct {
	ctx context.Context
} // NewRoleListService new RoleListService
func NewRoleListService(ctx context.Context) *RoleListService {
	return &RoleListService{ctx: ctx}
}

// Run create note info
func (s *RoleListService) Run(req *base.PageInfoReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
