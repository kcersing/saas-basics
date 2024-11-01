package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type MemberPropertyListService struct {
	ctx context.Context
} // NewMemberPropertyListService new MemberPropertyListService
func NewMemberPropertyListService(ctx context.Context) *MemberPropertyListService {
	return &MemberPropertyListService{ctx: ctx}
}

// Run create note info
func (s *MemberPropertyListService) Run(req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
