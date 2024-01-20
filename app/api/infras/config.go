package infras

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
	"saas_basics/shared/consts"
)

func INitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.ApiConfigPath)

	if err := v.ReadInConfig; err != nil {
		hlog.Fatalf("read viper config failed:%s", err.Error())
	}
}
