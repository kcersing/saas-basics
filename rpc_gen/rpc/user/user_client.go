package user

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/user/userservice"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ChangePassword(ctx context.Context, req *user.ChangePasswordReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	CreateUser(ctx context.Context, req *user.CreateOrUpdateUserReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateUser(ctx context.Context, req *user.CreateOrUpdateUserReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UserInfo(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *user.UserInfo, err error)
	UserList(ctx context.Context, req *user.UserListReq, callOptions ...callopt.Option) (r *user.UserListResp, err error)
	DeleteUser(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateUserStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	SetUserRole(ctx context.Context, req *user.SetUserRole, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	SetDefaultVenue(ctx context.Context, req *user.SetDefaultVenueReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
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
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.Login(ctx, req, callOptions...)
}

func (c *clientImpl) ChangePassword(ctx context.Context, req *user.ChangePasswordReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.ChangePassword(ctx, req, callOptions...)
}

func (c *clientImpl) CreateUser(ctx context.Context, req *user.CreateOrUpdateUserReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateUser(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateUser(ctx context.Context, req *user.CreateOrUpdateUserReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateUser(ctx, req, callOptions...)
}

func (c *clientImpl) UserInfo(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *user.UserInfo, err error) {
	return c.kitexClient.UserInfo(ctx, req, callOptions...)
}

func (c *clientImpl) UserList(ctx context.Context, req *user.UserListReq, callOptions ...callopt.Option) (r *user.UserListResp, err error) {
	return c.kitexClient.UserList(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteUser(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteUser(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateUserStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateUserStatus(ctx, req, callOptions...)
}

func (c *clientImpl) SetUserRole(ctx context.Context, req *user.SetUserRole, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.SetUserRole(ctx, req, callOptions...)
}

func (c *clientImpl) SetDefaultVenue(ctx context.Context, req *user.SetDefaultVenueReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.SetDefaultVenue(ctx, req, callOptions...)
}
