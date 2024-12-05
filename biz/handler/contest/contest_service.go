// Code generated by hertz generator.

package contest

import (
	"context"
	"saas/biz/infras/service"
	"saas/pkg/errno"
	"saas/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	contest "saas/idl_gen/model/contest"
)

// CreateContest .
//
//	@Summary		创建比赛 Summary
//	@Description	创建比赛 Description
//	@Param			request	body		contest.ContestInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/contest/create [POST]
func CreateContest(ctx context.Context, c *app.RequestContext) {
	var err error
	var req contest.ContestInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewContest(ctx, c).CreateContest(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateContest .
//
//	@Summary		更新比赛 Summary
//	@Description	更新比赛 Description
//	@Param			request	body		contest.ContestInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/contest/update [POST]
func UpdateContest(ctx context.Context, c *app.RequestContext) {
	var err error
	var req contest.ContestInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewContest(ctx, c).UpdateContest(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// ContestInfo .
//
//	@Summary		比赛信息 Summary
//	@Description	比赛信息 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/contest/info [POST]
func ContestInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewContest(ctx, c).ContestInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// ContestList .
//
//	@Summary		比赛列表 Summary
//	@Description	比赛列表 Description
//	@Param			request	body		contest.ContestListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/contest/list [POST]
func ContestList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req contest.ContestListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewContest(ctx, c).ContestList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// UpdateMemberStatus .
//
//	@Summary		更新比赛状态 Summary
//	@Description	更新比赛状态 Description
//	@Param			request	body		base.StatusCodeReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/contest/status [POST]
func UpdateMemberStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewContest(ctx, c).UpdateParticipantStatus(req.ID, req.Status)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// CreateParticipant .
//
//	@Summary		添加参赛人 Summary
//	@Description	添加参赛人 Description
//	@Param			request	body		contest.ParticipantInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/participant/create [POST]
func CreateParticipant(ctx context.Context, c *app.RequestContext) {
	var err error
	var req contest.ParticipantInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewContest(ctx, c).CreateParticipant(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateParticipant .
//
//	@Summary		更新参赛人 Summary
//	@Description	更新参赛人 Description
//	@Param			request	body		contest.ParticipantInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/participant/update [POST]
func UpdateParticipant(ctx context.Context, c *app.RequestContext) {
	var err error
	var req contest.ParticipantInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewContest(ctx, c).UpdateParticipant(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// ContestParticipantInfo .
//
//	@Summary		参赛人信息 Summary
//	@Description	参赛人信息 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/participant/info [POST]
func ContestParticipantInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewContest(ctx, c).ParticipantInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// ParticipantListList .
//
//	@Summary		参赛人列表 Summary
//	@Description	参赛人列表 Description
//	@Param			request	body		contest.ParticipantListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/participant/list [POST]
func ParticipantListList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req contest.ParticipantListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewContest(ctx, c).ParticipantList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// UpdateParticipantStatus .
//
//	@Summary		更新参赛人状态 Summary
//	@Description	更新参赛人状态 Description
//	@Param			request	body		base.StatusCodeReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/participant/status [POST]
func UpdateParticipantStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	err = service.NewContest(ctx, c).UpdateParticipantStatus(req.ID, req.Status)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}
