package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"system/biz/infra/service"
)

type DeleteRoleService struct {
	ctx context.Context
} // NewDeleteRoleService new DeleteRoleService
func NewDeleteRoleService(ctx context.Context) *DeleteRoleService {
	return &DeleteRoleService{ctx: ctx}
}

// Run create note info
func (s *DeleteRoleService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewRole(s.ctx).Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
