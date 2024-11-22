package service

import (
	"context"
	schedule "rpc_gen/kitex_gen/schedule"
	"testing"
)

func TestScheduleList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewScheduleListService(ctx)
	// init req and assert value

	req := &schedule.ScheduleListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
