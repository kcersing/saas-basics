package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
)

type MenuListsService struct {
	ctx context.Context
} // NewMenuListsService new MenuListsService
func NewMenuListsService(ctx context.Context) *MenuListsService {
	return &MenuListsService{ctx: ctx}
}

// Run create note info
func (s *MenuListsService) Run(req *base.PageInfoReq) (resp *menu.MenuInfoResp, err error) {
	// Finish your business logic.

	return
}
