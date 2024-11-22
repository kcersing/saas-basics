package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
)

type ScheduleListService struct {
	ctx context.Context
} // NewScheduleListService new ScheduleListService
func NewScheduleListService(ctx context.Context) *ScheduleListService {
	return &ScheduleListService{ctx: ctx}
}

// Run create note info
func (s *ScheduleListService) Run(req *schedule.ScheduleListReq) (resp *schedule.ScheduleListResp, err error) {
	// Finish your business logic.

	resp.Extra, resp.Resp.Total, err = service.NewSchedule(s.ctx).ScheduleList(*req)
	if err != nil {
		return nil, err
	}
	return
}
