package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/orders/order"
	"testing"
)

func TestListOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListOrderService(ctx)
	// init req and assert value

	req := &order.ListOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
