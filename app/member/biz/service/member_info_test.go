package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"testing"
)

func TestMemberInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMemberInfoService(ctx)
	// init req and assert value

	req := &base.IDReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
