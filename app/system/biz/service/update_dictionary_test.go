package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"testing"
)

func TestUpdateDictionary_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateDictionaryService(ctx)
	// init req and assert value

	req := &dictionary.DictionaryInfo{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
