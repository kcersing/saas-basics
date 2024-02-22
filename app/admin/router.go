// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "saas/app/admin/biz/handler"
	"saas/app/admin/biz/handler/mw"
	"saas/app/admin/infras"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/__vite_ping", handler.Ping)

	// your code ...

	r.POST("/api/login", mw.GetJWTMw(infras.CasbinEnforcer()).LoginHandler)
	r.POST("/api/logout", mw.GetJWTMw(infras.CasbinEnforcer()).LogoutHandler)
	r.POST("/api/refresh_token", mw.GetJWTMw(infras.CasbinEnforcer()).RefreshHandler)
}
