// Code generated by hertz generator.

package entry

import (
	"context"
	"saas/biz/infras/service"
	"saas/pkg/errno"
	"saas/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	entry "saas/idl_gen/model/entry"
)

// CreateEntry .
// @Summary 创建进场记录 Summary
// @Description 创建进场记录 Description
// @Param request body entry.CreateEntry true "query params"
// @Success      200  {object}  utils.Response
// @router /service/entry/list [POST]
func CreateEntry(ctx context.Context, c *app.RequestContext) {
	var err error
	var req entry.CreateEntry
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// EntryList .
// @Summary  进场记录列表 Summary
// @Description 进场记录列表 Description
// @Param request body entry.EntryListReq true "query params"
// @Success      200  {object}  utils.Response
// @router /service/entry/list [POST]
func EntryList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req entry.EntryListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	list, total, err := service.NewEntryLogs(ctx, c).List(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}
