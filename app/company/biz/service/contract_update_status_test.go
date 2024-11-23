package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"testing"
)

func TestContractUpdateStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewContractUpdateStatusService(ctx)
	// init req and assert value

	req := &base.StatusCodeReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
