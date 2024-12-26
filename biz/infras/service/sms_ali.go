package service

// This file is auto-generated, don't edit it. Thanks.

import (
	"context"
	"errors"
	"fmt"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cloudwego/hertz/pkg/app"
	"math/rand"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/sms"
	"saas/biz/infras/do"
	"saas/config"
	"time"
)

func (s Sms) SendVerificationCode(mobile string) (err error) {

	_, exist := s.cache.Get("VerificationCode_" + mobile)
	if exist {
		err = errors.New("请勿重复发送验证码")
		return
	}
	verifyCode := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	bizId, err := sms.SendAliyunSms(
		dysmsapi20170525.SendSmsRequest{
			SignName:      tea.String("阿里云短信测试"),
			TemplateCode:  tea.String("SMS_154950909"),
			PhoneNumbers:  tea.String(mobile),
			TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%s\"}", verifyCode)),
		},
	)

	if err != nil {
		return err
	}
	s.cache.SetWithTTL("VerificationCode_"+mobile, verifyCode, 1, 5*time.Minute)

	go func() {
		if bizId != "" {
			s.db.VenueSmsLog.Create().
				SetMobile(mobile).
				SetCode(verifyCode).
				SetBizID(bizId).
				SetTemplate("SMS_154950909").
				Save(s.ctx)
		}
	}()

	return nil
}

func (s Sms) CheckVerificationCode(mobile, verificationCode string) (err error) {
	code, exist := s.cache.Get("VerificationCode_" + mobile)
	if !exist {
		err = errors.New("验证码已失效")
		return
	}
	if vc, ok := code.(string); ok {
		if vc != verificationCode {
			err = errors.New("验证码输入错误")
			return
		}
	} else {
		err = errors.New("验证码解析错误")
		return
	}

	return nil
}

func NewAliyunSms(ctx context.Context, c *app.RequestContext) do.Sms {
	return &Sms{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
