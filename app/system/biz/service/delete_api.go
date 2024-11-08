package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type DeleteApiService struct {
	ctx context.Context
} // NewDeleteApiService new DeleteApiService
func NewDeleteApiService(ctx context.Context) *DeleteApiService {
	return &DeleteApiService{ctx: ctx}
}

// Run create note info
func (s *DeleteApiService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
