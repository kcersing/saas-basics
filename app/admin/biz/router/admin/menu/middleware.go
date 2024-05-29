// Code generated by hertz generator.

package menu

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

func _menuMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletemenuMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menulistsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _paramMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletemenuparamMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menuparamlistbymenuidMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menubyroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _param0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createmenuparamMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatemenuparamMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menu0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createmenuMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatemenuMw() []app.HandlerFunc {
	// your code...
	return nil
}
