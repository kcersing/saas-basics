package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	admin "system/biz/infra/service"
)

type MenuTreeService struct {
	ctx context.Context
} // NewMenuTreeService new MenuTreeService
func NewMenuTreeService(ctx context.Context) *MenuTreeService {
	return &MenuTreeService{ctx: ctx}
}

// Run create note info
func (s *MenuTreeService) Run(req *base.PageInfoReq) (resp []*menu.Tree, err error) {
	// Finish your business logic.
	resp, err = admin.NewMenu(s.ctx).MenuTree(req)
	if err != nil {
		return resp, kerrors.NewBizStatusError(40001, err.Error())
	}
	return resp, nil
}
