package service

import (
	"context"

	base "facade/hertz_gen/base"
	member "facade/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateMemberService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateMemberService(Context context.Context, RequestContext *app.RequestContext) *CreateMemberService {
	return &CreateMemberService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateMemberService) Run(req *member.CreateOrUpdateMemberReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
