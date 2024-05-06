/*
 * @Author: kcersing wt4@live.cn
 * @Date: 2024-04-21 16:12:04
 * @LastEditors: kcersing wt4@live.cn
 * @LastEditTime: 2024-04-22 23:41:52
 * @FilePath: \saas\app\admin\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/admin/pkg/minio"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/reverseproxy"
)

func init() {

	infras.InitLogger()
	infras.InitConfig()
	infras.InitDB()
	infras.InitCasbin()
	infras.InitCache()
	//infras.NewInitDatabase().InitDatabase()
	//infras.NewInitDatabase().InitDatabase2()
	minio.Init()
}
func minioReverseProxy(c context.Context, ctx *app.RequestContext) {
	proxy, _ := reverseproxy.NewSingleHostReverseProxy(config.GlobalServerConfig.Minio.Url)
	ctx.URI().SetPath(ctx.Param("name"))
	hlog.CtxInfof(c, string(ctx.Request.URI().Path()))
	proxy.ServeHTTP(c, ctx)
}

func main() {

	h := server.Default(
		server.WithStreamBody(true),
		server.WithHostPorts(fmt.Sprintf("%s:%d", config.GlobalServerConfig.Host, config.GlobalServerConfig.Port)),
	)

	h.Use(accesslog.New(
		accesslog.WithFormat("[${time}] | ${status} | ${latency} | ${method} | ${path} | ${queryParams}"),
		accesslog.WithTimeFormat(time.DateTime),
	),
	)
	// Set up /src/*name route forwarding to access minio from external network
	h.GET("/src/*name", minioReverseProxy)
	register(h)
	h.Spin()
}
