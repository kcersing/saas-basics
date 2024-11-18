package service

import (
	"context"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"testing"
)

func TestDetailByDictionaryList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDetailByDictionaryListService(ctx)
	// init req and assert value

	req := &dictionary.DetailListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
