package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"testing"
)

func TestCreateMenuAuthority_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateMenuAuthorityService(ctx)
	// init req and assert value

	req := &auth.MenuAuthorityInfoReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
