package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
)

type CreateMenuParamService struct {
	ctx context.Context
} // NewCreateMenuParamService new CreateMenuParamService
func NewCreateMenuParamService(ctx context.Context) *CreateMenuParamService {
	return &CreateMenuParamService{ctx: ctx}
}

// Run create note info
func (s *CreateMenuParamService) Run(req *menu.CreateOrUpdateMenuParamReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
