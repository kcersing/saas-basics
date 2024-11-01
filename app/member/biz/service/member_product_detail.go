package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type MemberProductDetailService struct {
	ctx context.Context
} // NewMemberProductDetailService new MemberProductDetailService
func NewMemberProductDetailService(ctx context.Context) *MemberProductDetailService {
	return &MemberProductDetailService{ctx: ctx}
}

// Run create note info
func (s *MemberProductDetailService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
