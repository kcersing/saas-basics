package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
	"testing"
)

func TestCreateUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateUserService(ctx)
	// init req and assert value

	req := &user.CreateOrUpdateUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
