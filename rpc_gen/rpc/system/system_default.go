package system

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
)

func CreateMenu(ctx context.Context, req *menu.CreateOrUpdateMenuReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateMenu(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateMenu call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateMenu(ctx context.Context, req *menu.CreateOrUpdateMenuReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateMenu(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateMenu call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteMenu(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.DeleteMenu(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteMenu call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MenuByRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *menu.MenuInfoTree, err error) {
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

func MenuTree(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MenuTree(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MenuTree call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateMenuParam(ctx context.Context, req *menu.CreateOrUpdateMenuParamReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateMenuParam(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateMenuParam call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateMenuParam(ctx context.Context, req *menu.CreateOrUpdateMenuParamReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateMenuParam(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateMenuParam call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteMenuParam(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.DeleteMenuParam(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteMenuParam call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MenuParamListByMenuID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MenuParamListByMenuID(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MenuParamListByMenuID call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
