// Code generated by hertz generator.

package menu

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"saas/biz/infras/service"
	base "saas/idl_gen/model/base"
	menu "saas/idl_gen/model/menu"
	"saas/pkg/errno"
	"saas/pkg/utils"
	"strconv"
)

// MenuAuth .
//
//	@Summary		获取角色菜单权限 Summary
//	@Description	获取角色菜单权限 Description
//	@Param			request	body		base.IDReq	true	"id是取得roleId"
//	@Success		200		{object}	utils.Response
//	@router			/service/auth/menu/role [POST]
func MenuAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	menuIDs, err := service.NewAuth(ctx, c).MenuAuth(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	m := map[string]interface{}{
		"roleID":  req.ID,
		"menuIDs": menuIDs,
	}
	utils.SendResponse(c, errno.Success, m, 0, "")
	return
}

// ApiList .
//
//	@Summary		获取api列表 Summary
//	@Description	获取api列表 Description
//	@Param			request	body		menu.ApiPageReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/api/list [POST]
func ApiList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.ApiPageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewApi(ctx, c).List(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// ApiTree .
//
//	@Summary		获取api树 Summary
//	@Description	获取api树 Description
//	@Param			request	body		menu.ApiPageReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/api/tree [POST]
func ApiTree(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.ApiPageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewApi(ctx, c).ApiTree(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// MenuLists .
//
//	@Summary		获取菜单列表 Summary
//	@Description	获取菜单列表 Description
//	@Param			request	body		menu.ApiPageReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/menu/list [POST]
func MenuLists(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.PageInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	menuTree, total, err := service.NewMenu(ctx, c).List(&req)
	hlog.Info(menuTree)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, menuTree, int64(total), "")
	return
}

// MenuTree .
//
//	@Summary		获取菜单树 Summary
//	@Description	获取菜单树 Description
//	@Param			request	body		menu.ApiPageReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/menu/tree [POST]
func MenuTree(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.PageInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	menuTree, err := service.NewMenu(ctx, c).MenuTree(&req)
	hlog.Info(menuTree)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, menuTree, 0, "")
	return
}

// MenuRole .
//
//	@Summary		获取角色菜单权限 Summary
//	@Description	获取角色菜单权限 （和 /service/auth/menu/role 重复 但是这个是通过token获取 无传值） Description
//	@Success		200	{object}	utils.Response
//	@router			/service/menu/role [POST]
func MenuRole(ctx context.Context, c *app.RequestContext) {

	roleIdInterface, exist := c.Get("roleId")
	if !exist || roleIdInterface == nil {
		utils.SendResponse(c, errno.Unauthorized, nil, 0, "")
		return
	}
	roleId, err := strconv.Atoi(roleIdInterface.(string))
	if err != nil {
		utils.SendResponse(c, errno.Unauthorized, nil, 0, "")
		return
	}
	menuTree, err := service.NewMenu(ctx, c).MenuRole(int64(roleId))
	if err != nil {
		utils.SendResponse(c, errors.New(err.Error()), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, menuTree, 0, "")
	return
}

// CreateApi .
//
//	@Summary		创建IPA Summary
//	@Description	创建IPA Description
//	@Param			request	body		menu.ApiInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/api/create [POST]
func CreateApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.ApiInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	err = service.NewApi(ctx, c).Create(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateApi .
//
//	@Summary		更新IPA Summary
//	@Description	更新IPA Description
//	@Param			request	body		menu.ApiInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/api/update [POST]
func UpdateApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.ApiInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewApi(ctx, c).Update(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DeleteApi .
//
//	@router	/service/api [POST]
func DeleteApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewApi(ctx, c).Delete(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// CreateMenu .
//
//	@Summary		创建菜单 Summary
//	@Description	创建菜单 Description
//	@Param			request	body		menu.CreateOrUpdateMenuReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/menu/create [POST]
func CreateMenu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.CreateOrUpdateMenuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewMenu(ctx, c).Create(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateMenu .
//
//	@Summary		更新菜单 Summary
//	@Description	更新菜单 Description
//	@Param			request	body		menu.CreateOrUpdateMenuReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/menu/update [POST]
func UpdateMenu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.CreateOrUpdateMenuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewMenu(ctx, c).Update(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DeleteMenu .
//
//	@Summary		删除菜单 Summary
//	@Description	删除菜单 Description
//	@Param			request	body		menu.CreateOrUpdateMenuReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/menu [POST]
func DeleteMenu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewMenu(ctx, c).Delete(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// MenuInfo .
//
//	@router	/service/menu/info [POST]
func MenuInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	info, err := service.NewMenu(ctx, c).MenuInfo(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}
