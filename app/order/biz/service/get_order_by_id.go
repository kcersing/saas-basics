package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type GetOrderByIdService struct {
	ctx context.Context
} // NewGetOrderByIdService new GetOrderByIdService
func NewGetOrderByIdService(ctx context.Context) *GetOrderByIdService {
	return &GetOrderByIdService{ctx: ctx}
}

// Run create note info
func (s *GetOrderByIdService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
