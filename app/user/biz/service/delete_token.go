package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"user/biz/infra/service"
)

type DeleteTokenService struct {
	ctx context.Context
} // NewDeleteTokenService new DeleteTokenService
func NewDeleteTokenService(ctx context.Context) *DeleteTokenService {
	return &DeleteTokenService{ctx: ctx}
}

// Run create note info
func (s *DeleteTokenService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	err = service.NewToken(s.ctx).Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
