package main

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"system/biz/service"
)

// RoleServiceImpl implements the last service interface defined in the IDL.
type RoleServiceImpl struct{}

// CreateRole implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) CreateRole(ctx context.Context, req *auth.RoleInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateRoleService(ctx).Run(req)

	return resp, err
}

// UpdateRole implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) UpdateRole(ctx context.Context, req *auth.RoleInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateRoleService(ctx).Run(req)

	return resp, err
}

// DeleteRole implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) DeleteRole(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteRoleService(ctx).Run(req)

	return resp, err
}

// RoleByID implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) RoleByID(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewRoleByIDService(ctx).Run(req)

	return resp, err
}

// RoleList implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) RoleList(ctx context.Context, req *base.PageInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewRoleListService(ctx).Run(req)

	return resp, err
}

// UpdateRoleStatus implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateRoleStatusService(ctx).Run(req)

	return resp, err
}

// CreateAuthority implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) CreateAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateAuthorityService(ctx).Run(req)

	return resp, err
}

// UpdateApiAuthority implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) UpdateApiAuthority(ctx context.Context, req *auth.CreateOrUpdateApiAuthorityReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateApiAuthorityService(ctx).Run(req)

	return resp, err
}

// ApiAuthority implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) ApiAuthority(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewApiAuthorityService(ctx).Run(req)

	return resp, err
}

// CreateMenuAuthority implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) CreateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateMenuAuthorityService(ctx).Run(req)

	return resp, err
}

// UpdateMenuAuthority implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) UpdateMenuAuthority(ctx context.Context, req *auth.MenuAuthorityInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateMenuAuthorityService(ctx).Run(req)

	return resp, err
}

// MenuAuthority implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) MenuAuthority(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMenuAuthorityService(ctx).Run(req)

	return resp, err
}

// CreateApi implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) CreateApi(ctx context.Context, req *auth.ApiInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateApiService(ctx).Run(req)

	return resp, err
}

// UpdateApi implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) UpdateApi(ctx context.Context, req *auth.ApiInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateApiService(ctx).Run(req)

	return resp, err
}

// DeleteApi implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) DeleteApi(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteApiService(ctx).Run(req)

	return resp, err
}

// ApiList implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) ApiList(ctx context.Context, req *auth.ApiPageReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewApiListService(ctx).Run(req)

	return resp, err
}

// ApiTree implements the RoleServiceImpl interface.
func (s *RoleServiceImpl) ApiTree(ctx context.Context, req *auth.ApiPageReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewApiTreeService(ctx).Run(req)

	return resp, err
}
