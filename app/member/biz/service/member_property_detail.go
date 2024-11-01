package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type MemberPropertyDetailService struct {
	ctx context.Context
} // NewMemberPropertyDetailService new MemberPropertyDetailService
func NewMemberPropertyDetailService(ctx context.Context) *MemberPropertyDetailService {
	return &MemberPropertyDetailService{ctx: ctx}
}

// Run create note info
func (s *MemberPropertyDetailService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
