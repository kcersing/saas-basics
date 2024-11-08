package service

import (
	"context"
	base "system/kitex_gen/base"
	auth "system/kitex_gen/system/auth"
)

type UpdateApiService struct {
	ctx context.Context
} // NewUpdateApiService new UpdateApiService
func NewUpdateApiService(ctx context.Context) *UpdateApiService {
	return &UpdateApiService{ctx: ctx}
}

// Run create note info
func (s *UpdateApiService) Run(req *auth.ApiInfo) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
