package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/biz/infras/service"
	"saas/idl_gen/model/auth"
	"strconv"
	"time"
)

func LogMw() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		c.Next(ctx)
		var logs auth.LogsInfo
		logs.Type = "Interface"
		logs.Method = string(c.Request.Method())
		logs.API = string(c.Request.Path())
		logs.UserAgent = string(c.Request.Header.UserAgent())
		logs.IP = c.ClientIP()

		reqBodyStr := string(c.Request.Body())
		if len(reqBodyStr) > 200 {
			reqBodyStr = reqBodyStr[:200]
		}
		logs.ReqContent = reqBodyStr

		respBodyStr := string(c.Request.Body())
		if len(respBodyStr) > 200 {
			respBodyStr = respBodyStr[:200]
		}

		if c.Response.Header.StatusCode() == 200 {
			logs.Success = true
		}

		costTime := time.Since(start).Milliseconds()
		logs.Time = int64(int32(costTime))
		var username string
		v, exist := c.Get("userId")
		if exist || v == nil {
			username = "member"
			v, exist = c.Get("memberId")
			if exist || v == nil {
				v = "0"
			}
		} else {
			username = "Anonymous"

		}
		var userIDStr string

		var ok bool

		userIDStr, ok = v.(string)

		if !ok {
			userIDStr = "0"
		}

		userID, err := strconv.Atoi(userIDStr)

		if err != nil {
			userID = 0
		}

		if username == "member" {
			userInfo, err := service.NewMember(ctx, c).MemberInfo(int64(userID))
			if err != nil {
				hlog.Error(err)
			}
			if userInfo != nil {
				username = userInfo.Name
			}
		} else {
			userInfo, err := service.NewUser(ctx, c).Info(int64(userID))
			if err != nil {
				hlog.Error(err)
			}
			if userInfo != nil {
				username = userInfo.Name
			}
		}

		logs.Operatorsr = username

		err = service.NewLogs(ctx, c).Create(&logs)

		if err != nil {
			hlog.Error(err)
		}

	}
}
