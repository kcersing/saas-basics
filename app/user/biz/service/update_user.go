package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
	"user/biz/infra/service"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user.CreateOrUpdateUserReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewUser(s.ctx).Update(req)
	if err != nil {
		return nil, err
	}
	return
}
