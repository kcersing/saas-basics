// Code generated by hertz generator.

package dictionary

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "saas/idl_gen/model/base"
	dictionary "saas/idl_gen/model/dictionary"
)

// CreateDictionary .
// @router /service/dict/create [POST]
func CreateDictionary(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictionaryInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateDictionary .
// @router /service/dict/update [POST]
func UpdateDictionary(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictionaryInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteDictionary .
// @router /service/dict [POST]
func DeleteDictionary(ctx context.Context, c *app.RequestContext) {
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

// DictionaryList .
// @router /service/dict/list [GET]
func DictionaryList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CreateDictionaryDetail .
// @router /service/dict/detail/create [POST]
func CreateDictionaryDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictionaryDetail
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateDictionaryDetail .
// @router /service/dict/detail/update [POST]
func UpdateDictionaryDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DictionaryDetail
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteDictionaryDetail .
// @router /service/dict/detail [GET]
func DeleteDictionaryDetail(ctx context.Context, c *app.RequestContext) {
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

// DetailByDictionaryList .
// @router /service/dict/detail/list [POST]
func DetailByDictionaryList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req dictionary.DetailListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
