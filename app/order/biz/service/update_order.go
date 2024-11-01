package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/order"
)

type UpdateOrderService struct {
	ctx context.Context
} // NewUpdateOrderService new UpdateOrderService
func NewUpdateOrderService(ctx context.Context) *UpdateOrderService {
	return &UpdateOrderService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
