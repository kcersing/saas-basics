package infras

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"saas/pkg/consts"
	"saas/pkg/utils"
)

func InitFlag() (string, int) {
	IP := flag.String(consts.IPFlagName, consts.IPFlagValue, consts.IPFlagUsage)
	Port := flag.Int(consts.PortFlagName, 0, consts.PortFlagUsage)
	flag.Parse()
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	klog.Info("ip: ", *IP)
	klog.Info("port: ", *Port)
	return *IP, *Port

}
