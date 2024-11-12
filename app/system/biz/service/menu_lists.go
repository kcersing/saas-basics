package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type MenuListsService struct {
	ctx context.Context
} // NewMenuListsService new MenuListsService
func NewMenuListsService(ctx context.Context) *MenuListsService {
	return &MenuListsService{ctx: ctx}
}

// Run create note info
func (s *MenuListsService) Run(req *base.PageInfoReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
