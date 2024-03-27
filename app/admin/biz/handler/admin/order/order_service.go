// Code generated by hertz generator.

package order

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	order "saas/app/admin/idl_gen/model/admin/order"
	base "saas/app/admin/idl_gen/model/base"
)

// CreateOrder .
// @router /api/admin/order/create [POST]
func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.CreateOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateUser .
// @router /api/admin/order/update [POST]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.UpdateOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CancelOrder .
// @router /api/admin/order/cancel [POST]
func CancelOrder(ctx context.Context, c *app.RequestContext) {
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

// ListOrder .
// @router /api/admin/order/list [POST]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.ListOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetOrderById .
// @router /api/admin/order/info [GET]
func GetOrderById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.GetOrderByIdReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
