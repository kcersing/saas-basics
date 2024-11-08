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

func RoleByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.RoleByID(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RoleByID call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RoleList(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
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

func CreateAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateAuthority(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateAuthority call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateApiAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateApiAuthority(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateApiAuthority call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ApiAuthority(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.ApiAuthority(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ApiAuthority call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateMenuAuthority(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateMenuAuthority call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateMenuAuthority(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateMenuAuthority call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MenuAuthority(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MenuAuthority(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MenuAuthority call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateApi(ctx context.Context, req *auth.ApiInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateApi(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateApi call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateApi(ctx context.Context, req *auth.ApiInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
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

func ApiList(ctx context.Context, req *auth.ApiPageReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.ApiList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ApiList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ApiTree(ctx context.Context, req *auth.ApiPageReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.ApiTree(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ApiTree call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
