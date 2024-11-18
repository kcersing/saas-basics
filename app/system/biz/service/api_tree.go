package service

import (
	"context"
	"rpc_gen/kitex_gen/base"
	"rpc_gen/kitex_gen/system/menu"
	"system/biz/infra/service"
)

type ApiTreeService struct {
	ctx context.Context
} // NewApiTreeService new ApiTreeService
func NewApiTreeService(ctx context.Context) *ApiTreeService {
	return &ApiTreeService{ctx: ctx}
}

// Run create note info
func (s *ApiTreeService) Run(req *menu.ApiPageReq) (resp []*base.Tree, err error) {
	// Finish your business logic.

	resp, _, err = service.NewApi(s.ctx).ApiTree(*req)
	if err != nil {
		return nil, err
	}
	return

}
