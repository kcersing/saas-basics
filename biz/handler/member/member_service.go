// Code generated by hertz generator.

package member

import (
	"context"
	"saas/biz/infras/service"
	"saas/pkg/errno"
	"saas/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	member "saas/idl_gen/model/member"
)

// CreateMember .
//
//	@Summary		创建会员 Summary
//	@Description	创建会员 Description
//	@Param			request	body		member.CreateOrUpdateMemberReq	true "query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/member/create [POST]
func CreateMember(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.CreateOrUpdateMemberReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewMember(ctx, c).CreateMember(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateMember .
//
//	@Summary		更新会员 Summary
//	@Description	更新会员 Description
//	@Param			request	body		member.CreateOrUpdateMemberReq	true "query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/member/update [POST]
func UpdateMember(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.CreateOrUpdateMemberReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewMember(ctx, c).UpdateMember(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// MemberInfo .
//
//	@Summary		会员详情 Summary
//	@Description	会员详情 Description
//	@Param			request	body		base.IDReq	true "query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/member/info [POST]
func MemberInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewMember(ctx, c).MemberInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// MemberFullList .
//
//	@Summary		会员列表 Summary
//	@Description	会员列表 Description
//	@Param			request	body		member.MemberListReq	true "query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/member/full-list [POST]
func MemberFullList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewMember(ctx, c).MemberFullList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// MemberPotentialList .
//
//	@Summary		潜在会员列表 Summary
//	@Description	潜在会员列表 Description
//	@Param			request	body		member.MemberListReq	true "query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/member/potential-list [POST]
func MemberPotentialList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewMember(ctx, c).MemberPotentialList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// UpdateMemberStatus .
//
//	@Summary		更新会员状态 Summary
//	@Description	更新会员状态 Description
//	@Param			request	body		base.StatusCodeReq	true "query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/member/status [POST]
func UpdateMemberStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewMember(ctx, c).UpdateMemberStatus(req.ID, req.Status)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateMemberFollow .
//
//	@Summary		更新会员关注 Summary
//	@Description	更新会员关注 Description
//	@Param			request	body		member.UpdateMemberFollowReq	true "query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/member/update-follow [POST]
func UpdateMemberFollow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.UpdateMemberFollowReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewMember(ctx, c).UpdateMemberFollow(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// MemberContractList .
//
//	@Summary		会员合同列表 Summary
//	@Description	会员合同列表 Description
//	@Param			request	body		member.MemberContractListReq	true "query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/member/contract-list [POST]
func MemberContractList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberContractListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	list, total, err := service.NewMember(ctx, c).ContractList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// MemberFullListExport .
// @router /service/member/full-list/export [POST]
func MemberFullListExport(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MemberPotentialListExport .
// @router /service/member/potential-list/export [POST]
func MemberPotentialListExport(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
