package do

type Captcha interface {
	GetCaptcha() (id, b64s string, err error)
}
