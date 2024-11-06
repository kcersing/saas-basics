package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type DeleteMenuParamService struct {
	ctx context.Context
} // NewDeleteMenuParamService new DeleteMenuParamService
func NewDeleteMenuParamService(ctx context.Context) *DeleteMenuParamService {
	return &DeleteMenuParamService{ctx: ctx}
}

// Run create note info
func (s *DeleteMenuParamService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
