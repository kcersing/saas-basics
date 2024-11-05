package service

import (
	"common/utils/errno"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"member/biz/infra/do"
	admin "member/biz/infra/service"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type CreateMemberService struct {
	ctx context.Context
} // NewCreateMemberService new CreateMemberService
func NewCreateMemberService(ctx context.Context) *CreateMemberService {
	return &CreateMemberService{ctx: ctx}
}

// Run create note info
func (s *CreateMemberService) Run(req *member.CreateOrUpdateMemberReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	var memberInfoReq do.CreateOrUpdateMemberReq
	err = copier.Copy(&memberInfoReq, &req)
	if err != nil {
		klog.Error(errno.ConvertErr(err))
	}
	err = admin.NewMember(s.ctx).Create(memberInfoReq)
	if err != nil {
		klog.Error(errno.ConvertErr(err))
	}
	return
}
