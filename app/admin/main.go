// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"time"
)

func main() {

	infras.InitLogger()
	infras.InitConfig()
	infras.InitDB()
	infras.InitCasbin()
	infras.InitCache()

	//tracer, trcCfg := hertztracing.NewServerTracer()
	// create a new server
	h := server.New(
		//tracer,
		//server.WithALPN(true),
		server.WithHostPorts(fmt.Sprintf("%s:%d", config.GlobalServerConfig.Host, config.GlobalServerConfig.Port)),
		//server.WithHandleMethodNotAllowed(true),
	)

	h.Use(accesslog.New(
		accesslog.WithFormat("[${time}] | ${status} | ${latency} | ${method} | ${path} | ${queryParams}"),
		accesslog.WithTimeFormat(time.DateTime),
	),
	)

	// add http2
	//h.AddProtocol(
	//	"h2",
	//	factory.NewServerFactory(
	//		cfg.WithReadTimeout(time.Minute),
	//		cfg.WithDisableKeepAlive(false),
	//	),
	//)
	// use pprof & tracer & sentinel
	//pprof.Register(h)
	//h.Use(hertztracing.ServerMiddleware(trcCfg))
	//h.Use(hertzSentinel.SentinelServerMiddleware(
	//	hertzSentinel.WithServerBlockFallback(func(c context.Context, ctx *app.RequestContext) {
	//		// abort with status 429 by default
	//		ctx.JSON(http.StatusTooManyRequests, nil)
	//		ctx.Abort()
	//	}),
	//))
	register(h)
	h.Spin()

}
