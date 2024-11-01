package service

import (
	"context"

	base "facade/hertz_gen/base"
	member "facade/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberPropertySearchService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberPropertySearchService(Context context.Context, RequestContext *app.RequestContext) *MemberPropertySearchService {
	return &MemberPropertySearchService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberPropertySearchService) Run(req *member.MemberPropertySearchReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
