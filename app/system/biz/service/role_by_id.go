package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type RoleByIDService struct {
	ctx context.Context
} // NewRoleByIDService new RoleByIDService
func NewRoleByIDService(ctx context.Context) *RoleByIDService {
	return &RoleByIDService{ctx: ctx}
}

// Run create note info
func (s *RoleByIDService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
