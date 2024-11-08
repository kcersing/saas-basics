package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"testing"
)

func TestUpdateApiAuthority_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateApiAuthorityService(ctx)
	// init req and assert value

	req := &auth.CreateOrUpdateApiAuthorityReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
