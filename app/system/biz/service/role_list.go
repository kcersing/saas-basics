package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"system/biz/infra/service"
)

type RoleListService struct {
	ctx context.Context
} // NewRoleListService new RoleListService
func NewRoleListService(ctx context.Context) *RoleListService {
	return &RoleListService{ctx: ctx}
}

// Run create note info
func (s *RoleListService) Run(req *base.PageInfoReq) (resp *auth.RoleListResp, err error) {
	// Finish your business logic.
	resp.Extra, resp.Resp.Total, err = service.NewRole(s.ctx).List(req)
	if err != nil {
		return nil, err
	}
	return
}
