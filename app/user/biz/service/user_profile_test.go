package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"testing"
)

func TestUserProfile_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUserProfileService(ctx)
	// init req and assert value

	req := &base.Empty{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
