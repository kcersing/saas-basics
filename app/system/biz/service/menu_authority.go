package service

import (
	"context"
	base "system/kitex_gen/base"
)

type MenuAuthorityService struct {
	ctx context.Context
} // NewMenuAuthorityService new MenuAuthorityService
func NewMenuAuthorityService(ctx context.Context) *MenuAuthorityService {
	return &MenuAuthorityService{ctx: ctx}
}

// Run create note info
func (s *MenuAuthorityService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
