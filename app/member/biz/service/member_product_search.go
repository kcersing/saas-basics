package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type MemberProductSearchService struct {
	ctx context.Context
} // NewMemberProductSearchService new MemberProductSearchService
func NewMemberProductSearchService(ctx context.Context) *MemberProductSearchService {
	return &MemberProductSearchService{ctx: ctx}
}

// Run create note info
func (s *MemberProductSearchService) Run(req *member.MemberProductSearchReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
