package mw

//
//import (
//	"context"
//	"github.com/casbin/casbin/v2"
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/common/hlog"
//	"github.com/hertz-contrib/jwt"
//	"saas/biz/infras/service"
//	"saas/config"
//	"saas/idl_gen/model/member"
//	"saas/idl_gen/model/token"
//	"saas/idl_gen/model/wx"
//	"saas/pkg/errno"
//	"saas/pkg/utils"
//	"strconv"
//	"time"
//)
//
//type jwtUsLogin struct {
//	Mobile     string `form:"mobile,required" json:"mobile,required"`
//	VerifyCode string `form:"verifyCode,required" json:"verifyCode,required"`
//	VenueId    string `form:"venueId,required" json:"venueId,required"`
//}
//
//// jwt identityUsKey
//var (
//	identityUsKey   = "jwt-id"
//	jwtUsMiddleware = new(jwt.HertzJWTMiddleware)
//)
//
//func GetUsJWTMw(e *casbin.Enforcer) *jwt.HertzJWTMiddleware {
//	jwtMiddleware, err := newUsJWT(e)
//	if err != nil {
//		hlog.Fatal(err, "JWT Init Error")
//	}
//	return jwtMiddleware
//}
//
//func newUsJWT(enforcer *casbin.Enforcer) (jwtMiddleware *jwt.HertzJWTMiddleware, err error) {
//	jwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
//		Realm:       "wxUs",
//		Key:         []byte(config.GlobalServerConfig.Auth.AccessSecret),
//		Timeout:     time.Duration(config.GlobalServerConfig.Auth.AccessExpire) * time.Second,
//		MaxRefresh:  time.Hour,
//		IdentityKey: identityUsKey,
//		// PayloadFunc is used to define a custom token payload.
//		PayloadFunc: func(data interface{}) jwt.MapClaims {
//			if v, ok := data.(map[string]interface{}); ok {
//				return jwt.MapClaims{
//					identityUsKey: v,
//				}
//			}
//			return jwt.MapClaims{}
//		},
//		// IdentityHandler is used to define a custom identity handler.
//		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
//			claims := jwt.ExtractClaims(ctx, c)
//			payloadMap, ok := claims[identityUsKey].(map[string]interface{})
//			if !ok {
//				hlog.Error("get payloadMap error:", " claims data:", claims[identityUsKey])
//				return nil
//			}
//			c.Set("userId", payloadMap["userId"])
//
//			return payloadMap
//		},
//		// Authenticator is used to validate the login data.
//		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
//			res := new(member.MemberInfo)
//
//			var loginVal jwtUsLogin
//			if err := c.BindAndValidate(&loginVal); err != nil {
//				return "", err
//			}
//			// 验证码
//
//			venueId := int64(loginVal.VenueId)
//
//			res, err = service.NewMember(ctx, c).Login(&wx.MemberLoginReq{
//				Mobile:  loginVal.Mobile,
//				Captcha: loginVal.VerifyCode,
//				VenueId: venueId,
//			})
//			hlog.Info(res)
//			if err != nil {
//				hlog.Error(err, "jwtLogin error")
//				return nil, err
//			}
//			//jwt
//			var tokenInfo token.UserTokenInfo
//			tokenInfo.MemberId = res.ID
//			tokenInfo.Name = res.Name
//			tokenInfo.ExpiredAt = time.Now().Add(time.Duration(config.GlobalServerConfig.Auth.AccessExpire) * time.Second).Format(time.DateTime)
//
//			err = service.NewToken(ctx, c).Create(&tokenInfo)
//			if err != nil {
//				hlog.Error(err, "jwtLogin error, store token error")
//				return nil, err
//			}
//
//			payLoadMap := make(map[string]interface{})
//			payLoadMap["userId"] = strconv.Itoa(int(res.UserId))
//
//			return payLoadMap, nil
//		},
//		// Authorizator is used to validate the authentication of the current request.
//		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
//			//obj := string(c.URI().Path())
//			//act := string(c.Method())
//			payloadMap, ok := data.(map[string]interface{})
//			hlog.Info(payloadMap)
//			if !ok {
//				hlog.Error("get payloadMap error:", " claims data:", data)
//				return false
//			}
//
//			if payloadMap["userId"] == "" {
//				hlog.Error("暂无角色", err)
//				return false
//			}
//			userId := payloadMap["userId"].(string)
//			// check token is valid
//			userIdInt, err := strconv.Atoi(userId)
//			if err != nil {
//				hlog.Error("get payloadMap error:", err)
//				return false
//			}
//
//			existToken := service.NewToken(ctx, c).IsExistByUserID(int64(userIdInt))
//			if !existToken {
//				return false
//			}
//
//			return true
//		},
//		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
//			c.JSON(code, map[string]interface{}{
//				"code":    code,
//				"message": message,
//			})
//		},
//		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
//			utils.SendResponse(c, errno.Success,
//				map[string]interface{}{
//					"token":  token,
//					"expire": expire.Format(time.RFC3339),
//				}, 0, "")
//		},
//	})
//
//	return
//
//}
