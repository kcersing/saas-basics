package do

import (
	"saas/idl_gen/model/sms"
)

type Sms interface {
	SendVerificationCode(mobile string) (err error)
	CheckVerificationCode(mobile, verificationCode string) (err error)

	SendList(req sms.SmsSendListReq) (resp []*sms.SmsSend, total int, err error)
	Info(id int64) (resp *sms.SmsInfo, err error)
	Buy(req sms.SmsBuyReq) (resp interface{}, err error)
}
