package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
)

type VenueListService struct {
	ctx context.Context
} // NewVenueListService new VenueListService
func NewVenueListService(ctx context.Context) *VenueListService {
	return &VenueListService{ctx: ctx}
}

// Run create note info
func (s *VenueListService) Run(req *sys.ListReq) (resp *sys.SysListResp, err error) {
	// Finish your business logic.

	return
}
