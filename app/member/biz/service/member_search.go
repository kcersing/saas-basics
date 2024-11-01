package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type MemberSearchService struct {
	ctx context.Context
} // NewMemberSearchService new MemberSearchService
func NewMemberSearchService(ctx context.Context) *MemberSearchService {
	return &MemberSearchService{ctx: ctx}
}

// Run create note info
func (s *MemberSearchService) Run(req *member.MemberSearchReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
