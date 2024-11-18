package system

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	base "rpc_gen/kitex_gen/base"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
)

func CreateDictionary(ctx context.Context, req *dictionary.DictionaryInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateDictionary(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateDictionary call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateDictionary(ctx context.Context, req *dictionary.DictionaryInfo, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateDictionary(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateDictionary call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteDictionary(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.DeleteDictionary(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteDictionary call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DictionaryList(ctx context.Context, req *dictionary.DictListReq, callOptions ...callopt.Option) (resp *dictionary.DictionaryListResp, err error) {
	resp, err = defaultClient.DictionaryList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DictionaryList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateDictionaryDetail(ctx context.Context, req *dictionary.DictionaryDetail, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.CreateDictionaryDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateDictionaryDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateDictionaryDetail(ctx context.Context, req *dictionary.DictionaryDetail, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.UpdateDictionaryDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateDictionaryDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteDictionaryDetail(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (resp *base.NilResponse, err error) {
	resp, err = defaultClient.DeleteDictionaryDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteDictionaryDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DetailByDictionaryList(ctx context.Context, req *dictionary.DetailListReq, callOptions ...callopt.Option) (resp *dictionary.DetailByDictionaryListResp, err error) {
	resp, err = defaultClient.DetailByDictionaryList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DetailByDictionaryList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
