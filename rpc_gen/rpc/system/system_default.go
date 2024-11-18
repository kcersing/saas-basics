package system

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	sys "rpc_gen/kitex_gen/system/sys"
)

func ProductList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.ProductList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ProductList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func PropertyList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.PropertyList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PropertyList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func PropertyType(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.PropertyType(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PropertyType call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VenueList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.VenueList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VenueList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.MemberList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ContractList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.ContractList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ContractList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func StaffList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.StaffList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "StaffList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func PlaceList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.PlaceList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RoleList(ctx context.Context, req *sys.ListReq, callOptions ...callopt.Option) (resp *sys.SysListResp, err error) {
	resp, err = defaultClient.RoleList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RoleList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
