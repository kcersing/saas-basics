package service

import (
	"context"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"system/biz/infra/service"
)

type DictionaryListService struct {
	ctx context.Context
} // NewDictionaryListService new DictionaryListService
func NewDictionaryListService(ctx context.Context) *DictionaryListService {
	return &DictionaryListService{ctx: ctx}
}

// Run create note info
func (s *DictionaryListService) Run(req *dictionary.DictListReq) (resp *dictionary.DictionaryListResp, err error) {
	// Finish your business logic.

	resp.Extra, resp.Resp.Total, err = service.NewDictionary(s.ctx).List(req)
	if err != nil {
		return nil, err
	}
	return
}
