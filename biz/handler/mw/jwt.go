package mw

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"github.com/spf13/cast"
	"saas/biz/infras/service"
	"saas/config"
	"saas/idl_gen/model/token"
	"saas/idl_gen/model/user"
	"saas/pkg/errno"
	"saas/pkg/utils"
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
		// PayloadFunc is used to define a custom token payload.
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// IdentityHandler is used to define a custom identity handler.
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			payloadMap, ok := claims[identityKey].(map[string]interface{})
			if !ok {
				hlog.Error("get payloadMap error:", " claims data:", claims[identityKey])
				return nil
			}
			hlog.Info("IdentityHandler")
			hlog.Info(payloadMap)
			c.Set("roleId", payloadMap["roleId"])
			c.Set("userId", payloadMap["userId"])
			return payloadMap
		},
		// Authenticator is used to validate the login data.
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			res := new(user.LoginResp)

			var loginVal jwtLogin
			if err := c.BindAndValidate(&loginVal); err != nil {
				return "", err
			}
			// 验证码

			username := loginVal.Username
			password := loginVal.Password
			res, err = service.NewLogin(ctx, c).Login(username, password)

			if err != nil {
				hlog.Error(err, "jwtLogin error")
				return nil, err
			}
			//jwt
			var tokenInfo token.TokenInfo
			tokenInfo.UserID = res.UserId
			tokenInfo.Username = res.Username
			tokenInfo.ExpiredAt = time.Now().Add(time.Duration(config.GlobalServerConfig.Auth.AccessExpire) * time.Second).Format(time.DateTime)

			err = service.NewToken(ctx, c).Create(&tokenInfo)
			if err != nil {
				hlog.Error(err, "jwtLogin error, store token error")
				return nil, err
			}

			payLoadMap := make(map[string]interface{})
			payLoadMap["roleId"] = strconv.Itoa(int(res.RoleId))
			payLoadMap["userId"] = strconv.Itoa(int(res.UserId))
			return payLoadMap, nil
		},
		// Authorizator is used to validate the authentication of the current request.
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			//obj := string(c.URI().Path())
			//act := string(c.Method())
			payloadMap, ok := data.(map[string]interface{})
			if !ok {
				hlog.Error("get payloadMap error:", " claims data:", data)
				return false
			}
			roleId := payloadMap["roleId"].(string)
			userId := payloadMap["userId"].(string)

			hlog.Info("Authorizator")
			hlog.Info(payloadMap)

			// check token is valid
			userIdInt, err := strconv.Atoi(userId)
			if err != nil {
				hlog.Error("get payloadMap error:", err)
				return false
			}
			existToken := service.NewToken(ctx, c).IsExistByUserID(int64(userIdInt))
			if !existToken {
				return false
			}
			// check the role status
			roleInfo, err := service.NewRole(ctx, c).RoleInfoByID(cast.ToInt64(roleId))
			// if the role is not exist or the role is not active, return false
			if err != nil {
				hlog.Error(err, "role is not exist")
				return false
			}
			if *roleInfo.Status != 1 {
				hlog.Error("role cache is not a valid *ent.Role or the role is not active")
				return false
			}

			//sub := roleId
			//check the permission
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
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			utils.SendResponse(c, errno.Success,
				map[string]interface{}{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				}, 0, "")
		},
	})

	return

}
