package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"testing"
)

func TestUpdateRoleStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateRoleStatusService(ctx)
	// init req and assert value

	req := &base.StatusCodeReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
