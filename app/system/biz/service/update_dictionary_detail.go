package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"system/biz/infra/service"
)

type UpdateDictionaryDetailService struct {
	ctx context.Context
} // NewUpdateDictionaryDetailService new UpdateDictionaryDetailService
func NewUpdateDictionaryDetailService(ctx context.Context) *UpdateDictionaryDetailService {
	return &UpdateDictionaryDetailService{ctx: ctx}
}

// Run create note info
func (s *UpdateDictionaryDetailService) Run(req *dictionary.DictionaryDetail) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewDictionary(s.ctx).UpdateDetail(req)
	if err != nil {
		return nil, err
	}
	return
}
