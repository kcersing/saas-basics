package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type ApiTreeService struct {
	ctx context.Context
} // NewApiTreeService new ApiTreeService
func NewApiTreeService(ctx context.Context) *ApiTreeService {
	return &ApiTreeService{ctx: ctx}
}

// Run create note info
func (s *ApiTreeService) Run(req *auth.ApiPageReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
