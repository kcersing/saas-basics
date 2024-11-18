package service

import (
	"context"
	dictionary "rpc_gen/kitex_gen/system/dictionary"
	"system/biz/infra/service"
)

type DetailByDictionaryListService struct {
	ctx context.Context
} // NewDetailByDictionaryListService new DetailByDictionaryListService
func NewDetailByDictionaryListService(ctx context.Context) *DetailByDictionaryListService {
	return &DetailByDictionaryListService{ctx: ctx}
}

// Run create note info
func (s *DetailByDictionaryListService) Run(req *dictionary.DetailListReq) (resp *dictionary.DetailByDictionaryListResp, err error) {
	// Finish your business logic.

	resp.Extra, resp.Resp.Total, err = service.NewDictionary(s.ctx).DetailListByDict(req)
	if err != nil {
		return nil, err
	}
	return
}
