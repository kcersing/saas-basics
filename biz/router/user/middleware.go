// Code generated by hertz generator.

package user

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

func _deleteuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}
