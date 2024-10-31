package service

import (
	"context"

	base "frontend/hertz_gen/base"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberPropertyDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberPropertyDetailService(Context context.Context, RequestContext *app.RequestContext) *MemberPropertyDetailService {
	return &MemberPropertyDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberPropertyDetailService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
