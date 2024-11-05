package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type MemberContractListService struct {
	ctx context.Context
} // NewMemberContractListService new MemberContractListService
func NewMemberContractListService(ctx context.Context) *MemberContractListService {
	return &MemberContractListService{ctx: ctx}
}

// Run create note info
func (s *MemberContractListService) Run(req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
