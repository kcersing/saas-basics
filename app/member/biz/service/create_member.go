package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

type CreateMemberService struct {
	ctx context.Context
} // NewCreateMemberService new CreateMemberService
func NewCreateMemberService(ctx context.Context) *CreateMemberService {
	return &CreateMemberService{ctx: ctx}
}

// Run create note info
func (s *CreateMemberService) Run(req *member.CreateOrUpdateMemberReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
