package service

import (
	"context"
	user "rpc_gen/kitex_gen/user"
	"testing"
)

func TestUserList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUserListService(ctx)
	// init req and assert value

	req := &user.UserListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
