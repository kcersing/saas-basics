package service

import (
	"context"
	user "rpc_gen/kitex_gen/user"
	"user/biz/infra/service"
)

type UserListService struct {
	ctx context.Context
} // NewUserListService new UserListService
func NewUserListService(ctx context.Context) *UserListService {
	return &UserListService{ctx: ctx}
}

// Run create note info
func (s *UserListService) Run(req *user.UserListReq) (resp *user.UserListResp, err error) {
	// Finish your business logic.

	resp.Extra, resp.Resp.Total, err = service.NewUser(s.ctx).List(req)
	if err != nil {
		return nil, err
	}
	return
}
