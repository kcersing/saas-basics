package service

import (
	"context"
	auth "rpc_gen/kitex_gen/system/auth"
	"testing"
)

func TestCreateAuth_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateAuthService(ctx)
	// init req and assert value

	req := &auth.CreateOrUpdateApiAuthReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
