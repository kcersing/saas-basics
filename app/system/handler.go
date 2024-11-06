package main

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	menu "rpc_gen/kitex_gen/system/menu"
	"system/biz/service"
)

// MenuServiceImpl implements the last service interface defined in the IDL.
type MenuServiceImpl struct{}

// CreateMenu implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) CreateMenu(ctx context.Context, req *menu.CreateOrUpdateMenuReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateMenuService(ctx).Run(req)

	return resp, err
}

// UpdateMenu implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) UpdateMenu(ctx context.Context, req *menu.CreateOrUpdateMenuReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateMenuService(ctx).Run(req)

	return resp, err
}

// DeleteMenu implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) DeleteMenu(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteMenuService(ctx).Run(req)

	return resp, err
}

// MenuByRole implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) MenuByRole(ctx context.Context, req *base.Empty) (resp *base.NilResponse, err error) {
	resp, err = service.NewMenuByRoleService(ctx).Run(req)

	return resp, err
}

// MenuLists implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) MenuLists(ctx context.Context, req *base.PageInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMenuListsService(ctx).Run(req)

	return resp, err
}

// MenuTree implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) MenuTree(ctx context.Context, req *base.PageInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMenuTreeService(ctx).Run(req)

	return resp, err
}

// CreateMenuParam implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) CreateMenuParam(ctx context.Context, req *menu.CreateOrUpdateMenuParamReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateMenuParamService(ctx).Run(req)

	return resp, err
}

// UpdateMenuParam implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) UpdateMenuParam(ctx context.Context, req *menu.CreateOrUpdateMenuParamReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateMenuParamService(ctx).Run(req)

	return resp, err
}

// DeleteMenuParam implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) DeleteMenuParam(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteMenuParamService(ctx).Run(req)

	return resp, err
}

// MenuParamListByMenuID implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) MenuParamListByMenuID(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMenuParamListByMenuIDService(ctx).Run(req)

	return resp, err
}
