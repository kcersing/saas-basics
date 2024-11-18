package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	user "rpc_gen/kitex_gen/user"
)

type SetDefaultVenueService struct {
	ctx context.Context
} // NewSetDefaultVenueService new SetDefaultVenueService
func NewSetDefaultVenueService(ctx context.Context) *SetDefaultVenueService {
	return &SetDefaultVenueService{ctx: ctx}
}

// Run create note info
func (s *SetDefaultVenueService) Run(req *user.SetDefaultVenueReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
