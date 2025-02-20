// Code generated by hertz generator.

package wx

import (
	"context"
	"saas/biz/infras/service"
	"saas/idl_gen/model/bootcamp"
	"saas/idl_gen/model/contest"
	"saas/idl_gen/model/user"
	venue "saas/idl_gen/model/venue"
	"saas/pkg/errno"
	"saas/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	community "saas/idl_gen/model/community"
	member "saas/idl_gen/model/member"
	memberProduct "saas/idl_gen/model/member/memberProduct"
	order "saas/idl_gen/model/order"
	product "saas/idl_gen/model/product"
	schedule "saas/idl_gen/model/schedule"
	wx "saas/idl_gen/model/wx"
)

// MemberInfo .
//
//	@Summary		会员信息 Summary
//	@Description	会员信息 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/info [POST]
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

// MemberProductInfo .
//
//	@Summary		会员产品信息 Summary
//	@Description	会员产品信息 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/product-info [POST]
func MemberProductInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewMemberProduct(ctx, c).MemberProductInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// MemberProductList .
//
//	@Summary		会员产品列表 Summary
//	@Description	会员产品列表 Description
//	@Param			request	body		memberProduct.MemberProductListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/product-list [POST]
func MemberProductList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req memberProduct.MemberProductListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	list, total, err := service.NewMemberProduct(ctx, c).MemberProductList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// MemberContestList .
//
//	@Summary		会员竞赛列表 Summary
//	@Description	会员竞赛列表 Description
//	@Param			request	body		member.MemberContestListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/contest-list [POST]
func MemberContestList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberContestListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	list, total, err := service.NewMember(ctx, c).ContestList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// MemberBootcampList .
//
//	@Summary		会员训练营列表 Summary
//	@Description	会员训练营列表 Description
//	@Param			request	body		member.MemberBootcampListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/bootcamp-list [POST]
func MemberBootcampList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberBootcampListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewMember(ctx, c).BootcampList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// MemberCommunityList .
//
//	@Summary		会员社群列表 Summary
//	@Description	会员社群列表 Description
//	@Param			request	body		member.MemberCommunityListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/community-list [POST]
func MemberCommunityList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberCommunityListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewMember(ctx, c).CommunityList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// MemberContractList .
//
//	@Summary		会员合同列表 Summary
//	@Description	会员合同列表 Description
//	@Param			request	body		member.MemberContractListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/contract-list [POST]
func MemberContractList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.MemberContractListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ScheduleMemberList .
//
//	@Summary		会员约课列表 Summary
//	@Description	会员约课列表 Description
//	@Param			request	body		schedule.ScheduleMemberListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/schedule-member-list [POST]
func ScheduleMemberList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleMemberListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewSchedule(ctx, c).ScheduleMemberList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// ContestList .
//
//	@Summary		竞赛列表 Summary
//	@Description	竞赛列表 Description
//	@Param			request	body		contest.ContestListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/contest-list [POST]
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

// ContestInfo .
//
//	@Summary		竞赛详情 Summary
//	@Description	竞赛详情 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/contest-info [POST]
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

// JoinContest .
//
//	@Summary		报名竞赛 Summary
//	@Description	报名竞赛 Description
//	@Param			request	body		wx.JoinContestReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/join-contest [POST]
func JoinContest(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.JoinContestReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// BootcampList .
//
//	@Summary		训练营列表 Summary
//	@Description	训练营列表 Description
//	@Param			request	body		bootcamp.BootcampListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/bootcamp-list [POST]
func BootcampList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req bootcamp.BootcampListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewBootcamp(ctx, c).BootcampList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// BootcampInfo .
//
//	@Summary		训练营详情 Summary
//	@Description	训练营详情 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/bootcamp-info [POST]
func BootcampInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewBootcamp(ctx, c).BootcampInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// JoinBootcamp .
//
//	@Summary		报名训练营 Summary
//	@Description	报名训练营 Description
//	@Param			request	body		wx.JoinBootcampReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/join-bootcamp [POST]
func JoinBootcamp(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.JoinBootcampReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommunityList .
//
//	@Summary		社群列表 Summary
//	@Description	社群列表 Description
//	@Param			request	body		community.CommunityListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/community-list [POST]
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

// CommunityInfo .
//
//	@Summary		社群详情 Summary
//	@Description	社群详情 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/community-info [POST]
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

// JoinCommunity .
//
//	@Summary		加入社群 Summary
//	@Description	加入社群 Description
//	@Param			request	body		wx.JoinCommunityReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/join-community [POST]
func JoinCommunity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.JoinCommunityReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// VenueList .
//
//	@Summary		场馆列表 Summary
//	@Description	场馆列表 Description
//	@Param			request	body		venue.VenueListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/venue-list [POST]
func VenueList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.VenueListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewVenue(ctx, c).List(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// VenueInfo .
//
//	@Summary		场馆详情 Summary
//	@Description	场馆详情 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/venue-info [POST]
func VenueInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	info, err := service.NewVenue(ctx, c).VenueInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// ScanQR .
//
//	@Summary		生成扫码信息 Summary
//	@Description	生成扫码信息 Description
//	@Param			request	body		wx.ScanQRReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/scan-QR [POST]
func ScanQR(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.ScanQRReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// Activation .
// @router /service/wx/member/activation [POST]
func Activation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CreateMemberScheduleCourse .
//
//	@Summary		会员约私教课 Summary
//	@Description	会员约私教课 Description
//	@Param			request	body		schedule.CreateOrUpdateScheduleCourseReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/create-member-schedule-course [POST]
func CreateMemberScheduleCourse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.CreateOrUpdateScheduleCourseReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CreateMemberScheduleCourseLessons .
//
//	@Summary		会员约团课 Summary
//	@Description	会员约团课 Description
//	@Param			request	body		wx.CreateMemberScheduleLessonsReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/create-member-schedule-lessons [POST]
func CreateMemberScheduleCourseLessons(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.CreateMemberScheduleLessonsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ProductInfo .
//
//	@Summary		产品详情 Summary
//	@Description	产品详情 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/product-info [POST]
func ProductInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewProduct(ctx, c).ProductInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// Buy .
//
//	@Summary		下单 Summary
//	@Description	下单 Description
//	@Param			request	body		order.BuyReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/buy [POST]
func Buy(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.BuyReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CoachList .
//
//	@Summary		教练列表 Summary
//	@Description	教练列表 Description
//	@Param			request	body		wx.CoachListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/coach-list [POST]
func CoachList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.CoachListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewUser(ctx, c).List(user.UserListReq{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Functions: req.Functions,
		Status:    req.Status,
		VenueId:   req.VenueId,
		TagId:     req.TagId,
		Type:      1,
	})
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// CoachInfo .
//
//	@Summary		教练详情 Summary
//	@Description	教练详情 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/coach-info [POST]
func CoachInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewUser(ctx, c).Info(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// PlaceInfo .
//
//	@Summary		场地详情 Summary
//	@Description	场地详情 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/place-info [POST]
func PlaceInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewVenue(ctx, c).PlaceInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// CreatePlaceSchedule .
//
//	@Summary		场地预约 Summary
//	@Description	场地预约 Description
//	@Param			request	body		wx.CreatePlaceScheduleReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/create-place-schedule [POST]
func CreatePlaceSchedule(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.CreatePlaceScheduleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MyMember .
//
//	@Summary		我的会员 Summary
//	@Description	我的会员 Description
//	@Param			request	body		wx.MyMemberReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/staff/my-member [POST]
func MyMember(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.MyMemberReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CoachScheduleList .
//
//	@Summary		教练上课列表 Summary
//	@Description	教练上课列表 Description
//	@Param			request	body		schedule.ScheduleCoachListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/staff/schedule-member-list [POST]
func CoachScheduleList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleCoachListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewSchedule(ctx, c).ScheduleCoachList(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// SignMemberSchedule .
//
//	@Summary		会员签到 Summary
//	@Description	会员签到 Description
//	@Param			request	body		wx.SignMemberScheduleReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/staff/sign-member-schedule [POST]
func SignMemberSchedule(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.SignMemberScheduleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// SignStaffSchedule .
//
//	@Summary		教练签到 Summary
//	@Description	教练签到 Description
//	@Param			request	body		wx.SignStaffScheduleReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/staff/sign-staff-schedule [POST]
func SignStaffSchedule(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.SignStaffScheduleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MemberRegister .
//
//	@Summary		会员注册 Summary
//	@Description	会员注册 Description
//	@Param			request	body		wx.MemberRegisterReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/register [POST]
func MemberRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.MemberRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MemberCaptcha .
//
//	@Summary		会员验证码 Summary
//	@Description	会员验证码 Description
//	@Param			request	body		wx.MemberCaptchaReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/captcha [POST]
func MemberCaptcha(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.MemberCaptchaReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// StaffCaptcha .
//
//	@Summary		教练验证码 Summary
//	@Description	教练验证码 Description
//	@Param			request	body		wx.StaffCaptchaReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/staff/captcha [POST]
func StaffCaptcha(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.StaffCaptchaReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MemberOrderList .
//
//	@Summary		会员订单列表 Summary
//	@Description	会员订单列表 Description
//	@Param			request	body		wx.MemberOrderListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/order-list [POST]
func MemberOrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wx.MemberOrderListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewOrder(ctx, c).List(&order.ListOrderReq{
		Page:              req.Page,
		PageSize:          req.PageSize,
		VenueId:           req.VenueId,
		Status:            req.Status,
		StartCompletionAt: req.StartCompletionAt,
		EndCompletionAt:   req.EndCompletionAt,
		ProductType:       req.ProductType,
		Nature:            req.Nature,
		Name:              req.Name,
		MemberName:        req.MemberName,
		MemberId:          req.MemberId,
	})
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// ProductList .
//
//	@Summary		产品列表 Summary
//	@Description	产品列表 Description
//	@Param			request	body		product.ProductListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/product-list [POST]
func ProductList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ProductListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	list, total, err := service.NewProduct(ctx, c).ProductList(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// PlaceList .
//
//	@Summary		场地列表 Summary
//	@Description	场地列表 Description
//	@Param			request	body		venue.VenuePlaceListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//
// @router /service/wx/member/place-list [POST]
func PlaceList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.VenuePlaceListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewVenue(ctx, c).PlaceList(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// Pay .
// @router /service/wx/member/pay [POST]
func Pay(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderPay
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
