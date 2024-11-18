package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
)

type PropertyTypeService struct {
	ctx context.Context
} // NewPropertyTypeService new PropertyTypeService
func NewPropertyTypeService(ctx context.Context) *PropertyTypeService {
	return &PropertyTypeService{ctx: ctx}
}

// Run create note info
func (s *PropertyTypeService) Run(req *sys.ListReq) (resp *sys.SysListResp, err error) {
	// Finish your business logic.

	return
}
