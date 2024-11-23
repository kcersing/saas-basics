package company

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	venue "rpc_gen/kitex_gen/company/venue"
)

func CreateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateVenuePlace(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateVenuePlace call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateVenuePlace(ctx context.Context, req *venue.VenuePlaceInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateVenuePlace(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateVenuePlace call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateVenuePlaceStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateVenuePlaceStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateVenuePlaceStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VenuePlaceList(ctx context.Context, req *venue.VenuePlaceListReq, callOptions ...callopt.Option) (resp *venue.VenuePlaceListResp, err error) {
	resp, err = defaultClient.VenuePlaceList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VenuePlaceList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateVenue(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateVenue call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateVenue(ctx context.Context, req *venue.VenueInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateVenue(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateVenue call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateVenueStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateVenueStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateVenueStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VenueList(ctx context.Context, req *venue.VenueListReq, callOptions ...callopt.Option) (resp *venue.VenueListResp, err error) {
	resp, err = defaultClient.VenueList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VenueList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
