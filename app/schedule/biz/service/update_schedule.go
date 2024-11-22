package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
)

type UpdateScheduleService struct {
	ctx context.Context
} // NewUpdateScheduleService new UpdateScheduleService
func NewUpdateScheduleService(ctx context.Context) *UpdateScheduleService {
	return &UpdateScheduleService{ctx: ctx}
}

// Run create note info
func (s *UpdateScheduleService) Run(req *schedule.CreateOrUpdateScheduleReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewSchedule(s.ctx).UpdateSchedule(*req)
	if err != nil {
		return nil, err
	}
	return
}
