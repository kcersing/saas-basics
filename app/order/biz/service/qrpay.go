package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/order"
)

type QRPayService struct {
	ctx context.Context
} // NewQRPayService new QRPayService
func NewQRPayService(ctx context.Context) *QRPayService {
	return &QRPayService{ctx: ctx}
}

// Run create note info
func (s *QRPayService) Run(req *order.QRPayReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
