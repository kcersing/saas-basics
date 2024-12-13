package wechat

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger/drivers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"saas/config"
	"saas/pkg/consts"
	"sync"
)

var MiniProgramApp *miniProgram.MiniProgram

var onceMiniProgramApp sync.Once

func InitMiniProgramApp() {
	onceMiniProgramApp.Do(func() {
		MiniProgramApp = NewMiniMiniProgramService()
	})
}

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

func NewMiniMiniProgramService() *miniProgram.MiniProgram {
	conf := config.GlobalServerConfig.Wechat.WXMiniProgram
	var cache kernel.CacheInterface
	if config.GlobalServerConfig.Redis.Host != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs: []string{config.GlobalServerConfig.Redis.Host},
		})
	}
	wechatFilePath := consts.WechatFilePath
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:        conf.AppID,  // 小程序、公众号或者企业微信的appid
		Secret:       conf.Secret, // 商户号 appID
		ResponseType: response.TYPE_MAP,
		Token:        conf.Token,
		AESKey:       conf.AESKey,

		AppKey:  conf.AppKey,
		OfferID: conf.OfferID,
		Http:    miniProgram.Http{},
		Log: miniProgram.Log{
			Driver: &drivers.SimpleLogger{},
			Level:  "debug",
			File:   wechatFilePath + "/mini_log.log",
		},
		//"sandbox": true,
		Cache:     cache,
		HttpDebug: true,
		Debug:     false,
	})
	if err != nil {

	}
	return app
}
