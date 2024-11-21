package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type UpdateCoachStatusService struct {
	ctx context.Context
} // NewUpdateCoachStatusService new UpdateCoachStatusService
func NewUpdateCoachStatusService(ctx context.Context) *UpdateCoachStatusService {
	return &UpdateCoachStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateCoachStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
