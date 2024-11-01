package main

import (
	"context"
	"member/biz/service"
	base "rpc_gen/kitex_gen/base"
	member "rpc_gen/kitex_gen/member"
)

// MemberServiceImpl implements the last service interface defined in the IDL.
type MemberServiceImpl struct{}

// CreateMember implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) CreateMember(ctx context.Context, req *member.CreateOrUpdateMemberReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateMemberService(ctx).Run(req)

	return resp, err
}

// UpdateMember implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) UpdateMember(ctx context.Context, req *member.CreateOrUpdateMemberReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateMemberService(ctx).Run(req)

	return resp, err
}

// MemberInfo implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberInfo(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberInfoService(ctx).Run(req)

	return resp, err
}

// MemberList implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberList(ctx context.Context, req *member.MemberListReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberListService(ctx).Run(req)

	return resp, err
}

// UpdateMemberStatus implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateMemberStatusService(ctx).Run(req)

	return resp, err
}

// MemberSearch implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberSearch(ctx context.Context, req *member.MemberSearchReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberSearchService(ctx).Run(req)

	return resp, err
}

// MemberProductList implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberProductList(ctx context.Context, req *member.MemberProductListReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberProductListService(ctx).Run(req)

	return resp, err
}

// MemberPropertyList implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberPropertyList(ctx context.Context, req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberPropertyListService(ctx).Run(req)

	return resp, err
}

// MemberProductDetail implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberProductDetail(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberProductDetailService(ctx).Run(req)

	return resp, err
}

// MemberPropertyDetail implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberPropertyDetail(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberPropertyDetailService(ctx).Run(req)

	return resp, err
}

// MemberPropertyUpdate implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberPropertyUpdate(ctx context.Context, req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberPropertyUpdateService(ctx).Run(req)

	return resp, err
}

// MemberContractList implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberContractList(ctx context.Context, req *member.MemberPropertyListReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberContractListService(ctx).Run(req)

	return resp, err
}

// MemberProductSearch implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberProductSearch(ctx context.Context, req *member.MemberProductSearchReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberProductSearchService(ctx).Run(req)

	return resp, err
}

// MemberPropertySearch implements the MemberServiceImpl interface.
func (s *MemberServiceImpl) MemberPropertySearch(ctx context.Context, req *member.MemberPropertySearchReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMemberPropertySearchService(ctx).Run(req)

	return resp, err
}
