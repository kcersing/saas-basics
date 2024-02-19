// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"saas/app/admin/config"
	"saas/app/admin/initialize"
)

func main() {

	infras.InitLogger()
	infras.InitConfig()

	//tracer, trcCfg := hertztracing.NewServerTracer()
	// create a new server
	h := server.New(
		//tracer,
		//server.WithALPN(true),
		server.WithHostPorts(fmt.Sprintf(":%d", config.GlobalServerConfig.Port)),
		//server.WithHandleMethodNotAllowed(true),
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