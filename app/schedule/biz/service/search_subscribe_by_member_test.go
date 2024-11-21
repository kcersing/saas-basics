package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"testing"
)

func TestSearchSubscribeByMember_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchSubscribeByMemberService(ctx)
	// init req and assert value

	req := &schedule.SearchSubscribeByMemberReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
