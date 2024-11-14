package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"system/biz/infra/service"
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
	err = service.NewRole(s.ctx).Create(req)
	if err != nil {
		return nil, err
	}
	return
}
