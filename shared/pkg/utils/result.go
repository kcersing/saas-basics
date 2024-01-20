package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"time"
)

// Result 返回的结果：
type Result struct {
	Time      string      `json:"time"`
	Code      int         `json:"code"`
	Message   interface{} `json:"message"`
	Data      interface{} `json:"data"`
	CacheTime string      `json:"cacheTime"`
}

// Success 成功
func Success(ctx context.Context, c *app.RequestContext, data interface{}, message interface{}, cacheTime string) {

	res := Result{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		Code:      1,
		Message:   message,
		Data:      data,
		CacheTime: cacheTime,
	}
	c.JSON(consts.StatusOK, res)
	return
}

// Error 出错
func Error(ctx context.Context, c *app.RequestContext, code int, message interface{}, data interface{}, cacheTime string) {
	res := Result{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		Code:      code,
		Message:   message,
		Data:      data,
		CacheTime: cacheTime,
	}
	c.JSON(consts.StatusOK, res)
	return
}
