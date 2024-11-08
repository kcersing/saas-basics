package system

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/system/auth/roleservice"
)

type RPCClient interface {
	KitexClient() roleservice.Client
	Service() string
	CreateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	RoleByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *auth.RoleInfo, err error)
	RoleList(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r []*auth.RoleInfo, err error)
	UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	CreateAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateApiAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ApiAuthority(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	CreateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MenuAuthority(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	CreateApi(ctx context.Context, req *auth.ApiInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateApi(ctx context.Context, req *auth.ApiInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteApi(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ApiList(ctx context.Context, req *auth.ApiPageReq, callOptions ...callopt.Option) (r []*auth.ApiInfo, err error)
	ApiTree(ctx context.Context, req *auth.ApiPageReq, callOptions ...callopt.Option) (r []*base.Tree, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := roleservice.NewClient(dstService, opts...)
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
	kitexClient roleservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() roleservice.Client {
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

func (c *clientImpl) RoleList(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r []*auth.RoleInfo, err error) {
	return c.kitexClient.RoleList(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateRoleStatus(ctx, req, callOptions...)
}

func (c *clientImpl) CreateAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateAuthority(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateApiAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateApiAuthority(ctx, req, callOptions...)
}

func (c *clientImpl) ApiAuthority(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.ApiAuthority(ctx, req, callOptions...)
}

func (c *clientImpl) CreateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateMenuAuthority(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateMenuAuthority(ctx, req, callOptions...)
}

func (c *clientImpl) MenuAuthority(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MenuAuthority(ctx, req, callOptions...)
}

func (c *clientImpl) CreateApi(ctx context.Context, req *auth.ApiInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateApi(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateApi(ctx context.Context, req *auth.ApiInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateApi(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteApi(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteApi(ctx, req, callOptions...)
}

func (c *clientImpl) ApiList(ctx context.Context, req *auth.ApiPageReq, callOptions ...callopt.Option) (r []*auth.ApiInfo, err error) {
	return c.kitexClient.ApiList(ctx, req, callOptions...)
}

func (c *clientImpl) ApiTree(ctx context.Context, req *auth.ApiPageReq, callOptions ...callopt.Option) (r []*base.Tree, err error) {
	return c.kitexClient.ApiTree(ctx, req, callOptions...)
}
