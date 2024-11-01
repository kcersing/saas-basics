package service

import (
	"context"

	base "facade/hertz_gen/base"
	member "facade/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberSearchService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberSearchService(Context context.Context, RequestContext *app.RequestContext) *MemberSearchService {
	return &MemberSearchService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberSearchService) Run(req *member.MemberSearchReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
