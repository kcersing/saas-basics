package main

import (
	"context"
	base "rpc_gen/kitex_gen/base"
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
