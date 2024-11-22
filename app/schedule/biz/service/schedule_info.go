package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
)

type ScheduleInfoService struct {
	ctx context.Context
} // NewScheduleInfoService new ScheduleInfoService
func NewScheduleInfoService(ctx context.Context) *ScheduleInfoService {
	return &ScheduleInfoService{ctx: ctx}
}

// Run create note info
func (s *ScheduleInfoService) Run(req *base.IDReq) (resp *schedule.ScheduleInfo, err error) {
	// Finish your business logic.
	resp, err = service.NewSchedule(s.ctx).ScheduleInfo(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
