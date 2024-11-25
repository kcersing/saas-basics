package utils

import (
	"github.com/pkg/errors"
	"saas/app/admin/idl_gen/model/base"
	"saas/app/admin/pkg/errno"
)

func BuildBaseResp(err error) *base.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *base.BaseResponse {
	return &base.BaseResponse{
		Code:    err.ErrCode,
		Message: err.ErrMsg,
	}
}

func ParseBaseResp(resp *base.BaseResponse) error {
	if resp.Code == errno.Success.ErrCode {
		return nil
	}
	return errno.NewErrNo(resp.Code, resp.Message)
}
