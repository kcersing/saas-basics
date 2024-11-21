package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
	"testing"
)

func TestUpdateToken_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateTokenService(ctx)
	// init req and assert value

	req := &user.TokenInfo{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
