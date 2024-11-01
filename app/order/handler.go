package main

import (
	"context"
	"order/biz/service"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateOrderService(ctx).Run(req)

	return resp, err
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateOrderService(ctx).Run(req)

	return resp, err
}

// UpdateStatus implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateStatusService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// GetOrderById implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderById(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewGetOrderByIdService(ctx).Run(req)

	return resp, err
}

// UnifyPay implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UnifyPay(ctx context.Context, req *order.UnifyPayReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUnifyPayService(ctx).Run(req)

	return resp, err
}

// QRPay implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) QRPay(ctx context.Context, req *order.QRPayReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewQRPayService(ctx).Run(req)

	return resp, err
}
