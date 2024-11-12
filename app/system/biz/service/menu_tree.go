package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type MenuTreeService struct {
	ctx context.Context
} // NewMenuTreeService new MenuTreeService
func NewMenuTreeService(ctx context.Context) *MenuTreeService {
	return &MenuTreeService{ctx: ctx}
}

// Run create note info
func (s *MenuTreeService) Run(req *base.PageInfoReq) (resp []*base.Tree, err error) {
	// Finish your business logic.

	return
}
