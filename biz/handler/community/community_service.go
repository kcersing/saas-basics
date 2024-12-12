// Code generated by hertz generator.

package community

import (
	"context"
	"github.com/skip2/go-qrcode"
	"saas/biz/infras/service"
	"saas/pkg/errno"
	"saas/pkg/utils"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	community "saas/idl_gen/model/community"
)

// CreateCommunity .
//
//	@Summary		创建运动社群 Summary
//	@Description	创建运动社群 Description
//	@Param			request	body		contest.ContestInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/create [POST]
func CreateCommunity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.CommunityInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewCommunity(ctx, c).CreateCommunity(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateCommunity .
//
//	@Summary		更新运动社群 Summary
//	@Description	更新运动社群 Description
//	@Param			request	body		community.ContestInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/update [POST]
func UpdateCommunity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.CommunityInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	err = service.NewCommunity(ctx, c).UpdateCommunity(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// CommunityInfo .
//
//	@Summary		运动社群信息 Summary
//
//	@Description	运动社群信息 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/info [POST]
func CommunityInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewCommunity(ctx, c).CommunityInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// CommunityList .
//
//	@Summary		运动社群列表 Summary
//
//	@Description	运动社群列表 Description
//	@Param			request	body		community.CommunityListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/list [POST]
func CommunityList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.CommunityListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewCommunity(ctx, c).CommunityList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// UpdateCommunityStatus .
//
//	@Summary		更新运动社群状态 Summary
//
//	@Description	更新运动社群状态 Description
//	@Param			request	body		base.StatusCodeReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/status [POST]
func UpdateCommunityStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewCommunity(ctx, c).UpdateCommunityStatus(req.ID, req.Status)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateCommunityShow .
//
//	@Summary		更新运动社群展示状态 Summary
//
//	@Description	更新运动社群展示状态 Description
//	@Param			request	body		base.StatusCodeReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/show [POST]
func UpdateCommunityShow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewCommunity(ctx, c).UpdateCommunityShow(req.ID, req.Status)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// CreateParticipant .
//
//	@Summary		创建运动社群报名 Summary
//
//	@Description	创建运动社群报名 Description
//	@Param			request	body		community.ParticipantInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/participant/create [POST]
func CreateParticipant(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.ParticipantInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewCommunity(ctx, c).CreateParticipant(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateParticipant .
//
//	@Summary		更新运动社群报名 Summary
//
//	@Description	更新运动社群报名 Description
//	@Param			request	body		community.ParticipantInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/participant/update [POST]
func UpdateParticipant(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.ParticipantInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewCommunity(ctx, c).UpdateParticipant(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// CommunityParticipantInfo .
//
//	@Summary		运动社群报名信息 Summary
//
//	@Description	运动社群报名信息 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/participant/info [POST]
func CommunityParticipantInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewCommunity(ctx, c).ParticipantInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// ParticipantListList .
//
//	@Summary		运动社群报名列表 Summary
//
//	@Description	运动社群报名列表 Description
//	@Param			request	body		community.ParticipantListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/participant/list [POST]
func ParticipantListList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.ParticipantListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewCommunity(ctx, c).ParticipantList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// UpdateParticipantStatus .
//
//	@Summary		更新运动社群报名状态 Summary
//
//	@Description	更新运动社群报名状态 Description
//	@Param			request	body		base.StatusCodeReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/participant/status [POST]
func UpdateParticipantStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewCommunity(ctx, c).UpdateParticipantStatus(req.ID, req.Status)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// ParticipantListListExport .
//
//	@Summary		运动社群报名列表导出 Summary
//
//	@Description	运动社群报名列表导出 Description
//	@Param			request	body		community.ParticipantListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/participant/export [POST]
func ParticipantListListExport(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.ParticipantListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	export, err := service.NewCommunity(ctx, c).ParticipantListExport(req)
	if err != nil {
		utils.SendResponse(c, errno.DirtyData, nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, map[string]string{
		"url": export,
	}, 0, "")
}

// ResultsUpload .
//
//	@Summary		运动社群报名结果上传 Summary
//
//	@Description	运动社群报名结果上传 Description
//	@Param			request	body		community.ResultsUploadReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/results-upload [POST]
func ResultsUpload(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.ResultsUploadReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// PromotionalLinks .
//
//	@Summary		训练营推广链接 Summary
//
//	@Description	训练营推广链接 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/promotional-links [POST]
func PromotionalLinks(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	qrCode, err := qrcode.Encode("weixin://dl/business/?t= *TICKET*&Community_id="+strconv.FormatInt(req.ID, 10), qrcode.Medium, 256)
	if err != nil {
		utils.SendResponse(c, errno.ServiceErr, nil, 0, "")
	}
	type img struct {
		Bs64 []byte `json:"bs64"`
	}
	utils.SendResponse(c, errno.Success, img{Bs64: qrCode}, 0, "")
}

// DelCommunity .
//
//	@Summary		删除运动社群 Summary
//
//	@Description	删除运动社群 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/del [POST]
func DelCommunity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewCommunity(ctx, c).DelCommunity(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// ParticipantFinish .
//
//	@Summary		运动社群报名完成 Summary
//
//	@Description	运动社群报名完成 Description
//	@Param			request	body		community.ParticipantFinishReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/community/participant/finish [POST]
func ParticipantFinish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req community.ParticipantFinishReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}
