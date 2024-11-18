package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"system/biz/infra/service"
)

type DeleteDictionaryDetailService struct {
	ctx context.Context
} // NewDeleteDictionaryDetailService new DeleteDictionaryDetailService
func NewDeleteDictionaryDetailService(ctx context.Context) *DeleteDictionaryDetailService {
	return &DeleteDictionaryDetailService{ctx: ctx}
}

// Run create note info
func (s *DeleteDictionaryDetailService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewDictionary(s.ctx).DeleteDetail(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
