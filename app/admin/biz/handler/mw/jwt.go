package mw

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"github.com/spf13/cast"
	"saas/app/admin/config"
	"saas/app/pkg/do"
	"saas/app/pkg/service/admin"
	"strconv"
	"time"
)

type jwtLogin struct {
	Username  string `form:"username,required" json:"username,required"`   //lint:ignore SA5008 ignoreCheck
	Password  string `form:"password,required" json:"password,required"`   //lint:ignore SA5008 ignoreCheck
	Captcha   string `form:"captcha,required" json:"captcha,required"`     //lint:ignore SA5008 ignoreCheck
	CaptchaID string `form:"captchaId,required" json:"captchaId,required"` //lint:ignore SA5008 ignoreCheck
}

// jwt identityKey
var (
	identityKey   = "jwt-id"
	jwtMiddleware = new(jwt.HertzJWTMiddleware)
)

func GetJWTMw(e *casbin.Enforcer) *jwt.HertzJWTMiddleware {
	jwtMiddleware, err := newJWT(e)
	if err != nil {
		hlog.Fatal(err, "JWT Init Error")
	}
	return jwtMiddleware
}

func newJWT(enforcer *casbin.Enforcer) (jwtMiddleware *jwt.HertzJWTMiddleware, err error) {
	jwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "saas",
		Key:         []byte(config.GlobalServerConfig.Auth.AccessSecret),
		Timeout:     time.Duration(config.GlobalServerConfig.Auth.AccessExpire) * time.Second,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			payloadMap, ok := claims[identityKey].(map[string]interface{})
			if !ok {
				hlog.Error("get payloadMap error:", " claims data:", claims[identityKey])
				return nil
			}
			hlog.Info("IdentityHandler")
			hlog.Info(payloadMap)
			c.Set("role_id", payloadMap["role_id"])
			c.Set("user_id", payloadMap["user_id"])
			return payloadMap
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			res := new(do.LoginResp)

			var loginVal jwtLogin
			if err := c.BindAndValidate(&loginVal); err != nil {
				return "", err
			}
			// 验证码

			username := loginVal.Username
			password := loginVal.Password
			res, err = admin.NewLogin(ctx, c).Login(username, password)

			if err != nil {
				hlog.Error(err, "jwtLogin error")
				return nil, err
			}
			//jwt
			var tokenInfo do.TokenInfo
			tokenInfo.UserID = res.UserID
			tokenInfo.UserName = res.Username
			tokenInfo.ExpiredAt = time.Now().Add(time.Duration(config.GlobalServerConfig.Auth.AccessExpire) * time.Second).Format("2006-01-02 15:04:05")

			err = admin.NewToken(ctx, c).Create(&tokenInfo)
			if err != nil {
				hlog.Error(err, "jwtLogin error, store token error")
				return nil, err
			}

			payLoadMap := make(map[string]interface{})
			payLoadMap["role_id"] = strconv.Itoa(int(res.RoleID))
			payLoadMap["user_id"] = strconv.Itoa(int(res.UserID))
			return payLoadMap, nil
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			//obj := string(c.URI().Path())
			//act := string(c.Method())
			payloadMap, ok := data.(map[string]interface{})
			if !ok {
				hlog.Error("get payloadMap error:", " claims data:", data)
				return false
			}
			roleId := payloadMap["role_id"].(string)
			userId := payloadMap["user_id"].(string)

			hlog.Info("Authorizator")
			hlog.Info(payloadMap)

			// check token is valid
			userIdInt, err := strconv.Atoi(userId)
			if err != nil {
				hlog.Error("get payloadMap error:", err)
				return false
			}
			existToken := admin.NewToken(ctx, c).IsExistByUserID(uint64(userIdInt))
			if !existToken {
				return false
			}
			// check the role status
			roleInfo, err := admin.NewRole(ctx, c).RoleInfoByID(cast.ToUint64(roleId))
			// if the role is not exist or the role is not active, return false
			if err != nil {
				hlog.Error(err, "role is not exist")
				return false
			}
			if roleInfo.Status != 1 {
				hlog.Error("role cache is not a valid *ent.Role or the role is not active")
				return false
			}
			//
			//sub := roleId
			//
			//pass, err := enforcer.Enforce(sub, obj, act)
			//if err != nil {
			//	hlog.Error("casbin err,  role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass, " err: ", err.Error())
			//	return false
			//}
			//if !pass {
			//	hlog.Info("casbin forbid role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass)
			//}
			//hlog.Info("casbin allow role id: ", roleId, " path: ", obj, " method: ", act, " pass: ", pass)
			//return pass

			return true
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
	})

	return

}
