// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/jinzhu/copier"
	"saas/app/admin/pkg/errno"
	"saas/app/admin/pkg/utils"
	"saas/app/pkg/do"
	"saas/app/pkg/service/admin"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "saas/app/admin/idl_gen/model/admin/user"
	base "saas/app/admin/idl_gen/model/base"
)

// Register .
// @router /api/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserPermCode .
// @router /api/admin/user/perm [POST]
func UserPermCode(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	roleId, exist := c.Get("role_id")
	if !exist || roleId == nil {
		utils.SendResponse(c, errno.Unauthorized, nil, 0, "")
	}

	utils.SendResponse(c, errno.Success, []string{roleId.(string)}, 0, "")
	return
}

// ChangePassword .
// @router /api/admin/user/change-password [POST]
func ChangePassword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ChangePasswordReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	err = admin.NewUser(ctx, c).ChangePassword(req.UserID, req.OldPassword, req.NewPassword)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return

}

// CreateUser .
// @router /api/admin/user/create [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateOrUpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	err = admin.NewUser(ctx, c).Create(do.CreateOrUpdateUserReq{
		Username: *req.Username,
		Password: *req.Password,
		Email:    *req.Email,
		Mobile:   *req.Mobile,
		RoleID:   *req.RoleID,
		Avatar:   *req.Avatar,
		Nickname: *req.Nickname,
		Status:   *req.Status,
	})

	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateUser .
// @router /api/admin/user/update [POST]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateOrUpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	err = admin.NewUser(ctx, c).Update(do.CreateOrUpdateUserReq{
		ID:       req.ID,
		Username: *req.Username,
		Password: *req.Password,
		Email:    *req.Email,
		Mobile:   *req.Mobile,
		RoleID:   *req.RoleID,
		Avatar:   *req.Avatar,
		Nickname: *req.Nickname,
		Status:   *req.Status,
	})
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return

}

// UserInfo .
// @router /api/admin/user/info [POST]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	v, exist := c.Get("user_id")
	if !exist || v == nil {
		c.JSON(consts.StatusUnauthorized, "Unauthorized")
		return
	}
	i, err := strconv.Atoi(v.(string))
	if err != nil {
		c.JSON(consts.StatusUnauthorized, "Unauthorized,"+err.Error())
		return
	}
	userID := uint64(i)

	userInfo, err := admin.NewUser(ctx, c).UserInfo(int64(userID))
	if err != nil {
		c.JSON(consts.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(c, errno.Success, userInfo, 0, "")
	return
}

// UserList .
// @router /api/admin/user/list [POST]
func UserList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	var userListReq do.UserListReq
	err = copier.Copy(&userListReq, &req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	userList, total, err := admin.NewUser(ctx, c).List(userListReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, userList, int64(total), "")
	return
}

// DeleteUser .
// @router /api/admin/user [POST]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	err = admin.NewUser(ctx, c).DeleteUser(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateProfile .
// @router /api/admin/user/profile-update [POST]
func UpdateProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ProfileReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	var profileReq do.UpdateUserProfileReq
	err = copier.Copy(&profileReq, &req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	err = admin.NewUser(ctx, c).UpdateProfile(profileReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UserProfile .
// @router /api/admin/user/profile [POST]
func UserProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	v, exist := c.Get("user_id")
	if !exist || v == nil {
		utils.SendResponse(c, errno.ConvertErr(errno.NewErrNo(501, "Unauthorized")), nil, 0, "")
		return
	}
	i, err := strconv.Atoi(v.(string))
	if err != nil {

		utils.SendResponse(c, errno.ConvertErr(errno.NewErrNo(401, "Unauthorized,"+err.Error())), nil, 0, "")
		return
	}
	userInfo, err := admin.NewUser(ctx, c).UserInfo(int64(i))
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(errno.NewErrNo(500, err.Error())), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, userInfo, 0, "")
	return
}

// UpdateUserStatus .
// @router /api/admin/user/status [POST]
func UpdateUserStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	err = admin.NewUser(ctx, c).UpdateUserStatus(req.ID, req.Status)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}
