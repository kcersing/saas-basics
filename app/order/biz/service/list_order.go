package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
