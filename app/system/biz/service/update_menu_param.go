package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
)

type UpdateMenuParamService struct {
	ctx context.Context
} // NewUpdateMenuParamService new UpdateMenuParamService
func NewUpdateMenuParamService(ctx context.Context) *UpdateMenuParamService {
	return &UpdateMenuParamService{ctx: ctx}
}

// Run create note info
func (s *UpdateMenuParamService) Run(req *menu.CreateOrUpdateMenuParamReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
