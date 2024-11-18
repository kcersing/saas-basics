package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
)

type PlaceListService struct {
	ctx context.Context
} // NewPlaceListService new PlaceListService
func NewPlaceListService(ctx context.Context) *PlaceListService {
	return &PlaceListService{ctx: ctx}
}

// Run create note info
func (s *PlaceListService) Run(req *sys.ListReq) (resp *sys.SysListResp, err error) {
	// Finish your business logic.

	return
}
