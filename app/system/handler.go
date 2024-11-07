package main

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	"rpc_gen/kitex_gen/system/dictionary"
	menu "rpc_gen/kitex_gen/system/menu"
	"rpc_gen/kitex_gen/system/role"
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
func (s *MenuServiceImpl) MenuByRole(ctx context.Context, req *base.IDReq) (resp []*menu.MenuInfoTree, err error) {
	resp, err = service.NewMenuByRoleService(ctx).Run(req)

	return resp, err
}

// MenuLists implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) MenuLists(ctx context.Context, req *base.PageInfoReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewMenuListsService(ctx).Run(req)

	return resp, err
}

// MenuTree implements the MenuServiceImpl interface.
func (s *MenuServiceImpl) MenuTree(ctx context.Context, req *base.PageInfoReq) (resp []*menu.Tree, err error) {
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

// CreateDictionary implements the DictionaryServiceImpl interface.
func (s *DictionaryServiceImpl) CreateDictionary(ctx context.Context, req *dictionary.DictionaryInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateDictionaryService(ctx).Run(req)

	return resp, err
}

// UpdateDictionary implements the DictionaryServiceImpl interface.
func (s *DictionaryServiceImpl) UpdateDictionary(ctx context.Context, req *dictionary.DictionaryInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateDictionaryService(ctx).Run(req)

	return resp, err
}

// DeleteDictionary implements the DictionaryServiceImpl interface.
func (s *DictionaryServiceImpl) DeleteDictionary(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteDictionaryService(ctx).Run(req)

	return resp, err
}

// DictionaryList implements the DictionaryServiceImpl interface.
func (s *DictionaryServiceImpl) DictionaryList(ctx context.Context, req *dictionary.DictionaryPageReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDictionaryListService(ctx).Run(req)

	return resp, err
}

// CreateDictionaryDetail implements the DictionaryServiceImpl interface.
func (s *DictionaryServiceImpl) CreateDictionaryDetail(ctx context.Context, req *dictionary.DictionaryDetail) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateDictionaryDetailService(ctx).Run(req)

	return resp, err
}

// UpdateDictionaryDetail implements the DictionaryServiceImpl interface.
func (s *DictionaryServiceImpl) UpdateDictionaryDetail(ctx context.Context, req *dictionary.DictionaryDetail) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateDictionaryDetailService(ctx).Run(req)

	return resp, err
}

// DeleteDictionaryDetail implements the DictionaryServiceImpl interface.
func (s *DictionaryServiceImpl) DeleteDictionaryDetail(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteDictionaryDetailService(ctx).Run(req)

	return resp, err
}

// DetailByDictionaryList implements the DictionaryServiceImpl interface.
func (s *DictionaryServiceImpl) DetailByDictionaryList(ctx context.Context, req *dictionary.DictionaryDetailReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDetailByDictionaryListService(ctx).Run(req)

	return resp, err
}

// CreateApi implements the ApisServiceImpl interface.
func (s *ApisServiceImpl) CreateApi(ctx context.Context, req *role.ApiInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewCreateApiService(ctx).Run(req)

	return resp, err
}

// UpdateApi implements the ApisServiceImpl interface.
func (s *ApisServiceImpl) UpdateApi(ctx context.Context, req *role.ApiInfo) (resp *base.NilResponse, err error) {
	resp, err = service.NewUpdateApiService(ctx).Run(req)

	return resp, err
}

// DeleteApi implements the ApisServiceImpl interface.
func (s *ApisServiceImpl) DeleteApi(ctx context.Context, req *base.IDReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteApiService(ctx).Run(req)

	return resp, err
}

// ApiList implements the ApisServiceImpl interface.
func (s *ApisServiceImpl) ApiList(ctx context.Context, req *role.ApiPageReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewApiListService(ctx).Run(req)

	return resp, err
}

// ApiTree implements the ApisServiceImpl interface.
func (s *ApisServiceImpl) ApiTree(ctx context.Context, req *role.ApiPageReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewApiTreeService(ctx).Run(req)

	return resp, err
}
