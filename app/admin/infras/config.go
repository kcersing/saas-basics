package infras

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
	"saas/app/admin/config"
	"saas/app/admin/pkg/consts"
	"saas/pkg/utils"
)

func InitConfig() {

	v := viper.New()
	v.SetConfigFile(consts.ApiConfigPath)

	if err := v.ReadInConfig(); err != nil {
		hlog.Fatalf("read viper config failed:%s", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalServerConfig); err != nil {
		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	hlog.Info("config Info: %v", config.GlobalServerConfig)

	//cfg := api.DefaultConfig()
	//cfg.Address = net.JoinHostPort(
	//	config.GlobalConsulConfig.Host,
	//	strconv.Itoa(config.GlobalConsulConfig.Port),
	//)
	//consulClient, err := api.NewClient(cfg)
	//if err != nil {
	//	hlog.Fatalf("new consul client failed:%s", err.Error())
	//}
	//content, _, err := consulClient.KV().Get(config.GlobalConsulConfig.Key, nil)
	//if err != nil {
	//	hlog.Fatalf("consul kv failed :%s", err.Error())
	//}
	//err = sonic.Unmarshal(content.Value, &config.GlobalServerConfig)
	//if err != nil {
	//	hlog.Fatalf("sonic unmarshal config failed: %s", err.Error())
	//}
	if config.GlobalServerConfig.Host == "" {
		address, err := utils.GetLocalIPv4Address()
		if err != nil {
			hlog.Fatalf("get localIpv4Addr failed:%s", err.Error())
		} else {
			config.GlobalServerConfig.Host = address
		}
	}
}
