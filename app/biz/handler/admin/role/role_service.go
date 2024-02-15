// Code generated by hertz generator.

package role

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	role "saas/app/biz/model/admin/role"
	base "saas/app/biz/model/base"
)

// CreateRole .
// @router /api/admin/role/create [POST]
func CreateRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req role.RoleInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateRole .
// @router /api/admin/role/update [POST]
func UpdateRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req role.RoleInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteRole .
// @router /api/admin/role [POST]
func DeleteRole(ctx context.Context, c *app.RequestContext) {
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

// RoleByID .
// @router /api/admin/role [GET]
func RoleByID(ctx context.Context, c *app.RequestContext) {
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

// RoleList .
// @router /api/admin/role/list [GET]
func RoleList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.PageInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateRoleStatus .
// @router /api/admin/role/status [POST]
func UpdateRoleStatus(ctx context.Context, c *app.RequestContext) {
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