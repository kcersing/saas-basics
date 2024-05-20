// Code generated by hertz generator.

package product

import (
	"context"
	"github.com/jinzhu/copier"
	"saas/app/admin/pkg/errno"
	"saas/app/admin/pkg/utils"
	"saas/app/pkg/do"
	"saas/app/pkg/service/admin"

	"github.com/cloudwego/hertz/pkg/app"
	product "saas/app/admin/idl_gen/model/admin/product"
	base "saas/app/admin/idl_gen/model/base"
)

// Create .
// @router /api/admin/product/create [POST]
func Create(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateOrUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	var productReq do.CreateOrUpdateProduct
	err = copier.Copy(&productReq, &req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	err = admin.NewProduct(ctx, c).Create(productReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// Update .
// @router /api/admin/product/update [POST]
func Update(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateOrUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	var productReq do.ProductInfo
	err = copier.Copy(&productReq, &req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	err = admin.NewProduct(ctx, c).Update(productReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// Delete .
// @router /api/admin/product/del [POST]
func Delete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	err = admin.NewProduct(ctx, c).Delete(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateStatus .
// @router /api/admin/product/status [POST]
func UpdateStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	err = admin.NewProduct(ctx, c).UpdateStatus(int64(req.ID), req.Status)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// InfoByID .
// @router /api/admin/product/info [POST]
func InfoByID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	info, err := admin.NewProduct(ctx, c).InfoByID(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, info, 0, "")
	return
}

// List .
// @router /api/admin/product/list [POST]
func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	var listReq do.ProductListReq
	listReq.Page = req.Page
	listReq.PageSize = req.PageSize
	list, total, err := admin.NewProduct(ctx, c).List(listReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}

// CreateProperty .
// @router /api/admin/property/create [POST]
func CreateProperty(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateOrUpdatePropertyReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	var propertyReq do.PropertyInfo
	err = copier.Copy(&propertyReq, &req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	err = admin.NewProduct(ctx, c).CreateProperty(propertyReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// UpdateProperty .
// @router /api/admin/property/update [POST]
func UpdateProperty(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.CreateOrUpdatePropertyReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	var propertyReq do.PropertyInfo
	err = copier.Copy(&propertyReq, &req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	err = admin.NewProduct(ctx, c).UpdateProperty(propertyReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// DeleteProperty .
// @router /api/admin/property/del [POST]
func DeleteProperty(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	err = admin.NewProduct(ctx, c).DeleteProperty(req.ID)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, nil, 0, "")
	return
}

// ListProperty .
// @router /api/admin/property/list [POST]
func ListProperty(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.PropertyListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	var listReq do.ProductListReq
	err = copier.Copy(&listReq, &req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	list, total, err := admin.NewProduct(ctx, c).PropertyList(listReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	utils.SendResponse(c, errno.Success, list, int64(total), "")
	return
}
