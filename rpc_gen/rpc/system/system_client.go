package system

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	role "rpc_gen/kitex_gen/system/role"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/system/role/apisservice"
)

type RPCClient interface {
	KitexClient() apisservice.Client
	Service() string
	CreateApi(ctx context.Context, req *role.ApiInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateApi(ctx context.Context, req *role.ApiInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteApi(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ApiList(ctx context.Context, req *role.ApiPageReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ApiTree(ctx context.Context, req *role.ApiPageReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := apisservice.NewClient(dstService, opts...)
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
	kitexClient apisservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() apisservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateApi(ctx context.Context, req *role.ApiInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateApi(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateApi(ctx context.Context, req *role.ApiInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateApi(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteApi(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteApi(ctx, req, callOptions...)
}

func (c *clientImpl) ApiList(ctx context.Context, req *role.ApiPageReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.ApiList(ctx, req, callOptions...)
}

func (c *clientImpl) ApiTree(ctx context.Context, req *role.ApiPageReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.ApiTree(ctx, req, callOptions...)
}
