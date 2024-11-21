package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"testing"
)

func TestUpdateMemberStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateMemberStatusService(ctx)
	// init req and assert value

	req := &base.StatusCodeReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
