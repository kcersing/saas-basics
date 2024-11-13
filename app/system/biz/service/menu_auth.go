package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	"system/biz/infra/service"
)

type MenuAuthService struct {
	ctx context.Context
} // NewMenuAuthService new MenuAuthService
func NewMenuAuthService(ctx context.Context) *MenuAuthService {
	return &MenuAuthService{ctx: ctx}
}

// Run create note info
func (s *MenuAuthService) Run(req *base.IDReq) (resp []*menu.MenuInfoTree, err error) {
	// Finish your business logic.
	resp, err = service.NewMenu(s.ctx).MenuRole(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
