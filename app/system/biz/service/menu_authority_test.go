package service

import (
	"context"
	base "system/kitex_gen/base"
	"testing"
)

func TestMenuAuthority_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMenuAuthorityService(ctx)
	// init req and assert value

	req := &base.IDReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
