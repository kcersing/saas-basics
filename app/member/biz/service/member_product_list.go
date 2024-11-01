package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type MemberProductListService struct {
	ctx context.Context
} // NewMemberProductListService new MemberProductListService
func NewMemberProductListService(ctx context.Context) *MemberProductListService {
	return &MemberProductListService{ctx: ctx}
}

// Run create note info
func (s *MemberProductListService) Run(req *member.MemberProductListReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
