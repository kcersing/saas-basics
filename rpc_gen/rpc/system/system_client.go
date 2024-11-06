package system

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/system/menu/menuservice"
)

type RPCClient interface {
	KitexClient() menuservice.Client
	Service() string
	CreateMenu(ctx context.Context, req *menu.CreateOrUpdateMenuReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateMenu(ctx context.Context, req *menu.CreateOrUpdateMenuReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteMenu(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MenuByRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *menu.MenuInfoTree, err error)
	MenuLists(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MenuTree(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	CreateMenuParam(ctx context.Context, req *menu.CreateOrUpdateMenuParamReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateMenuParam(ctx context.Context, req *menu.CreateOrUpdateMenuParamReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteMenuParam(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MenuParamListByMenuID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := menuservice.NewClient(dstService, opts...)
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
	kitexClient menuservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() menuservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateMenu(ctx context.Context, req *menu.CreateOrUpdateMenuReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateMenu(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateMenu(ctx context.Context, req *menu.CreateOrUpdateMenuReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateMenu(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteMenu(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteMenu(ctx, req, callOptions...)
}

func (c *clientImpl) MenuByRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *menu.MenuInfoTree, err error) {
	return c.kitexClient.MenuByRole(ctx, req, callOptions...)
}

func (c *clientImpl) MenuLists(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MenuLists(ctx, req, callOptions...)
}

func (c *clientImpl) MenuTree(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MenuTree(ctx, req, callOptions...)
}

func (c *clientImpl) CreateMenuParam(ctx context.Context, req *menu.CreateOrUpdateMenuParamReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateMenuParam(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateMenuParam(ctx context.Context, req *menu.CreateOrUpdateMenuParamReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateMenuParam(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteMenuParam(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteMenuParam(ctx, req, callOptions...)
}

func (c *clientImpl) MenuParamListByMenuID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MenuParamListByMenuID(ctx, req, callOptions...)
}
