package service

import (
	"context"
	base "system/kitex_gen/base"
	auth "system/kitex_gen/system/auth"
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
