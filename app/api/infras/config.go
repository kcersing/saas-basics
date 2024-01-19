package infras

import (
	"github.com/spf13/viper"
	"saas_basics/shared/consts"
)

func INitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.ApiConfigPath)
}
