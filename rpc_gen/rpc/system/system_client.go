package system

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/system/menu/systemservice"
)

type RPCClient interface {
	KitexClient() systemservice.Client
	Service() string
	MenuAuth(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r []*menu.MenuInfoTree, err error)
	MenuRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Ids, err error)
	ApiList(ctx context.Context, req *menu.ApiPageReq, callOptions ...callopt.Option) (r *menu.ApiInfoResp, err error)
	ApiTree(ctx context.Context, req *menu.ApiPageReq, callOptions ...callopt.Option) (r []*base.Tree, err error)
	MenuLists(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *menu.MenuInfoResp, err error)
	MenuTree(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r []*base.Tree, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := systemservice.NewClient(dstService, opts...)
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
	kitexClient systemservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() systemservice.Client {
	return c.kitexClient
}

func (c *clientImpl) MenuAuth(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r []*menu.MenuInfoTree, err error) {
	return c.kitexClient.MenuAuth(ctx, req, callOptions...)
}

func (c *clientImpl) MenuRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.Ids, err error) {
	return c.kitexClient.MenuRole(ctx, req, callOptions...)
}

func (c *clientImpl) ApiList(ctx context.Context, req *menu.ApiPageReq, callOptions ...callopt.Option) (r *menu.ApiInfoResp, err error) {
	return c.kitexClient.ApiList(ctx, req, callOptions...)
}

func (c *clientImpl) ApiTree(ctx context.Context, req *menu.ApiPageReq, callOptions ...callopt.Option) (r []*base.Tree, err error) {
	return c.kitexClient.ApiTree(ctx, req, callOptions...)
}

func (c *clientImpl) MenuLists(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *menu.MenuInfoResp, err error) {
	return c.kitexClient.MenuLists(ctx, req, callOptions...)
}

func (c *clientImpl) MenuTree(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r []*base.Tree, err error) {
	return c.kitexClient.MenuTree(ctx, req, callOptions...)
}
