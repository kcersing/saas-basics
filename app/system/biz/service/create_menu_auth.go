package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"system/biz/infra/service"
)

type CreateMenuAuthService struct {
	ctx context.Context
} // NewCreateMenuAuthService new CreateMenuAuthService
func NewCreateMenuAuthService(ctx context.Context) *CreateMenuAuthService {
	return &CreateMenuAuthService{ctx: ctx}
}

// Run create note info
func (s *CreateMenuAuthService) Run(req *auth.MenuAuthInfoReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewAuth(s.ctx).UpdateMenuAuth(req.RoleId, req.MenuIds)
	if err != nil {
		return nil, err
	}
	return
}
