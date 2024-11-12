package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	"testing"
)

func TestApiTree_Run(t *testing.T) {
	ctx := context.Background()
	s := NewApiTreeService(ctx)
	// init req and assert value

	req := &menu.ApiPageReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
