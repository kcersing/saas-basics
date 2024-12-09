// Code generated by hertz generator. DO NOT EDIT.

package order

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	order "saas/biz/handler/order"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_service := root.Group("/service", _serviceMw()...)
		{
			_order := _service.Group("/order", _orderMw()...)
			_order.GET("/info", append(_getorderbyidMw(), order.GetOrderById)...)
			_order.POST("/list", append(_listorderMw(), order.ListOrder)...)
			_order.POST("/status", append(_updatestatusMw(), order.UpdateStatus)...)
		}
	}
}
