package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type MenuParamListByMenuIDService struct {
	ctx context.Context
} // NewMenuParamListByMenuIDService new MenuParamListByMenuIDService
func NewMenuParamListByMenuIDService(ctx context.Context) *MenuParamListByMenuIDService {
	return &MenuParamListByMenuIDService{ctx: ctx}
}

// Run create note info
func (s *MenuParamListByMenuIDService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
