package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/jinzhu/copier"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	"system/biz/infra/do"
	admin "system/biz/infra/service"
)

type UpdateMenuService struct {
	ctx context.Context
} // NewUpdateMenuService new UpdateMenuService
func NewUpdateMenuService(ctx context.Context) *UpdateMenuService {
	return &UpdateMenuService{ctx: ctx}
}

// Run create note info
func (s *UpdateMenuService) Run(req *menu.CreateOrUpdateMenuReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	var cReq do.MenuInfo
	err = copier.Copy(&cReq, req)
	if err != nil {
		return resp, kerrors.NewBizStatusError(40001, err.Error())
	}

	err = admin.NewMenu(s.ctx).Update(&cReq)
	if err != nil {
		return resp, kerrors.NewBizStatusError(40001, err.Error())
	}

	return resp, nil
}
