package service

import (
	"context"
	base "system/kitex_gen/base"
	auth "system/kitex_gen/system/auth"
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
