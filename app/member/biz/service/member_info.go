package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type MemberInfoService struct {
	ctx context.Context
} // NewMemberInfoService new MemberInfoService
func NewMemberInfoService(ctx context.Context) *MemberInfoService {
	return &MemberInfoService{ctx: ctx}
}

// Run create note info
func (s *MemberInfoService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
