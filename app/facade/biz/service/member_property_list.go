package service

import (
	"context"

	base "facade/hertz_gen/base"
	member "facade/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberPropertyListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberPropertyListService(Context context.Context, RequestContext *app.RequestContext) *MemberPropertyListService {
	return &MemberPropertyListService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberPropertyListService) Run(req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
