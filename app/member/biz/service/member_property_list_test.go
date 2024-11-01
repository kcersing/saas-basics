package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
	"testing"
)

func TestMemberPropertyList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMemberPropertyListService(ctx)
	// init req and assert value

	req := &member.MemberPropertyListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
