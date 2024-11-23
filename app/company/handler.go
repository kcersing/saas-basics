package main

import (
	"company/biz/service"
	"context"
	base "rpc_gen/kitex_gen/base"
	contract "rpc_gen/kitex_gen/company/contract"
)

// CompanyServiceImpl implements the last service interface defined in the IDL.
type CompanyServiceImpl struct{}

// ContractList implements the CompanyServiceImpl interface.
func (s *CompanyServiceImpl) ContractList(ctx context.Context, req *contract.ContractListReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewContractListService(ctx).Run(req)

	return resp, err
}

// ContractCreate implements the CompanyServiceImpl interface.
func (s *CompanyServiceImpl) ContractCreate(ctx context.Context, req *contract.CreateOrUpdateContractReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewContractCreateService(ctx).Run(req)

	return resp, err
}

// ContractUpdate implements the CompanyServiceImpl interface.
func (s *CompanyServiceImpl) ContractUpdate(ctx context.Context, req *contract.CreateOrUpdateContractReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewContractUpdateService(ctx).Run(req)

	return resp, err
}

// ContractUpdateStatus implements the CompanyServiceImpl interface.
func (s *CompanyServiceImpl) ContractUpdateStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewContractUpdateStatusService(ctx).Run(req)

	return resp, err
}

// ContractByID implements the CompanyServiceImpl interface.
func (s *CompanyServiceImpl) ContractByID(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewContractByIDService(ctx).Run(req)

	return resp, err
}
