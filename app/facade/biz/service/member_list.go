package service

import (
	"context"

	base "facade/hertz_gen/base"
	"facade/hertz_gen/facade/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberListService(Context context.Context, RequestContext *app.RequestContext) *MemberListService {
	return &MemberListService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberListService) Run(req *member.MemberListReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
