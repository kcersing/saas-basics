package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"system/biz/infra/service"
)

type RoleByIDService struct {
	ctx context.Context
} // NewRoleByIDService new RoleByIDService
func NewRoleByIDService(ctx context.Context) *RoleByIDService {
	return &RoleByIDService{ctx: ctx}
}

// Run create note info
func (s *RoleByIDService) Run(req *base.IDReq) (resp *auth.RoleInfo, err error) {
	// Finish your business logic.
	resp, err = service.NewRole(s.ctx).RoleInfoByID(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
