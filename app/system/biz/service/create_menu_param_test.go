package service

import (
	"context"
	menu "rpc_gen/kitex_gen/system/menu"
	"testing"
)

func TestCreateMenuParam_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateMenuParamService(ctx)
	// init req and assert value

	req := &menu.CreateOrUpdateMenuParamReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
