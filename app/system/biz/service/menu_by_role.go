package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	admin "system/biz/infra/service"
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
	resp, err = admin.NewMenu(s.ctx).MenuByRole(req.Id)
	if err != nil {
		return resp, kerrors.NewBizStatusError(40001, err.Error())
	}
	return resp, nil
}
