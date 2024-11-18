package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
)

type ProductListService struct {
	ctx context.Context
} // NewProductListService new ProductListService
func NewProductListService(ctx context.Context) *ProductListService {
	return &ProductListService{ctx: ctx}
}

// Run create note info
func (s *ProductListService) Run(req *sys.ListReq) (resp *sys.SysListResp, err error) {
	// Finish your business logic.

	return
}
