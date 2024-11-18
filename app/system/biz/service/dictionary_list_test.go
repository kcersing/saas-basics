package service

import (
	"context"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"testing"
)

func TestDictionaryList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDictionaryListService(ctx)
	// init req and assert value

	req := &dictionary.DictListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
