package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"testing"
)

func TestCreateMember_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateMemberService(ctx)
	// init req and assert value

	req := &schedule.CreateMemberReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
