package system

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/system/auth/systemservice"
)

type RPCClient interface {
	KitexClient() systemservice.Client
	Service() string
	CreateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	RoleByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *auth.RoleInfo, err error)
	CreateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	RoleList(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *auth.RoleListResp, err error)
	UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	CreateAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateApiAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ApiAuth(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r []*auth.ApiAuthInfo, err error)
	GetLogsList(ctx context.Context, req *auth.LogsListReq, callOptions ...callopt.Option) (r *auth.LogsListReq, err error)
	DeleteLogs(ctx context.Context, req *base.Ids, callOptions ...callopt.Option) (r *base.NilResponse, err error)
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

func (c *clientImpl) CreateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateRole(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateRole(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteRole(ctx, req, callOptions...)
}

func (c *clientImpl) RoleByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *auth.RoleInfo, err error) {
	return c.kitexClient.RoleByID(ctx, req, callOptions...)
}

func (c *clientImpl) CreateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateMenuAuth(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateMenuAuth(ctx, req, callOptions...)
}

func (c *clientImpl) RoleList(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *auth.RoleListResp, err error) {
	return c.kitexClient.RoleList(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateRoleStatus(ctx, req, callOptions...)
}

func (c *clientImpl) CreateAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateAuth(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateApiAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateApiAuth(ctx, req, callOptions...)
}

func (c *clientImpl) ApiAuth(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r []*auth.ApiAuthInfo, err error) {
	return c.kitexClient.ApiAuth(ctx, req, callOptions...)
}

func (c *clientImpl) GetLogsList(ctx context.Context, req *auth.LogsListReq, callOptions ...callopt.Option) (r *auth.LogsListReq, err error) {
	return c.kitexClient.GetLogsList(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteLogs(ctx context.Context, req *base.Ids, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteLogs(ctx, req, callOptions...)
}
