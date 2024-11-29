// Code generated by hertz generator.

package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"saas/biz/infras/service"
	auth "saas/idl_gen/model/auth"
	base "saas/idl_gen/model/base"
	"saas/pkg/errno"
	"saas/pkg/utils"
	"strconv"
)

// CreateRole .
// @Summary 创建角色 Summary
// @Description 创建角色 Description
// @Param request body auth.RoleInfo true "query params"
// @Success      200  {object}  utils.Response
// @router /service/role/create [POST]
func CreateRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RoleInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewRole(ctx, c).Create(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateRole .
// @Summary 更新角色 Summary
// @Description 更新角色 Description
// @Param request body auth.RoleInfo true "query params"
// @Success      200  {object}  utils.Response
// @router /service/role/update [POST]
func UpdateRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RoleInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewRole(ctx, c).Update(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DeleteRole .
// @Summary  删除角色 Summary
// @Description 删除角色 Description
// @Param request body base.IDReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/role/del [POST]
func DeleteRole(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewRole(ctx, c).Delete(int64(req.ID))
	if err != nil {

	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// RoleByID .
// @Summary  获取角色信息 Summary
// @Description 获取角色信息 Description
// @Param request body base.IDReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/role [GET]
func RoleByID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	roleInfo, err := service.NewRole(ctx, c).RoleInfoByID(int64(req.ID))
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, roleInfo, 0, "")
	return
}

// CreateMenuAuth .
// @Summary  创建角色菜单权限 Summary
// @Description 创建角色菜单权限 Description
// @Param request body auth.MenuAuthInfoReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/auth/menu/create [POST]
func CreateMenuAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MenuAuthInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewAuth(ctx, c).UpdateMenuAuth(req.RoleID, req.MenuIds)
	if err != nil {

	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateMenuAuth .
// @Summary  修改角色菜单权限 Summary
// @Description 修改角色菜单权限 Description
// @Param request body auth.MenuAuthInfoReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/auth/menu/update [POST]
func UpdateMenuAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MenuAuthInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	err = service.NewAuth(ctx, c).UpdateMenuAuth(req.RoleID, req.MenuIds)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// RoleList .
// @Summary  获取角色列表 Summary
// @Description 获取角色列表 Description
// @Param request body base.PageInfoReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/role/list [GET]
func RoleList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.PageInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewRole(ctx, c).List(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// UpdateRoleStatus .
// @Summary  修改角色状态 Summary
// @Description 修改角色状态 Description
// @Param request body base.StatusCodeReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/role/status [POST]
func UpdateRoleStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewRole(ctx, c).UpdateStatus(req.ID, req.Status)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// CreateAuth .
// @Summary  创建角色接口权限 Summary
// @Description 创建角色接口权限 Description
// @Param request body auth.CreateOrUpdateApiAuthReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/auth/api/create [POST]
func CreateAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.CreateOrUpdateApiAuthReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewAuth(ctx, c).UpdateApiAuth(strconv.FormatInt(req.RoleID, 10), req.Apis)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateAuth .
// @Summary  修改角色接口权限 Summary
// @Description 修改角色接口权限 Description
// @Param request body auth.CreateOrUpdateApiAuthReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/auth/api/update [POST]
func UpdateAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.CreateOrUpdateApiAuthReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	err = service.NewAuth(ctx, c).UpdateApiAuth(strconv.FormatInt(req.RoleID, 10), req.Apis)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// ApiAuth .
// @Summary  获取角色接口权限 Summary
// @Description 获取角色接口权限 Description
// @Param request body base.IDReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/auth/api/role [POST]
func ApiAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	roleId := strconv.FormatInt(req.ID, 10)
	if roleId == "0" {
		c.String(consts.StatusBadRequest, "roleId 不应为0")
		return
	}
	policies, err := service.NewAuth(ctx, c).ApiAuth(roleId)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	total := uint64(len(policies))
	utils.SendResponse(c, errno.Success, policies, int64(total), "")
	return
}

// GetLogsList .
// @Summary  获取操作日志列表 Summary
// @Description 获取操作日志列表 Description
// @Param request body auth.LogsListReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/logs/list [POST]
func GetLogsList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LogsListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	logsList, total, err := service.NewLogs(ctx, c).List(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, logsList, int64(total), "")
	return
}

// DeleteLogs .
// @Summary  删除操作日志 Summary
// @Description 删除操作日志 Description
// @Param request body base.Ids true "query params"
// @Success      200  {object}  utils.Response
// @router /service/logs/deleteAll [POST]
func DeleteLogs(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Ids
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewLogs(ctx, c).DeleteAll(req.Ids)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}
