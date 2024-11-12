package system

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
)

func MenuAuth(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MenuAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MenuAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ApiList(ctx context.Context, req *menu.ApiPageReq, callOptions ...callopt.Option) (resp *menu.ApiInfoResp, err error) {
	resp, err = defaultClient.ApiList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ApiList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ApiTree(ctx context.Context, req *menu.ApiPageReq, callOptions ...callopt.Option) (resp []*base.Tree, err error) {
	resp, err = defaultClient.ApiTree(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ApiTree call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MenuByRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp []*menu.MenuInfoTree, err error) {
	resp, err = defaultClient.MenuByRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MenuByRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MenuLists(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MenuLists(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MenuLists call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MenuTree(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (resp []*base.Tree, err error) {
	resp, err = defaultClient.MenuTree(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MenuTree call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
