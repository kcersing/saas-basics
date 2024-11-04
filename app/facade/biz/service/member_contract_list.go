package service

import (
	"context"
	"facade/hertz_gen/facade/member"

	base "facade/hertz_gen/base"
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
