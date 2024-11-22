package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	base "rpc_gen/kitex_gen/base"
	"schedule/biz/infra/service"
)

type UpdateStatusService struct {
	ctx context.Context
} // NewUpdateStatusService new UpdateStatusService
func NewUpdateStatusService(ctx context.Context) *UpdateStatusService {
	return &UpdateStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewSchedule(s.ctx).UpdateScheduleStatus(req.Id, req.Status)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, "修改失败")
	}
	return
}
