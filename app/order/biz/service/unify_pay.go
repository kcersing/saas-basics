package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/orders/order"
)

type UnifyPayService struct {
	ctx context.Context
} // NewUnifyPayService new UnifyPayService
func NewUnifyPayService(ctx context.Context) *UnifyPayService {
	return &UnifyPayService{ctx: ctx}
}

// Run create note info
func (s *UnifyPayService) Run(req *order.UnifyPayReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
