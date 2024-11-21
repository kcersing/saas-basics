package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type UpdateStatusService struct {
	ctx context.Context
} // NewUpdateStatusService new UpdateStatusService
func NewUpdateStatusService(ctx context.Context) *UpdateStatusService {
	return &UpdateStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
