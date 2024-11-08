package service

import (
	"context"
	base "system/kitex_gen/base"
	auth "system/kitex_gen/system/auth"
)

type RoleByIDService struct {
	ctx context.Context
} // NewRoleByIDService new RoleByIDService
func NewRoleByIDService(ctx context.Context) *RoleByIDService {
	return &RoleByIDService{ctx: ctx}
}

// Run create note info
func (s *RoleByIDService) Run(req *base.IDReq) (resp *auth.RoleInfo, err error) {
	// Finish your business logic.

	return
}
