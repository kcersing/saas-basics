package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
	"user/biz/infra/service"
)

type UpdateTokenService struct {
	ctx context.Context
} // NewUpdateTokenService new UpdateTokenService
func NewUpdateTokenService(ctx context.Context) *UpdateTokenService {
	return &UpdateTokenService{ctx: ctx}
}

// Run create note info
func (s *UpdateTokenService) Run(req *user.TokenInfo) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewToken(s.ctx).Update(req)
	if err != nil {
		return nil, err
	}
	return
}
