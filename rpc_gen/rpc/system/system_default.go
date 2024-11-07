package system

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	role "rpc_gen/kitex_gen/system/role"
)

func CreateApi(ctx context.Context, req *role.ApiInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateApi(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateApi call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateApi(ctx context.Context, req *role.ApiInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateApi(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateApi call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteApi(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.DeleteApi(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteApi call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ApiList(ctx context.Context, req *role.ApiPageReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.ApiList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ApiList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ApiTree(ctx context.Context, req *role.ApiPageReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.ApiTree(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ApiTree call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
