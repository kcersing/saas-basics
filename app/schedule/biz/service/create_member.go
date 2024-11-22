package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
)

type CreateMemberService struct {
	ctx context.Context
} // NewCreateMemberService new CreateMemberService
func NewCreateMemberService(ctx context.Context) *CreateMemberService {
	return &CreateMemberService{ctx: ctx}
}

// Run create note info
func (s *CreateMemberService) Run(req *schedule.CreateMemberReq) (resp *schedule.ScheduleMemberListResp, err error) {
	// Finish your business logic.

	err = service.NewSchedule(s.ctx).CreateMember(*req)
	if err != nil {
		return nil, err
	}
	return
}
