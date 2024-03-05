// Code generated by hertz generator.

package role

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"saas/app/admin/idl_gen/model/admin/role"
	"saas/app/admin/idl_gen/model/base"
	"saas/app/admin/pkg/errno"
	"saas/app/admin/pkg/utils"
	"saas/app/pkg/do"
	"saas/app/pkg/service/admin"
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

	err = admin.NewRole(ctx, c).Create(do.RoleInfo{
		Name:          req.Name,
		Value:         req.Value,
		DefaultRouter: req.DefaultRouter,
		Status:        req.Status,
		Remark:        req.Remark,
		OrderNo:       int32(req.OrderNo),
	})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
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

	err = admin.NewRole(ctx, c).Update(do.RoleInfo{
		ID:            req.ID,
		Name:          req.Name,
		Value:         req.Value,
		DefaultRouter: req.DefaultRouter,
		Status:        req.Status,
		Remark:        req.Remark,
		OrderNo:       int32(req.OrderNo),
	})
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
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

	err = admin.NewRole(ctx, c).Delete(int64(req.ID))
	if err != nil {

	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
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

	roleInfo, err := admin.NewRole(ctx, c).RoleInfoByID(int64(req.ID))
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	utils.SendResponse(c, errno.Success, roleInfo, 0, "")
	return
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

	var listReq do.RoleListReq
	listReq.Page = (req.Page)
	listReq.PageSize = (req.PageSize)
	list, total, err := admin.NewRole(ctx, c).List(&listReq)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
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

	err = admin.NewRole(ctx, c).UpdateStatus(int64(req.ID), int8(req.Status))
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}
