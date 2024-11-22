package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
)

type MemberListService struct {
	ctx context.Context
} // NewMemberListService new MemberListService
func NewMemberListService(ctx context.Context) *MemberListService {
	return &MemberListService{ctx: ctx}
}

// Run create note info
func (s *MemberListService) Run(req *schedule.ScheduleMemberListReq) (resp *schedule.ScheduleMemberListResp, err error) {
	// Finish your business logic.

	resp.Extra, resp.Resp.Total, err = service.NewSchedule(s.ctx).MemberList(*req)
	if err != nil {
		return nil, err
	}
	return
}
