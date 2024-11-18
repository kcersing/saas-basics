package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
)

type StaffListService struct {
	ctx context.Context
} // NewStaffListService new StaffListService
func NewStaffListService(ctx context.Context) *StaffListService {
	return &StaffListService{ctx: ctx}
}

// Run create note info
func (s *StaffListService) Run(req *sys.ListReq) (resp *sys.SysListResp, err error) {
	// Finish your business logic.

	return
}
