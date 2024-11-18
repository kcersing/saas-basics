package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"system/biz/infra/service"
)

type UpdateDictionaryService struct {
	ctx context.Context
} // NewUpdateDictionaryService new UpdateDictionaryService
func NewUpdateDictionaryService(ctx context.Context) *UpdateDictionaryService {
	return &UpdateDictionaryService{ctx: ctx}
}

// Run create note info
func (s *UpdateDictionaryService) Run(req *dictionary.DictionaryInfo) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewDictionary(s.ctx).Update(req)
	if err != nil {
		return nil, err
	}
	return
}
