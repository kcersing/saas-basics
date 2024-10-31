package member

import (
	"context"

	"frontend/biz/service"
	"frontend/biz/utils"
	base "frontend/hertz_gen/base"
	member "frontend/member"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateMember .
// @router /api/admin/member/create [POST]
func CreateMember(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.CreateOrUpdateMemberReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewCreateMemberService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateMember .
// @router /api/admin/member/update [POST]
func UpdateMember(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.CreateOrUpdateMemberReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewUpdateMemberService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberInfo .
// @router /api/admin/member/info [POST]
func MemberInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberInfoService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberList .
// @router /api/admin/member/list [POST]
func MemberList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateMemberStatus .
// @router /api/admin/member/status [POST]
func UpdateMemberStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewUpdateMemberStatusService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberSearch .
// @router /api/admin/member/search [POST]
func MemberSearch(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberSearchReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberSearchService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberProductList .
// @router /api/admin/member/product-list [POST]
func MemberProductList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberProductListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberProductListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberPropertyList .
// @router /api/admin/member/property-list [POST]
func MemberPropertyList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberPropertyListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberPropertyListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberProductDetail .
// @router /api/admin/member/product-detail [POST]
func MemberProductDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberProductDetailService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberPropertyDetail .
// @router /api/admin/member/property-detail [POST]
func MemberPropertyDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberPropertyDetailService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberPropertyUpdate .
// @router /api/admin/member/property-update [POST]
func MemberPropertyUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberPropertyListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberPropertyUpdateService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberContractList .
// @router /api/admin/member/contract-list [POST]
func MemberContractList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberPropertyListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberContractListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberProductSearch .
// @router /api/admin/member/search-product [POST]
func MemberProductSearch(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberProductSearchReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberProductSearchService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MemberPropertySearch .
// @router /api/admin/member/search-property [POST]
func MemberPropertySearch(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberPropertySearchReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &base.NilResponse{}
	resp, err = service.NewMemberPropertySearchService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
