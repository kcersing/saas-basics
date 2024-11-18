package service

import (
	"context"
	sys "rpc_gen/kitex_gen/system/sys"
)

type ContractListService struct {
	ctx context.Context
} // NewContractListService new ContractListService
func NewContractListService(ctx context.Context) *ContractListService {
	return &ContractListService{ctx: ctx}
}

// Run create note info
func (s *ContractListService) Run(req *sys.ListReq) (resp *sys.SysListResp, err error) {
	// Finish your business logic.
	//resp.Extra, resp.Resp.Total, err = service.NewApi(s.ctx).List(*req)
	//if err != nil {
	//	return nil, err
	//}
	//return

}
