package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"schedule/biz/infra/service"
)

type UpdateMemberStatusService struct {
	ctx context.Context
} // NewUpdateMemberStatusService new UpdateMemberStatusService
func NewUpdateMemberStatusService(ctx context.Context) *UpdateMemberStatusService {
	return &UpdateMemberStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateMemberStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewSchedule(s.ctx).UpdateMemberStatus(req.Id, req.Status)
	if err != nil {
		return nil, err
	}
	return
}
