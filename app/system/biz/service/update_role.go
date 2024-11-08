package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type UpdateRoleService struct {
	ctx context.Context
} // NewUpdateRoleService new UpdateRoleService
func NewUpdateRoleService(ctx context.Context) *UpdateRoleService {
	return &UpdateRoleService{ctx: ctx}
}

// Run create note info
func (s *UpdateRoleService) Run(req *auth.RoleInfo) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
