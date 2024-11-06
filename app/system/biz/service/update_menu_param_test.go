package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	"testing"
)

func TestUpdateMenuParam_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateMenuParamService(ctx)
	// init req and assert value

	req := &menu.CreateOrUpdateMenuParamReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
