package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/orders/order"
)

type CreateOrderService struct {
	ctx context.Context
} // NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{ctx: ctx}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
