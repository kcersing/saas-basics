package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
)

type CoachListService struct {
	ctx context.Context
} // NewCoachListService new CoachListService
func NewCoachListService(ctx context.Context) *CoachListService {
	return &CoachListService{ctx: ctx}
}

// Run create note info
func (s *CoachListService) Run(req *schedule.ScheduleCoachListReq) (resp *schedule.ScheduleCoachListResp, err error) {
	// Finish your business logic.

	return
}
