package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type MemberPropertyUpdateService struct {
	ctx context.Context
} // NewMemberPropertyUpdateService new MemberPropertyUpdateService
func NewMemberPropertyUpdateService(ctx context.Context) *MemberPropertyUpdateService {
	return &MemberPropertyUpdateService{ctx: ctx}
}

// Run create note info
func (s *MemberPropertyUpdateService) Run(req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
