package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"user/biz/infra/service"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	err = service.NewUser(s.ctx).DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
