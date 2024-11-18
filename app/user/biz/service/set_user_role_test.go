package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
	"testing"
)

func TestSetUserRole_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSetUserRoleService(ctx)
	// init req and assert value

	req := &user.SetUserRole{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
