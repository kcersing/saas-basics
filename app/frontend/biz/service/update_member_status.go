package service

import (
	"context"

	base "frontend/hertz_gen/base"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateMemberStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateMemberStatusService(Context context.Context, RequestContext *app.RequestContext) *UpdateMemberStatusService {
	return &UpdateMemberStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateMemberStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
