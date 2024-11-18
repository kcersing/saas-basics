package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"system/biz/infra/service"
)

type CreateDictionaryService struct {
	ctx context.Context
} // NewCreateDictionaryService new CreateDictionaryService
func NewCreateDictionaryService(ctx context.Context) *CreateDictionaryService {
	return &CreateDictionaryService{ctx: ctx}
}

// Run create note info
func (s *CreateDictionaryService) Run(req *dictionary.DictionaryInfo) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	err = service.NewDictionary(s.ctx).Create(req)
	if err != nil {
		return nil, err
	}
	return
}
