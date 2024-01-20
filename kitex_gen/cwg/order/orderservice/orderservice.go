// Code generated by Kitex v0.8.0. DO NOT EDIT.

package orderservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	order "saas/kitex_gen/cwg/order"
)

func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

var orderServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateOrder":  kitex.NewMethodInfo(createOrderHandler, newOrderServiceCreateOrderArgs, newOrderServiceCreateOrderResult, false),
		"CancelOrder":  kitex.NewMethodInfo(cancelOrderHandler, newOrderServiceCancelOrderArgs, newOrderServiceCancelOrderResult, false),
		"ListOrder":    kitex.NewMethodInfo(listOrderHandler, newOrderServiceListOrderArgs, newOrderServiceListOrderResult, false),
		"GetOrderById": kitex.NewMethodInfo(getOrderByIdHandler, newOrderServiceGetOrderByIdArgs, newOrderServiceGetOrderByIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "order",
		"ServiceFilePath": `..\..\idl\order.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func createOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceCreateOrderArgs)
	realResult := result.(*order.OrderServiceCreateOrderResult)
	success, err := handler.(order.OrderService).CreateOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceCreateOrderArgs() interface{} {
	return order.NewOrderServiceCreateOrderArgs()
}

func newOrderServiceCreateOrderResult() interface{} {
	return order.NewOrderServiceCreateOrderResult()
}

func cancelOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceCancelOrderArgs)
	realResult := result.(*order.OrderServiceCancelOrderResult)
	success, err := handler.(order.OrderService).CancelOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceCancelOrderArgs() interface{} {
	return order.NewOrderServiceCancelOrderArgs()
}

func newOrderServiceCancelOrderResult() interface{} {
	return order.NewOrderServiceCancelOrderResult()
}

func listOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceListOrderArgs)
	realResult := result.(*order.OrderServiceListOrderResult)
	success, err := handler.(order.OrderService).ListOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceListOrderArgs() interface{} {
	return order.NewOrderServiceListOrderArgs()
}

func newOrderServiceListOrderResult() interface{} {
	return order.NewOrderServiceListOrderResult()
}

func getOrderByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceGetOrderByIdArgs)
	realResult := result.(*order.OrderServiceGetOrderByIdResult)
	success, err := handler.(order.OrderService).GetOrderById(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceGetOrderByIdArgs() interface{} {
	return order.NewOrderServiceGetOrderByIdArgs()
}

func newOrderServiceGetOrderByIdResult() interface{} {
	return order.NewOrderServiceGetOrderByIdResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (r *order.CreateOrderResp, err error) {
	var _args order.OrderServiceCreateOrderArgs
	_args.Req = req
	var _result order.OrderServiceCreateOrderResult
	if err = p.c.Call(ctx, "CreateOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CancelOrder(ctx context.Context, req *order.CancelOrderReq) (r *order.CancelOrderResp, err error) {
	var _args order.OrderServiceCancelOrderArgs
	_args.Req = req
	var _result order.OrderServiceCancelOrderResult
	if err = p.c.Call(ctx, "CancelOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListOrder(ctx context.Context, req *order.ListOrderReq) (r *order.ListOrderResp, err error) {
	var _args order.OrderServiceListOrderArgs
	_args.Req = req
	var _result order.OrderServiceListOrderResult
	if err = p.c.Call(ctx, "ListOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetOrderById(ctx context.Context, req *order.GetOrderByIdReq) (r *order.GetOrderByIdResp, err error) {
	var _args order.OrderServiceGetOrderByIdArgs
	_args.Req = req
	var _result order.OrderServiceGetOrderByIdResult
	if err = p.c.Call(ctx, "GetOrderById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}