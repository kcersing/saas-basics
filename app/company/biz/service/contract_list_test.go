package service

import (
	"context"
	contract "rpc_gen/kitex_gen/company/contract"
	"testing"
)

func TestContractList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewContractListService(ctx)
	// init req and assert value

	req := &contract.ContractListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
