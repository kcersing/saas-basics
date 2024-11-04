package service

import (
	"context"

	base "facade/hertz_gen/base"
	"facade/hertz_gen/facade/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateMemberService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateMemberService(Context context.Context, RequestContext *app.RequestContext) *UpdateMemberService {
	return &UpdateMemberService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateMemberService) Run(req *member.CreateOrUpdateMemberReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
