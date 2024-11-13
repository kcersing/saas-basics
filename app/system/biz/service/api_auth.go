package service

import (
	"context"
	"rpc_gen/kitex_gen/base"
	"rpc_gen/kitex_gen/system/auth"
	"strconv"
	"system/biz/infra/service"
)

type ApiAuthService struct {
	ctx context.Context
} // NewApiAuthService new ApiAuthService
func NewApiAuthService(ctx context.Context) *ApiAuthService {
	return &ApiAuthService{ctx: ctx}
}

// Run create note info
func (s *ApiAuthService) Run(req *base.IDReq) (resp []*auth.ApiAuthInfo, err error) {
	// Finish your business logic.
	resp, err = service.NewAuth(s.ctx).ApiAuth(strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}
	return
}
