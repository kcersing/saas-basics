package service

import (
	"context"
	base "system/kitex_gen/base"
	auth "system/kitex_gen/system/auth"
)

type ApiTreeService struct {
	ctx context.Context
} // NewApiTreeService new ApiTreeService
func NewApiTreeService(ctx context.Context) *ApiTreeService {
	return &ApiTreeService{ctx: ctx}
}

// Run create note info
func (s *ApiTreeService) Run(req *auth.ApiPageReq) (resp []*base.Tree, err error) {
	// Finish your business logic.

	return
}
