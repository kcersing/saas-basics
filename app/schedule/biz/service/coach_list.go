package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
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
	resp.Extra, resp.Resp.Total, err = service.NewSchedule(s.ctx).CoachList(*req)
	if err != nil {
		return nil, err
	}
	return
}
