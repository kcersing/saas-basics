package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type UpdateUserStatusService struct {
	ctx context.Context
} // NewUpdateUserStatusService new UpdateUserStatusService
func NewUpdateUserStatusService(ctx context.Context) *UpdateUserStatusService {
	return &UpdateUserStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
