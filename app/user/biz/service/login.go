package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	//err = service.NewUser(s.ctx).(req.Id)
	//if err != nil {
	//	return nil, err
	//}
	return
}
