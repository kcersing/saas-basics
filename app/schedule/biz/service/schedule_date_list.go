package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
)

type ScheduleDateListService struct {
	ctx context.Context
} // NewScheduleDateListService new ScheduleDateListService
func NewScheduleDateListService(ctx context.Context) *ScheduleDateListService {
	return &ScheduleDateListService{ctx: ctx}
}

// Run create note info
func (s *ScheduleDateListService) Run(req *schedule.ScheduleListReq) (resp *schedule.ScheduleListResp, err error) {
	// Finish your business logic.

	return
}
