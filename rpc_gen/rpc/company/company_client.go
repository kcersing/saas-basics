package company

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	venue "rpc_gen/kitex_gen/company/venue"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"rpc_gen/kitex_gen/company/venue/venueservice"
)

type RPCClient interface {
	KitexClient() venueservice.Client
	Service() string
	CreateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateVenuePlaceStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	VenuePlaceList(ctx context.Context, req *venue.VenuePlaceListReq, callOptions ...callopt.Option) (r *venue.VenuePlaceListResp, err error)
	CreateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateVenueStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	VenueList(ctx context.Context, req *venue.VenueListReq, callOptions ...callopt.Option) (r *venue.VenueListResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := venueservice.NewClient(dstService, opts...)
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
	kitexClient venueservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() venueservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateVenuePlace(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateVenuePlace(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateVenuePlaceStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateVenuePlaceStatus(ctx, req, callOptions...)
}

func (c *clientImpl) VenuePlaceList(ctx context.Context, req *venue.VenuePlaceListReq, callOptions ...callopt.Option) (r *venue.VenuePlaceListResp, err error) {
	return c.kitexClient.VenuePlaceList(ctx, req, callOptions...)
}

func (c *clientImpl) CreateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.CreateVenue(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateVenue(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateVenueStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	return c.kitexClient.UpdateVenueStatus(ctx, req, callOptions...)
}

func (c *clientImpl) VenueList(ctx context.Context, req *venue.VenueListReq, callOptions ...callopt.Option) (r *venue.VenueListResp, err error) {
	return c.kitexClient.VenueList(ctx, req, callOptions...)
}
