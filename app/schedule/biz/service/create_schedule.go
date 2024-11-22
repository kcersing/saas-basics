package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
)

type CreateScheduleService struct {
	ctx context.Context
} // NewCreateScheduleService new CreateScheduleService
func NewCreateScheduleService(ctx context.Context) *CreateScheduleService {
	return &CreateScheduleService{ctx: ctx}
}

// Run create note info
func (s *CreateScheduleService) Run(req *schedule.CreateOrUpdateScheduleReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewSchedule(s.ctx).CreateSchedule(*req)
	if err != nil {
		return nil, err
	}
	return
}
