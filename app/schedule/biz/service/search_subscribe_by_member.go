package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"schedule/biz/infra/service"
)

type SearchSubscribeByMemberService struct {
	ctx context.Context
} // NewSearchSubscribeByMemberService new SearchSubscribeByMemberService
func NewSearchSubscribeByMemberService(ctx context.Context) *SearchSubscribeByMemberService {
	return &SearchSubscribeByMemberService{ctx: ctx}
}

// Run create note info
func (s *SearchSubscribeByMemberService) Run(req *schedule.SearchSubscribeByMemberReq) (resp *schedule.SearchSubscribeByMemberResp, err error) {
	// Finish your business logic.

	resp.Extra, resp.Resp.Total, err = service.NewSchedule(s.ctx).SearchSubscribeByMember(*req)
	if err != nil {
		return nil, err
	}
	return
}
