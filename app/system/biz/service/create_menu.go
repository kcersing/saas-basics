package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	admin "system/biz/infra/service"
)

type CreateMenuService struct {
	ctx context.Context
} // NewCreateMenuService new CreateMenuService
func NewCreateMenuService(ctx context.Context) *CreateMenuService {
	return &CreateMenuService{ctx: ctx}
}

// Run create note info
func (s *CreateMenuService) Run(req *menu.CreateOrUpdateMenuReq) (resp *base.NilResponse, err error) {

	err = admin.NewMenu(s.ctx).Create(req)
	if err != nil {
		return resp, kerrors.NewBizStatusError(40001, err.Error())
	}

	return resp, nil
}
