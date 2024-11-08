package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

type CreateApiService struct {
	ctx context.Context
} // NewCreateApiService new CreateApiService
func NewCreateApiService(ctx context.Context) *CreateApiService {
	return &CreateApiService{ctx: ctx}
}

// Run create note info
func (s *CreateApiService) Run(req *auth.ApiInfo) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
