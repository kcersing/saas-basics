package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/wechat"
	"saas/biz/infras/do"
	"saas/config"
)

type Wechat struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (w Wechat) GetOpenid(code string) (string, error) {

	rs, err := wechat.MiniProgramApp.Auth.Session(w.ctx, code)
	if err != nil {
		return "", err
	}
	wechat.MiniProgramApp.AccessToken.Refresh(w.ctx)
	print(rs)

	return rs.OpenID, err

	//encrypt := "DRdLarcsm8c2jsdSOubaLWh/FuNgC38OszQly2n5OTchOMdTMJb6FqwfSKD3PV3B5UKfljlwpkdQxFAd3q8orqfBK/X9+lth5zXvSZ+OxICncxqbycpJFf+o695N2aCE66oC+GPebSpYld0aiUYtbPZQdwr1uqg1toCIDkZCKuuXkFFl5QeKgwg6VJVZk5w5IM2mDzAnUr8pIQllvlA/4A=="
	//sessionKey := "7o91EfeHO4PEnoeVzJxvqw=="
	////openId:="o4QEk5Kc_y8QTrENCpKoxYhS4jkg"
	////unionId:="orLIIs_Tfr0DLoG2iMwSq7RuaYRg"
	//iv := "JhojkHhkdgqHPkqD9uYZYQ=="
	//rs_e, err_e := wechat.MiniProgramApp.Encryptor.DecryptData(encrypt, sessionKey, iv)
	////services.MiniProgramApp.Base.CheckEncryptedData(c.Request.Context())
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err_e != nil {
	//	panic(err_e)
	//}
	//return rs_e, err
}

func NewWechat(ctx context.Context, c *app.RequestContext) do.Wechat {

	return &Wechat{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
