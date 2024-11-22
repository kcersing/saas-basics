package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
)

type ScheduleDateListService struct {
	ctx context.Context
} // NewScheduleDateListService new ScheduleDateListService
func NewScheduleDateListService(ctx context.Context) *ScheduleDateListService {
	return &ScheduleDateListService{ctx: ctx}
}

// Run create note info
func (s *ScheduleDateListService) Run(req *schedule.ScheduleListReq) (resp *schedule.ScheduleDateListResp, err error) {
	// Finish your business logic.
	resp.Extra, resp.Resp.Total, err = service.NewSchedule(s.ctx).ScheduleDateList(*req)
	if err != nil {
		return nil, err
	}
	return
}
