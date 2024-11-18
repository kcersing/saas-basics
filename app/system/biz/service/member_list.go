package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
)

type MemberListService struct {
	ctx context.Context
} // NewMemberListService new MemberListService
func NewMemberListService(ctx context.Context) *MemberListService {
	return &MemberListService{ctx: ctx}
}

// Run create note info
func (s *MemberListService) Run(req *sys.ListReq) (resp *sys.SysListResp, err error) {
	// Finish your business logic.

	return
}
