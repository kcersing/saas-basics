// Code generated by hertz generator.

package role

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	role "saas/app/admin/idl_gen/model/admin/role"
	base "saas/app/admin/idl_gen/model/base"
	"saas/app/admin/pkg/errno"
	"saas/app/admin/pkg/utils"
	"saas/app/pkg/do"
	"saas/app/pkg/service/admin"
)

// CreateApi .
// @router /api/admin/api/create [POST]
func CreateApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req role.ApiInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var ApiInfoReq do.ApiInfo
	err = copier.Copy(&ApiInfoReq, &req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	err = admin.NewApi(ctx, c).Create(ApiInfoReq)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateApi .
// @router /api/admin/api/update [POST]
func UpdateApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req role.ApiInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var ApiInfoReq do.ApiInfo
	err = copier.Copy(&ApiInfoReq, &req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	err = admin.NewApi(ctx, c).Update(ApiInfoReq)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DeleteApi .
// @router /api/admin/api [POST]
func DeleteApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = admin.NewApi(ctx, c).Delete(req.ID)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// ApiList .
// @router /api/admin/api/list [POST]
func ApiList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req role.ApiPageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var ApiPageReq do.ListApiReq
	err = copier.Copy(&ApiPageReq, &req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	list, total, err := admin.NewApi(ctx, c).List(ApiPageReq)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}
