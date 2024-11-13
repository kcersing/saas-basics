package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"system/biz/infra/service"
)

type DeleteLogsService struct {
	ctx context.Context
} // NewDeleteLogsService new DeleteLogsService
func NewDeleteLogsService(ctx context.Context) *DeleteLogsService {
	return &DeleteLogsService{ctx: ctx}
}

// Run create note info
func (s *DeleteLogsService) Run(req *base.Ids) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	err = service.NewLogs(s.ctx).DeleteAll(req.Ids)
	if err != nil {
		return nil, err
	}
	return
}
