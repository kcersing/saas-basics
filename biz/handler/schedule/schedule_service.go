// Code generated by hertz generator.

package schedule

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	schedule "saas/idl_gen/model/schedule"
)

// CreateScheduleUserTimePeriod .
// @router /service/schedule/create-user-time-period [POST]
func CreateScheduleUserTimePeriod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.UserTimePeriodReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateScheduleUserTimePeriod .
// @router /service/schedule/update-user-time-period [POST]
func UpdateScheduleUserTimePeriod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.UserTimePeriodReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CreateScheduleCours .
// @router /service/schedule/create-cours [POST]
func CreateScheduleCours(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.CreateOrUpdateScheduleCoursReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CreateScheduleLessons .
// @router /service/schedule/create-lessons [POST]
func CreateScheduleLessons(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.CreateOrUpdateScheduleLessonsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateScheduleStatus .
// @router /service/schedule/status [POST]
func UpdateScheduleStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ScheduleList .
// @router /service/schedule/list [POST]
func ScheduleList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ScheduleDateList .
// @router /service/schedule/date-list [POST]
func ScheduleDateList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetScheduleInfo .
// @router /service/schedule/info [POST]
func GetScheduleInfo(ctx context.Context, c *app.RequestContext) {
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

// CreateMemberSubscribe .
// @router /service/schedule/create-member-subscribe [POST]
func CreateMemberSubscribe(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.MemberSubscribeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ScheduleMemberList .
// @router /service/schedule/schedule-member-list [POST]
func ScheduleMemberList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleMemberListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateMemberStatus .
// @router /service/schedule/schedule-member-status [POST]
func UpdateMemberStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ScheduleMemberInfo .
// @router /service/schedule/schedule-member-info [POST]
func ScheduleMemberInfo(ctx context.Context, c *app.RequestContext) {
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

// ScheduleCoachList .
// @router /service/schedule/schedule-coach-list [POST]
func ScheduleCoachList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleCoachListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateCoachStatus .
// @router /service/schedule/schedule-coach-status [POST]
func UpdateCoachStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ScheduleCoachInfo .
// @router /service/schedule/schedule-coach-info [POST]
func ScheduleCoachInfo(ctx context.Context, c *app.RequestContext) {
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

// CreateScheduleCourse .
// @router /service/schedule/create-cours [POST]
func CreateScheduleCourse(ctx context.Context, c *app.RequestContext) {
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

// ScheduleInfo .
// @router /service/schedule/info [POST]
func ScheduleInfo(ctx context.Context, c *app.RequestContext) {
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

// CreateMemberSubscribeLessons .
// @router /service/schedule/create-member-subscribe-lessons [POST]
func CreateMemberSubscribeLessons(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.MemberSubscribeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateScheduleCoachStatus .
// @router /service/schedule/schedule-coach-status [POST]
func UpdateScheduleCoachStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
