package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
)

type MenuByRoleService struct {
	ctx context.Context
} // NewMenuByRoleService new MenuByRoleService
func NewMenuByRoleService(ctx context.Context) *MenuByRoleService {
	return &MenuByRoleService{ctx: ctx}
}

// Run create note info
func (s *MenuByRoleService) Run(req *base.IDReq) (resp []*menu.MenuInfoTree, err error) {
	// Finish your business logic.

	return
}
