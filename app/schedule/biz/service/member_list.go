package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
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

	return
}
