package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type MemberPropertySearchService struct {
	ctx context.Context
} // NewMemberPropertySearchService new MemberPropertySearchService
func NewMemberPropertySearchService(ctx context.Context) *MemberPropertySearchService {
	return &MemberPropertySearchService{ctx: ctx}
}

// Run create note info
func (s *MemberPropertySearchService) Run(req *member.MemberPropertySearchReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
