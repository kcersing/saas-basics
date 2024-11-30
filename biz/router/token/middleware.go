// Code generated by hertz generator.

package token

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

func _tokenMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletetokenMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _tokenlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _token0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatetokenMw() []app.HandlerFunc {
	// your code...
	return nil
}
