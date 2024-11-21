package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
	"user/biz/infra/service"
)

type UserInfoService struct {
	ctx context.Context
} // NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

// Run create note info
func (s *UserInfoService) Run(req *base.IDReq) (resp *user.UserInfo, err error) {
	// Finish your business logic.

	resp, err = service.NewUser(s.ctx).Info(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
