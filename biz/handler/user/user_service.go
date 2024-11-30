// Code generated by hertz generator.

package user

import (
	"context"
	"saas/biz/infras/service"
	"saas/pkg/errno"
	"saas/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	user "saas/idl_gen/model/user"
)

// CreateUser .
// @Summary  创建员工 Summary
// @Description 创建员工 Description
// @Param request body user.CreateOrUpdateUserReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/user/create [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateOrUpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewUser(ctx, c).Create(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateUser .
// @Summary  更新员工 Summary
// @Description 更新员工 Description
// @Param request body user.CreateOrUpdateUserReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/user/update [POST]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateOrUpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewUser(ctx, c).Update(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UserInfo .
// @Summary  获取员工信息 Summary
// @Description 获取员工信息 Description
// @Summary  创建员工 Summary
// @Description 创建员工 Description
// @Param request body base.IDReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/user/info [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
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

// UserList .
// @Summary  获取员工列表 Summary
// @Description 获取员工列表 Description
// @Param request body user.UserListReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/user/list [POST]
func UserList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewUser(ctx, c).List(req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// DeleteUser .
// @Summary  删除员工 Summary
// @Description 删除员工 Description
// @Param request body base.IDReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/user [POST]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = service.NewUser(ctx, c).DeleteUser(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}
