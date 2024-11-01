package service

import (
	"context"

	base "facade/hertz_gen/base"
	member "facade/member"
	"github.com/cloudwego/hertz/pkg/app"
)

type MemberContractListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMemberContractListService(Context context.Context, RequestContext *app.RequestContext) *MemberContractListService {
	return &MemberContractListService{RequestContext: RequestContext, Context: Context}
}

func (h *MemberContractListService) Run(req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
