package service

import (
	"context"

	base "frontend/hertz_gen/base"
	member "frontend/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberProductListService(Context context.Context, RequestContext *app.RequestContext) *MemberProductListService {
	return &MemberProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberProductListService) Run(req *member.MemberProductListReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
