package service

import (
	"context"

	base "facade/hertz_gen/base"
	member "facade/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberPropertyUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberPropertyUpdateService(Context context.Context, RequestContext *app.RequestContext) *MemberPropertyUpdateService {
	return &MemberPropertyUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberPropertyUpdateService) Run(req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
