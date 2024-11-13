package main

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"rpc_gen/kitex_gen/system/auth"
	menu "rpc_gen/kitex_gen/system/menu"
	"system/biz/service"
)

// SystemServiceImpl implements the last service interface defined in the IDL.
type SystemServiceImpl struct{}

// MenuAuth implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) MenuAuth(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMenuAuthService(ctx).Run(req)

	return resp, err
}

// ApiList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) ApiList(ctx context.Context, req *menu.ApiPageReq) (resp []*menu.ApiInfo, err error) {
	resp, err = service.NewApiListService(ctx).Run(req)

	return resp, err
}

// ApiTree implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) ApiTree(ctx context.Context, req *menu.ApiPageReq) (resp []*base.Tree, err error) {
	resp, err = service.NewApiTreeService(ctx).Run(req)

	return resp, err
}

// MenuByRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) MenuByRole(ctx context.Context, req *base.IDReq) (resp []*menu.MenuInfoTree, err error) {
	resp, err = service.NewMenuByRoleService(ctx).Run(req)

	return resp, err
}

// MenuLists implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) MenuLists(ctx context.Context, req *base.PageInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMenuListsService(ctx).Run(req)

	return resp, err
}

// MenuTree implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) MenuTree(ctx context.Context, req *base.PageInfoReq) (resp []*base.Tree, err error) {
	resp, err = service.NewMenuTreeService(ctx).Run(req)

	return resp, err
}

// MenuRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) MenuRole(ctx context.Context, req *base.IDReq) (resp []*menu.MenuInfoTree, err error) {
	resp, err = service.NewMenuRoleService(ctx).Run(req)

	return resp, err
}

// CreateRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateRole(ctx context.Context, req *auth.RoleInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateRoleService(ctx).Run(req)

	return resp, err
}

// UpdateRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateRole(ctx context.Context, req *auth.RoleInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateRoleService(ctx).Run(req)

	return resp, err
}

// DeleteRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteRole(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteRoleService(ctx).Run(req)

	return resp, err
}

// RoleByID implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) RoleByID(ctx context.Context, req *base.IDReq) (resp *auth.RoleInfo, err error) {
	resp, err = service.NewRoleByIDService(ctx).Run(req)

	return resp, err
}

// CreateMenuAuth implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateMenuAuthService(ctx).Run(req)

	return resp, err
}

// UpdateMenuAuth implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateMenuAuthService(ctx).Run(req)

	return resp, err
}

// RoleList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) RoleList(ctx context.Context, req *base.PageInfoReq) (resp []*auth.RoleInfo, err error) {
	resp, err = service.NewRoleListService(ctx).Run(req)

	return resp, err
}

// UpdateRoleStatus implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateRoleStatusService(ctx).Run(req)

	return resp, err
}

// CreateAuth implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateAuthService(ctx).Run(req)

	return resp, err
}

// UpdateApiAuth implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateApiAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateApiAuthService(ctx).Run(req)

	return resp, err
}

// ApiAuth implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) ApiAuth(ctx context.Context, req *base.IDReq) (resp []*auth.ApiAuthInfo, err error) {
	resp, err = service.NewApiAuthService(ctx).Run(req)

	return resp, err
}

// GetLogsList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) GetLogsList(ctx context.Context, req *auth.LogsListReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewGetLogsListService(ctx).Run(req)

	return resp, err
}

// DeleteLogs implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteLogs(ctx context.Context, req *base.Empty) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteLogsService(ctx).Run(req)

	return resp, err
}
