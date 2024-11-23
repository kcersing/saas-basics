package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	contract "rpc_gen/kitex_gen/company/contract"
	"testing"
)

func TestContractUpdate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewContractUpdateService(ctx)
	// init req and assert value

	req := &contract.CreateOrUpdateContractReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
