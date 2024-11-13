package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"testing"
)

func TestDeleteLogs_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteLogsService(ctx)
	// init req and assert value

	req := &base.Ids{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
