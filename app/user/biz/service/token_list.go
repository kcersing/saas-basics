package service

import (
	"context"
	user "rpc_gen/kitex_gen/user"
	"user/biz/infra/service"
)

type TokenListService struct {
	ctx context.Context
} // NewTokenListService new TokenListService
func NewTokenListService(ctx context.Context) *TokenListService {
	return &TokenListService{ctx: ctx}
}

// Run create note info
func (s *TokenListService) Run(req *user.TokenListReq) (resp *user.TokenListResp, err error) {
	// Finish your business logic.

	resp.Extra, resp.Resp.Total, err = service.NewToken(s.ctx).List(req)
	if err != nil {
		return nil, err
	}
	return
}
