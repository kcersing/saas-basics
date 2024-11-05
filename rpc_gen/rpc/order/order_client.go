package order

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/order"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/order/orderservice"
)

type RPCClient interface {
	KitexClient() orderservice.Client
	Service() string
	CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateOrder(ctx context.Context, req *order.UpdateOrderReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	GetOrderById(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UnifyPay(ctx context.Context, req *order.UnifyPayReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	QRPay(ctx context.Context, req *order.QRPayReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := orderservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient orderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() orderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateOrder(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateOrder(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateStatus(ctx, req, callOptions...)
}

func (c *clientImpl) ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.ListOrder(ctx, req, callOptions...)
}

func (c *clientImpl) GetOrderById(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.GetOrderById(ctx, req, callOptions...)
}

func (c *clientImpl) UnifyPay(ctx context.Context, req *order.UnifyPayReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UnifyPay(ctx, req, callOptions...)
}

func (c *clientImpl) QRPay(ctx context.Context, req *order.QRPayReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.QRPay(ctx, req, callOptions...)
}
