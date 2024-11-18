package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
	"testing"
)

func TestMemberList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMemberListService(ctx)
	// init req and assert value

	req := &sys.ListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
