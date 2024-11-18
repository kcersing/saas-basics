package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"system/biz/infra/service"
)

type CreateDictionaryDetailService struct {
	ctx context.Context
} // NewCreateDictionaryDetailService new CreateDictionaryDetailService
func NewCreateDictionaryDetailService(ctx context.Context) *CreateDictionaryDetailService {
	return &CreateDictionaryDetailService{ctx: ctx}
}

// Run create note info
func (s *CreateDictionaryDetailService) Run(req *dictionary.DictionaryDetail) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewDictionary(s.ctx).CreateDetail(req)
	if err != nil {
		return nil, err
	}
	return
}
