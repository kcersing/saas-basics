package infras

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
	"saas/app/user/config"
	"saas/pkg/consts"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.UserConfigPath)
	if err := v.ReadInConfig(); err != nil {
		hlog.Fatalf("read viper config failed: %d", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalConsulConfig); err != nil {
		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	hlog.Infof("Config Info: %v", config.GlobalConsulConfig)

	v1 := viper.New()
	v1.SetConfigFile(consts.UserConfigPath1)

	if err := v1.ReadInConfig(); err != nil {
		hlog.Fatalf("read viper config failed:%s", err.Error())
	}
	if err := v1.Unmarshal(&config.GlobalServerConfig); err != nil {
		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	hlog.Info("config Info: %v", config.GlobalServerConfig)
}
