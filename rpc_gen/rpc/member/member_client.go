package member

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/member/memberservice"
)

type RPCClient interface {
	KitexClient() memberservice.Client
	Service() string
	CreateMember(ctx context.Context, req *member.CreateOrUpdateMemberReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateMember(ctx context.Context, req *member.CreateOrUpdateMemberReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberInfo(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberList(ctx context.Context, req *member.MemberListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberSearch(ctx context.Context, req *member.MemberSearchReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberProductList(ctx context.Context, req *member.MemberProductListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberPropertyList(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberProductDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberPropertyDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberPropertyUpdate(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberContractList(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberProductSearch(ctx context.Context, req *member.MemberProductSearchReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	MemberPropertySearch(ctx context.Context, req *member.MemberPropertySearchReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := memberservice.NewClient(dstService, opts...)
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
	kitexClient memberservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() memberservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateMember(ctx context.Context, req *member.CreateOrUpdateMemberReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateMember(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateMember(ctx context.Context, req *member.CreateOrUpdateMemberReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateMember(ctx, req, callOptions...)
}

func (c *clientImpl) MemberInfo(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberInfo(ctx, req, callOptions...)
}

func (c *clientImpl) MemberList(ctx context.Context, req *member.MemberListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberList(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateMemberStatus(ctx, req, callOptions...)
}

func (c *clientImpl) MemberSearch(ctx context.Context, req *member.MemberSearchReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberSearch(ctx, req, callOptions...)
}

func (c *clientImpl) MemberProductList(ctx context.Context, req *member.MemberProductListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberProductList(ctx, req, callOptions...)
}

func (c *clientImpl) MemberPropertyList(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberPropertyList(ctx, req, callOptions...)
}

func (c *clientImpl) MemberProductDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberProductDetail(ctx, req, callOptions...)
}

func (c *clientImpl) MemberPropertyDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberPropertyDetail(ctx, req, callOptions...)
}

func (c *clientImpl) MemberPropertyUpdate(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberPropertyUpdate(ctx, req, callOptions...)
}

func (c *clientImpl) MemberContractList(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberContractList(ctx, req, callOptions...)
}

func (c *clientImpl) MemberProductSearch(ctx context.Context, req *member.MemberProductSearchReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberProductSearch(ctx, req, callOptions...)
}

func (c *clientImpl) MemberPropertySearch(ctx context.Context, req *member.MemberPropertySearchReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.MemberPropertySearch(ctx, req, callOptions...)
}
