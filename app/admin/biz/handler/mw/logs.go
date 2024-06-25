package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/app/pkg/do"
	"saas/app/pkg/service/admin"
	"strconv"
	"time"
)

func LogMw() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		c.Next(ctx)
		var logs do.LogsInfo
		logs.Type = "Interface"
		logs.Method = string(c.Request.Method())
		logs.Api = string(c.Request.Path())
		logs.UserAgent = string(c.Request.Header.UserAgent())
		logs.Ip = c.ClientIP()

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
		logs.Time = int32(costTime)

		v, exist := c.Get("user_id")
		if exist || v == nil {
			v = "0"
		}
		var userIDStr string
		var username = "Anonymous"
		var ok bool

		userIDStr, ok = v.(string)

		if !ok {
			userIDStr = "0"
		}

		userID, _ := strconv.Atoi(userIDStr)

		userInfo, _ := admin.NewUser(ctx, c).Info(int64(userID))

		if userInfo != nil {
			username = userInfo.Nickname
		}

		logs.Operator = username

		err := admin.NewLogs(ctx, c).Create(&logs)

		if err != nil {
			hlog.Error(err)
		}

	}
}
