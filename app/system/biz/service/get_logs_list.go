package service

import (
	"context"
	auth "rpc_gen/kitex_gen/system/auth"
	"system/biz/infra/service"
)

type GetLogsListService struct {
	ctx context.Context
} // NewGetLogsListService new GetLogsListService
func NewGetLogsListService(ctx context.Context) *GetLogsListService {
	return &GetLogsListService{ctx: ctx}
}

// Run create note info
func (s *GetLogsListService) Run(req *auth.LogsListReq) (resp *auth.LogsListResp, err error) {
	// Finish your business logic.
	resp.Extra, resp.Resp.Total, err = service.NewLogs(s.ctx).List(req)
	if err != nil {
		return nil, err
	}
	return
}
