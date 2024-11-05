package system

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	dictionary "rpc_gen/kitex_gen/system/dictionary"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/system/dictionary/dictionaryservice"
)

type RPCClient interface {
	KitexClient() dictionaryservice.Client
	Service() string
	CreateDictionary(ctx context.Context, req *dictionary.DictionaryInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateDictionary(ctx context.Context, req *dictionary.DictionaryInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteDictionary(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DictionaryList(ctx context.Context, req *dictionary.DictionaryPageReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	CreateDictionaryDetail(ctx context.Context, req *dictionary.DictionaryDetail, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateDictionaryDetail(ctx context.Context, req *dictionary.DictionaryDetail, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteDictionaryDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DetailByDictionaryList(ctx context.Context, req *dictionary.DictionaryDetailReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := dictionaryservice.NewClient(dstService, opts...)
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
	kitexClient dictionaryservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() dictionaryservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateDictionary(ctx context.Context, req *dictionary.DictionaryInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateDictionary(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateDictionary(ctx context.Context, req *dictionary.DictionaryInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateDictionary(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteDictionary(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteDictionary(ctx, req, callOptions...)
}

func (c *clientImpl) DictionaryList(ctx context.Context, req *dictionary.DictionaryPageReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DictionaryList(ctx, req, callOptions...)
}

func (c *clientImpl) CreateDictionaryDetail(ctx context.Context, req *dictionary.DictionaryDetail, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateDictionaryDetail(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateDictionaryDetail(ctx context.Context, req *dictionary.DictionaryDetail, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateDictionaryDetail(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteDictionaryDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DeleteDictionaryDetail(ctx, req, callOptions...)
}

func (c *clientImpl) DetailByDictionaryList(ctx context.Context, req *dictionary.DictionaryDetailReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.DetailByDictionaryList(ctx, req, callOptions...)
}
