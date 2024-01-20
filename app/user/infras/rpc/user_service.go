package rpc

import (
	"github.com/cloudwego/kitex/client"
	_ "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	_ "github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	_ "github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	_ "github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	_ "github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"saas/app/user/config"
	"saas/kitex_gen/cwg/user/userservice"
)

func initUser() {
	// init resolver
	r, err := consul.NewConsulResolver("")
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}
	// init OpenTelemetry
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(""),
		provider.WithExportEndpoint(""),
		provider.WithInsecure(),
	)
	//create a new client
	c, err := userservice.NewClient(
		"",
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "",
			Method:      "",
			Tags:        nil,
		}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	config.GlobalUserClient = c
}
