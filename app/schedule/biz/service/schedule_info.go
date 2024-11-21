package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	schedule "rpc_gen/kitex_gen/schedule"
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

	return
}
