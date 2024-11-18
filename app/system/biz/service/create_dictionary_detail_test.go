package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"testing"
)

func TestCreateDictionaryDetail_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateDictionaryDetailService(ctx)
	// init req and assert value

	req := &dictionary.DictionaryDetail{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
