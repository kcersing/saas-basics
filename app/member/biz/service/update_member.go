package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type UpdateMemberService struct {
	ctx context.Context
} // NewUpdateMemberService new UpdateMemberService
func NewUpdateMemberService(ctx context.Context) *UpdateMemberService {
	return &UpdateMemberService{ctx: ctx}
}

// Run create note info
func (s *UpdateMemberService) Run(req *member.CreateOrUpdateMemberReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
