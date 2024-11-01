package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type MemberListService struct {
	ctx context.Context
} // NewMemberListService new MemberListService
func NewMemberListService(ctx context.Context) *MemberListService {
	return &MemberListService{ctx: ctx}
}

// Run create note info
func (s *MemberListService) Run(req *member.MemberListReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
