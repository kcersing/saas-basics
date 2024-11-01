package member

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

func CreateMember(ctx context.Context, req *member.CreateOrUpdateMemberReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateMember(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateMember call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateMember(ctx context.Context, req *member.CreateOrUpdateMemberReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateMember(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateMember call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberInfo(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberList(ctx context.Context, req *member.MemberListReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateMemberStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateMemberStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberSearch(ctx context.Context, req *member.MemberSearchReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberSearch(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberSearch call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberProductList(ctx context.Context, req *member.MemberProductListReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberProductList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberProductList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberPropertyList(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberPropertyList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberPropertyList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberProductDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberProductDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberProductDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberPropertyDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberPropertyDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberPropertyDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberPropertyUpdate(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberPropertyUpdate(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberPropertyUpdate call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberContractList(ctx context.Context, req *member.MemberPropertyListReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberContractList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberContractList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberProductSearch(ctx context.Context, req *member.MemberProductSearchReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberProductSearch(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberProductSearch call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MemberPropertySearch(ctx context.Context, req *member.MemberPropertySearchReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.MemberPropertySearch(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MemberPropertySearch call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
