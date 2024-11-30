// Code generated by hertz generator.

package auth

import (
	"github.com/cloudwego/hertz/pkg/app"
	"saas/biz/dal/casbin"
	"saas/biz/handler/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _serviceMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.GetJWTMw(casbin.CasbinEnforcer()).MiddlewareFunc(),
		mw.LogMw(),
	}
}

func _roleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _rolebyidMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _rolelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updaterolestatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _authMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createauthMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiauthMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateauthMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _menuMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createmenuauthMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatemenuauthMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _logsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletelogsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getlogslistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _role0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteroleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateroleMw() []app.HandlerFunc {
	// your code...
	return nil
}
