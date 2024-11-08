package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type ApiListService struct {
	ctx context.Context
} // NewApiListService new ApiListService
func NewApiListService(ctx context.Context) *ApiListService {
	return &ApiListService{ctx: ctx}
}

// Run create note info
func (s *ApiListService) Run(req *auth.ApiPageReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
