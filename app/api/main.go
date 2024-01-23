// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	cfg "github.com/hertz-contrib/http2/config"
	"github.com/hertz-contrib/http2/factory"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	hertzSentinel "github.com/hertz-contrib/opensergo/sentinel/adapter"
	"github.com/hertz-contrib/pprof"
	"net/http"
	"saas/app/api/config"
	"saas/app/api/infras"
	"saas/app/api/infras/rpc"
	"time"
)

func main() {

	infras.InitLogger()
	infras.InitConfig()
	r, info := infras.InitRegistry()
	rpc.Init()
	tracer, trcCfg := hertztracing.NewServerTracer()
	tlsCfg := infras.InitTLS()

	rpc.Init()
	// create a new server
	h := server.New(
		tracer,
		server.WithALPN(true),
		server.WithTLS(tlsCfg),
		server.WithHostPorts(fmt.Sprintf(":%d", config.GlobalServerConfig.Port)),
		server.WithRegistry(r, info),
		server.WithHandleMethodNotAllowed(true),
	)
	// add http2
	h.AddProtocol(
		"h2",
		factory.NewServerFactory(
			cfg.WithReadTimeout(time.Minute),
			cfg.WithDisableKeepAlive(false),
		),
	)

	tlsCfg.NextProtos = append(tlsCfg.NextProtos, "h2")
	// use pprof & tracer & sentinel
	pprof.Register(h)
	h.Use(hertztracing.ServerMiddleware(trcCfg))
	h.Use(hertzSentinel.SentinelServerMiddleware(
		hertzSentinel.WithServerBlockFallback(func(c context.Context, ctx *app.RequestContext) {
			// abort with status 429 by default
			ctx.JSON(http.StatusTooManyRequests, nil)
			ctx.Abort()
		}),
	))

	register(h)
	h.Spin()
}
