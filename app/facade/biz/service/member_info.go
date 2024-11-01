package service

import (
	"context"

	base "facade/hertz_gen/base"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberInfoService(Context context.Context, RequestContext *app.RequestContext) *MemberInfoService {
	return &MemberInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberInfoService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
