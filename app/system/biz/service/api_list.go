package service

import (
	"context"
	auth "system/kitex_gen/system/auth"
)

type ApiListService struct {
	ctx context.Context
} // NewApiListService new ApiListService
func NewApiListService(ctx context.Context) *ApiListService {
	return &ApiListService{ctx: ctx}
}

// Run create note info
func (s *ApiListService) Run(req *auth.ApiPageReq) (resp []*auth.ApiInfo, err error) {
	// Finish your business logic.

	return
}
