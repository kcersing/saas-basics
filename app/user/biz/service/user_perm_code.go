package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
)

type UserPermCodeService struct {
	ctx context.Context
} // NewUserPermCodeService new UserPermCodeService
func NewUserPermCodeService(ctx context.Context) *UserPermCodeService {
	return &UserPermCodeService{ctx: ctx}
}

// Run create note info
func (s *UserPermCodeService) Run(req *base.Empty) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
