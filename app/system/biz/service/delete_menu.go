package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	base "rpc_gen/kitex_gen/base"
	admin "system/biz/infra/service"
)

type DeleteMenuService struct {
	ctx context.Context
} // NewDeleteMenuService new DeleteMenuService
func NewDeleteMenuService(ctx context.Context) *DeleteMenuService {
	return &DeleteMenuService{ctx: ctx}
}

// Run create note info
func (s *DeleteMenuService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	err = admin.NewMenu(s.ctx).Delete(req.Id)
	if err != nil {
		return resp, kerrors.NewBizStatusError(40001, err.Error())
	}
	return resp, nil
}
