package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"system/biz/infra/service"
)

type DeleteDictionaryService struct {
	ctx context.Context
} // NewDeleteDictionaryService new DeleteDictionaryService
func NewDeleteDictionaryService(ctx context.Context) *DeleteDictionaryService {
	return &DeleteDictionaryService{ctx: ctx}
}

// Run create note info
func (s *DeleteDictionaryService) Run(req *base.IDReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewDictionary(s.ctx).Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return
}
