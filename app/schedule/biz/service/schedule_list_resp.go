package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
)

type ScheduleListRespService struct {
	ctx context.Context
} // NewScheduleListRespService new ScheduleListRespService
func NewScheduleListRespService(ctx context.Context) *ScheduleListRespService {
	return &ScheduleListRespService{ctx: ctx}
}

// Run create note info
func (s *ScheduleListRespService) Run(req *schedule.ScheduleListReq) (resp *schedule.ScheduleListResp, err error) {
	// Finish your business logic.

	return
}
