// Code generated by hertz generator.

package pub

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"path"
	"saas/app/admin/config"
	"saas/app/admin/pkg/errno"
	"saas/app/admin/pkg/minio"
	"saas/app/admin/pkg/utils"
	"strconv"
	"time"
)

// Upload .
// @router /api/pub/upload/ [POST]
func Upload(ctx context.Context, c *app.RequestContext) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	nowTime := time.Now()
	filename := minio.NewFileName(0, nowTime.UnixMicro())
	file.Filename = filename + path.Ext(file.Filename)
	uploadinfo, err := minio.PutToBucket(ctx, config.GlobalServerConfig.Minio.ImgBucketName, file)
	hlog.CtxInfof(ctx, "image upload size:"+strconv.FormatInt(uploadinfo.Size, 10))
	if err != nil {
		hlog.CtxInfof(ctx, "err:"+err.Error())
	}
	url := minio.URLconvert(ctx, c, config.GlobalServerConfig.Minio.ImgBucketName+"/"+uploadinfo.Key)
	utils.SendResponse(c, errno.Success, map[string]string{
		"name": uploadinfo.Key,
		"url":  url,
		"path": config.GlobalServerConfig.Minio.ImgBucketName + "/" + uploadinfo.Key,
	}, 1, "")

	return
}
