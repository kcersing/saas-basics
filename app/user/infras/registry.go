package infras

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	"net"
	"saas/app/api/config"
	"saas/pkg/consts"
	"strconv"
)

func InitRegistry(Port int) (registry.Registry, *registry.Info) {

	r, err := consul.NewConsulRegister(
		net.JoinHostPort(
			config.GlobalConsulConfig.Host,
			strconv.Itoa(config.GlobalConsulConfig.Port),
		),
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       consts.ConsulCheckInterval,
			Timeout:                        consts.ConsulCheckTimeout,
			DeregisterCriticalServiceAfter: consts.ConsulCheckDeregisterCriticalServiceAfter,
		}),
	)
	if err != nil {
		klog.Fatalf("new consul register failed: %s", err.Error())
	}

	sf, err := snowflake.NewNode(2)
	if err != nil {
		klog.Fatalf("generate servive name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: config.GlobalServerConfig.Name,
		Addr:        utils.NewNetAddr(consts.TCP, net.JoinHostPort(config.GlobalServerConfig.Host, strconv.Itoa(Port))),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
