package service

import (
	"context"
	base "system/kitex_gen/base"
)

type UpdateRoleStatusService struct {
	ctx context.Context
} // NewUpdateRoleStatusService new UpdateRoleStatusService
func NewUpdateRoleStatusService(ctx context.Context) *UpdateRoleStatusService {
	return &UpdateRoleStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateRoleStatusService) Run(req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
