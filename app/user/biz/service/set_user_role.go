package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

type SetUserRoleService struct {
	ctx context.Context
} // NewSetUserRoleService new SetUserRoleService
func NewSetUserRoleService(ctx context.Context) *SetUserRoleService {
	return &SetUserRoleService{ctx: ctx}
}

// Run create note info
func (s *SetUserRoleService) Run(req *user.SetUserRole) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
