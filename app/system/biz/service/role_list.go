package service

import (
	"context"
	base "system/kitex_gen/base"
	auth "system/kitex_gen/system/auth"
)

type RoleListService struct {
	ctx context.Context
} // NewRoleListService new RoleListService
func NewRoleListService(ctx context.Context) *RoleListService {
	return &RoleListService{ctx: ctx}
}

// Run create note info
func (s *RoleListService) Run(req *base.PageInfoReq) (resp []*auth.RoleInfo, err error) {
	// Finish your business logic.

	return
}
