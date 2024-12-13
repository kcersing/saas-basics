package do

type Wechat interface {
	GetOpenid(code string) (string, error)
}
