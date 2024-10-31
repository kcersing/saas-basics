package service

import (
	"context"

	base "frontend/hertz_gen/base"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberProductDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberProductDetailService(Context context.Context, RequestContext *app.RequestContext) *MemberProductDetailService {
	return &MemberProductDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberProductDetailService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
