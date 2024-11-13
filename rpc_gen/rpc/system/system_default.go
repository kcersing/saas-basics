package system

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

func CreateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.DeleteRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RoleByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *auth.RoleInfo, err error) {
	resp, err = defaultClient.RoleByID(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RoleByID call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateMenuAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateMenuAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateMenuAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateMenuAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RoleList(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (resp *auth.RoleListResp, err error) {
	resp, err = defaultClient.RoleList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RoleList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateRoleStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateRoleStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateApiAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateApiAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateApiAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ApiAuth(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp []*auth.ApiAuthInfo, err error) {
	resp, err = defaultClient.ApiAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ApiAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetLogsList(ctx context.Context, req *auth.LogsListReq, callOptions ...callopt.Option) (resp *auth.LogsListReq, err error) {
	resp, err = defaultClient.GetLogsList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetLogsList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteLogs(ctx context.Context, req *base.Ids, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.DeleteLogs(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteLogs call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
