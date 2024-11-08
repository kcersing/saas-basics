package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type CreateRoleService struct {
	ctx context.Context
} // NewCreateRoleService new CreateRoleService
func NewCreateRoleService(ctx context.Context) *CreateRoleService {
	return &CreateRoleService{ctx: ctx}
}

// Run create note info
func (s *CreateRoleService) Run(req *auth.RoleInfo) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
