// Code generated by hertz generator.

package member

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	member "saas/app/admin/idl_gen/model/admin/member"
	base "saas/app/admin/idl_gen/model/base"
)

// CreateMember .
// @router /api/admin/member/create [POST]
func CreateMember(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.CreateOrUpdateMemberReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateMember .
// @router /api/admin/member/update [POST]
func UpdateMember(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.CreateOrUpdateMemberReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MemberInfo .
// @router /api/admin/member/info [GET]
func MemberInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MemberList .
// @router /api/admin/member/list [POST]
func MemberList(ctx context.Context, c *app.RequestContext) {
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

// UpdateProfile .
// @router /api/admin/member/profile-update [POST]
func UpdateProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req member.ProfileReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MemberProfile .
// @router /api/admin/member/profile [GET]
func MemberProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateMemberStatus .
// @router /api/admin/member/status [POST]
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