// Code generated by hertz generator.

package admin

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	admin "saas/app/admin/idl_gen/model/admin/admin"
	base "saas/app/admin/idl_gen/model/base"
)

// HealthCheck .
// @router /api/health [GET]
func HealthCheck(ctx context.Context, c *app.RequestContext) {
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

// Captcha .
// @router /api/captcha [POST]
func Captcha(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.CaptchaInfoResp
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// AdminLogin .
// @router /admin/login [POST]
func AdminLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.AdminLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// AdminChangePassword .
// @router /admin/password [POST]
func AdminChangePassword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.AdminChangePasswordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// AdminAddUser .
// @router /admin/add-user [POST]
func AdminAddUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.AddUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// AdminDeleteUser .
// @router /admin/delete-user [POST]
func AdminDeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.DeleteUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}