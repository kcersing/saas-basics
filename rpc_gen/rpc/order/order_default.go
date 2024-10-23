package order

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	order "rpc_gen/kitex_gen/orders/order"
)

func CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrder(ctx context.Context, req *order.UpdateOrderReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOrderById(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.GetOrderById(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrderById call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UnifyPay(ctx context.Context, req *order.UnifyPayReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UnifyPay(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UnifyPay call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QRPay(ctx context.Context, req *order.QRPayReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.QRPay(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QRPay call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
