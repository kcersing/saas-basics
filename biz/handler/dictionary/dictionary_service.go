// Code generated by hertz generator.

package dictionary

import (
	"context"
	admin "saas/biz/infras/service"
	"saas/pkg/errno"
	"saas/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	dictionary "saas/idl_gen/model/dictionary"
)

// CreateDictionary .
//
//	@Summary		创建字典 Summary
//	@Description	创建字典 Description
//	@Param			request	body		dictionary.DictionaryInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/dict/create [POST]
func CreateDictionary(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictionaryInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = admin.NewDictionary(ctx, c).Create(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateDictionary .
//
//	@Summary		更新字典 Summary
//	@Description	更新字典 Description
//	@Param			request	body		dictionary.DictionaryInfo	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/dict/update [POST]
func UpdateDictionary(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictionaryInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = admin.NewDictionary(ctx, c).Update(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DeleteDictionary .
//
//	@Summary		删除字典 Summary
//	@Description	删除字典 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/dict [POST]
func DeleteDictionary(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = admin.NewDictionary(ctx, c).DeleteDetail(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DictionaryList .
//
//	@Summary		字典列表 Summary
//	@Description	字典列表 Description
//	@Param			request	body		dictionary.DictListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/dict/list [GET]
func DictionaryList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// get dict list
	dictList, total, err := admin.NewDictionary(ctx, c).List(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, dictList, int64(total), "")
	return
}

// CreateDictionaryDetail .
//
//	@Summary		创建字典字段 Summary
//	@Description	创建字典字段 Description
//	@Param			request	body		dictionary.DictionaryDetail	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/dict/detail/create [POST]
func CreateDictionaryDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictionaryDetail
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = admin.NewDictionary(ctx, c).CreateDetail(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateDictionaryDetail .
//
//	@Summary		更新字典字段 Summary
//	@Description	更新字典字段 Description
//	@Param			request	body		dictionary.DictionaryDetail	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/dict/detail/update [POST]
func UpdateDictionaryDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictionaryDetail
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = admin.NewDictionary(ctx, c).UpdateDetail(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DeleteDictionaryDetail .
//
//	@Summary		删除字典字段 Summary
//	@Description	删除字典字段 Description
//	@Param			request	body		base.IDReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/dict/detail [GET]
func DeleteDictionaryDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	err = admin.NewDictionary(ctx, c).DeleteDetail(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DetailByDictionaryList .
//
//	@Summary		字典字段列表 Summary
//	@Description	字典字段列表 Description
//	@Param			request	body		dictionary.DetailListReq	true	"query params"
//	@Success		200		{object}	utils.Response
//	@router			/service/dict/detail/list [POST]
func DetailByDictionaryList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DetailListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	dictDetailList, total, err := admin.NewDictionary(ctx, c).DetailListByDict(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, dictDetailList, total, "")
	return
}
