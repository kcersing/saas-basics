package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
)

type PropertyListService struct {
	ctx context.Context
} // NewPropertyListService new PropertyListService
func NewPropertyListService(ctx context.Context) *PropertyListService {
	return &PropertyListService{ctx: ctx}
}

// Run create note info
func (s *PropertyListService) Run(req *sys.ListReq) (resp *sys.SysListResp, err error) {
	// Finish your business logic.

	return
}
