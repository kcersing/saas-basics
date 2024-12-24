package do

type Sms interface {
	SendVerificationCode(mobile string) (err error)
	CheckVerificationCode(mobile, verificationCode string) (err error)
}
