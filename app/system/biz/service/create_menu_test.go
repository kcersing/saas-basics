package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	"testing"
)

func TestCreateMenu_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateMenuService(ctx)
	// init req and assert value

	req := &menu.CreateOrUpdateMenuReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
