package service

import (
	"context"

	base "frontend/hertz_gen/base"
	member "frontend/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberProductSearchService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberProductSearchService(Context context.Context, RequestContext *app.RequestContext) *MemberProductSearchService {
	return &MemberProductSearchService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberProductSearchService) Run(req *member.MemberProductSearchReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
