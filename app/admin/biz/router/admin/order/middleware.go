// Code generated by hertz generator.

package order

import (
	"github.com/cloudwego/hertz/pkg/app"
	"saas/app/admin/biz/handler/mw"
	"saas/app/admin/infras"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.GetJWTMw(infras.CasbinEnforcer()).MiddlewareFunc(),
		mw.LogMw(),
	}
}

func _orderMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _cancelorderMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createorderMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getorderbyidMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listorderMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatestatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateorderMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _qrpayMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _unifypayMw() []app.HandlerFunc {
	// your code...
	return nil
}
