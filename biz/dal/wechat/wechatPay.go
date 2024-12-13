package wechat

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"os"
	"saas/config"
	"saas/pkg/consts"
	"sync"
)

var PaymentWechatApp *payment.Payment
var oncePaymentWechat sync.Once

func InitWXPaymentApp() {
	oncePaymentWechat.Do(func() {
		PaymentWechatApp = NewWXPaymentApp()
	})
}

func NewWXPaymentApp() *payment.Payment {

	conf := config.GlobalServerConfig.Payment.Wechatpay

	var cache kernel.CacheInterface
	if config.GlobalServerConfig.Redis.Host != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs: []string{config.GlobalServerConfig.Redis.Host},
		})
	}

	wechatFilePath := consts.WechatFilePath
	if err := os.MkdirAll(wechatFilePath, 0o777); err != nil {
		panic(err)
	}

	Payment, err := payment.NewPayment(&payment.UserConfig{
		AppID:        conf.Appid,           // 小程序、公众号或者企业微信的appid
		MchID:        conf.MchId,           // 商户号 appID
		MchApiV3Key:  conf.ApiV3Key,        // 微信V3接口调用必填
		Key:          conf.ApiKey,          // 微信V2接口调用必填
		CertPath:     conf.CertFileContent, // 商户后台支付的Cert证书路径
		KeyPath:      conf.KeyFileContent,  // 商户后台支付的Key证书路径
		SerialNo:     conf.SerialNo,        // 商户支付证书序列号
		NotifyURL:    conf.NotifyUrl,
		ResponseType: response.TYPE_MAP,

		CertificateKeyPath: conf.CertificateKeyPath,
		WechatPaySerial:    conf.WechatPaySerial,
		RSAPublicKeyPath:   conf.RSAPublicKeyPath,
		SubMchID:           conf.SubMchID,
		SubAppID:           conf.SubAppID,

		Cache: cache,
		Log: payment.Log{
			Level: "debug",
			File:  wechatFilePath + "/pay_log.log",
		},
		Http: payment.Http{
			Timeout: 30.0,
			//BaseURI: "http://127.0.0.1:8888",
			BaseURI: "https://api.mch.weixin.qq.com",
		},

		HttpDebug: false,
		Debug:     false,
		//Debug:     true,
	})

	if err != nil || Payment == nil {
		hlog.Error(err)
		return nil
	}

	return Payment
}
