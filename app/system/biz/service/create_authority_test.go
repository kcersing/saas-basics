package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"testing"
)

func TestCreateAuthority_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateAuthorityService(ctx)
	// init req and assert value

	req := &auth.CreateOrUpdateApiAuthorityReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
