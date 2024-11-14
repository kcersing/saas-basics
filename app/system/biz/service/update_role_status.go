package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"system/biz/infra/service"
)

type UpdateRoleStatusService struct {
	ctx context.Context
} // NewUpdateRoleStatusService new UpdateRoleStatusService
func NewUpdateRoleStatusService(ctx context.Context) *UpdateRoleStatusService {
	return &UpdateRoleStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateRoleStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewRole(s.ctx).UpdateStatus(req.Id, req.Status)
	if err != nil {
		return nil, err
	}
	return
}
