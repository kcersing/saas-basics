package user

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UserPermCode(ctx context.Context, req *base.Empty, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UserPermCode(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UserPermCode call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ChangePassword(ctx context.Context, req *user.ChangePasswordReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.ChangePassword(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ChangePassword call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateUser(ctx context.Context, req *user.CreateOrUpdateUserReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateUser(ctx context.Context, req *user.CreateOrUpdateUserReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UserInfo(ctx context.Context, req *base.Empty, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UserList(ctx context.Context, req *user.UserListReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UserList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UserList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.DeleteUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProfile(ctx context.Context, req *user.ProfileReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateProfile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProfile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UserProfile(ctx context.Context, req *base.Empty, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UserProfile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UserProfile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateUserStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateUserStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateUserStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SetUserRole(ctx context.Context, req *user.SetUserRole, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.SetUserRole(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SetUserRole call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SetDefaultVenue(ctx context.Context, req *user.SetDefaultVenueReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.SetDefaultVenue(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SetDefaultVenue call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
