// Code generated by Kitex v0.9.1. DO NOT EDIT.

package venueservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	base "rpc_gen/kitex_gen/base"
	venue "rpc_gen/kitex_gen/company/venue"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateVenuePlaceStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	VenuePlaceList(ctx context.Context, req *venue.VenuePlaceListReq, callOptions ...callopt.Option) (r *venue.VenuePlaceListResp, err error)
	CreateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateVenueStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	VenueList(ctx context.Context, req *venue.VenueListReq, callOptions ...callopt.Option) (r *venue.VenueListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kVenueServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVenueServiceClient struct {
	*kClient
}

func (p *kVenueServiceClient) CreateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateVenuePlace(ctx, req)
}

func (p *kVenueServiceClient) UpdateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateVenuePlace(ctx, req)
}

func (p *kVenueServiceClient) UpdateVenuePlaceStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateVenuePlaceStatus(ctx, req)
}

func (p *kVenueServiceClient) VenuePlaceList(ctx context.Context, req *venue.VenuePlaceListReq, callOptions ...callopt.Option) (r *venue.VenuePlaceListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VenuePlaceList(ctx, req)
}

func (p *kVenueServiceClient) CreateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateVenue(ctx, req)
}

func (p *kVenueServiceClient) UpdateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateVenue(ctx, req)
}

func (p *kVenueServiceClient) UpdateVenueStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateVenueStatus(ctx, req)
}

func (p *kVenueServiceClient) VenueList(ctx context.Context, req *venue.VenueListReq, callOptions ...callopt.Option) (r *venue.VenueListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VenueList(ctx, req)
}
