package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
)

type SearchSubscribeByMemberService struct {
	ctx context.Context
} // NewSearchSubscribeByMemberService new SearchSubscribeByMemberService
func NewSearchSubscribeByMemberService(ctx context.Context) *SearchSubscribeByMemberService {
	return &SearchSubscribeByMemberService{ctx: ctx}
}

// Run create note info
func (s *SearchSubscribeByMemberService) Run(req *schedule.SearchSubscribeByMemberReq) (resp *schedule.ScheduleMemberInfo, err error) {
	// Finish your business logic.

	return
}
