package service

import (
	"context"
	menu "rpc_gen/kitex_gen/system/menu"
	"system/biz/infra/service"
)

type ApiListService struct {
	ctx context.Context
} // NewApiListService new ApiListService
func NewApiListService(ctx context.Context) *ApiListService {
	return &ApiListService{ctx: ctx}
}

// Run create note info
func (s *ApiListService) Run(req *menu.ApiPageReq) (resp *menu.ApiInfoResp, err error) {
	// Finish your business logic.

	resp.Extra, resp.Resp.Total, err = service.NewApi(s.ctx).List(*req)
	if err != nil {
		return nil, err
	}
	return
}
