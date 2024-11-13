package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"system/biz/infra/service"
)

type MenuRoleService struct {
	ctx context.Context
} // NewMenuRoleService new MenuRoleService
func NewMenuRoleService(ctx context.Context) *MenuRoleService {
	return &MenuRoleService{ctx: ctx}
}

// Run create note info
func (s *MenuRoleService) Run(req *base.IDReq) (resp *base.Ids, err error) {
	// Finish your business logic.
	resp.Ids, err = service.NewAuth(s.ctx).MenuAuth(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
