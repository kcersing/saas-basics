// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/logger/accesslog"
	prometheus "github.com/hertz-contrib/monitor-prometheus"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	"saas/biz/dal"
	"saas/config"
	_ "saas/docs"
	"time"
)

func minioReverseProxy(c context.Context, ctx *app.RequestContext) {
	proxy, _ := reverseproxy.NewSingleHostReverseProxy(config.GlobalServerConfig.Minio.Url)
	ctx.URI().SetPath(ctx.Param("name"))
	hlog.CtxInfof(c, string(ctx.Request.URI().Path()))
	proxy.ServeHTTP(c, ctx)
}
func main() {
	dal.Init()

	h := server.Default(
		server.WithStreamBody(true),
		server.WithHostPorts(fmt.Sprintf("%s:%d", config.GlobalServerConfig.Host, config.GlobalServerConfig.Port)),
		server.WithTracer(
			prometheus.NewServerTracer(":9091", "/hertz",
				prometheus.WithEnableGoCollector(true), // enable go runtime metric collector
			),
		),
	)

	//h.SetCustomSignalWaiter(func(err chan error) error {
	//	return nil
	//})
	//pprof.Register(h, "debug/pprof")

	h.Use(accesslog.New(
		accesslog.WithFormat("[${time}] | ${status} | ${latency} | ${method} | ${path} | ${queryParams}"),
		accesslog.WithTimeFormat(time.DateTime),
	))

	h.NoHijackConnPool = true
	h.GET("/src/*name", minioReverseProxy)
	h.Static("/export", "./tmp/")
	url := swagger.URL(config.GlobalServerConfig.Swagger.Url) // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	register(h)
	h.Spin()
}
