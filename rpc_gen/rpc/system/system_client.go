package system

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/system/sys/systemservice"
)

type RPCClient interface {
	KitexClient() systemservice.Client
	Service() string
	ProductList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
	PropertyList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
	PropertyType(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
	VenueList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
	MemberList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
	ContractList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
	StaffList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
	PlaceList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
	RoleList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error)
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

func (c *clientImpl) ProductList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.ProductList(ctx, req, callOptions...)
}

func (c *clientImpl) PropertyList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.PropertyList(ctx, req, callOptions...)
}

func (c *clientImpl) PropertyType(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.PropertyType(ctx, req, callOptions...)
}

func (c *clientImpl) VenueList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.VenueList(ctx, req, callOptions...)
}

func (c *clientImpl) MemberList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.MemberList(ctx, req, callOptions...)
}

func (c *clientImpl) ContractList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.ContractList(ctx, req, callOptions...)
}

func (c *clientImpl) StaffList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.StaffList(ctx, req, callOptions...)
}

func (c *clientImpl) PlaceList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.PlaceList(ctx, req, callOptions...)
}

func (c *clientImpl) RoleList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (r *sys.SysListResp, err error) {
	return c.kitexClient.RoleList(ctx, req, callOptions...)
}
