package infras

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"net"
	"saas/app/api/config"
	"saas/pkg/consts"
	"strconv"
)

func InitRegistry() (registry.Registry, *registry.Info) {

	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port),
	)
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}
	r := consul.NewConsulRegister(
		consulClient,
		consul.WithCheck(
			&api.AgentServiceCheck{
				Interval:                       consts.ConsulCheckInterval,
				Timeout:                        consts.ConsulCheckTimeout,
				DeregisterCriticalServiceAfter: consts.ConsulCheckDeregisterCriticalServiceAfter,
			},
		),
	)

	sf, err := snowflake.NewNode(2)
	if err != nil {
		hlog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: config.GlobalServerConfig.Name,
		Addr: utils.NewNetAddr(
			consts.TCP, net.JoinHostPort(config.GlobalServerConfig.Host,
				strconv.Itoa(config.GlobalServerConfig.Port),
			),
		),
		Weight: registry.DefaultWeight,
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
