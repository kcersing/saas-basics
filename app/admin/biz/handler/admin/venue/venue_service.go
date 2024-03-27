// Code generated by hertz generator.

package venue

import (
	"context"
	"github.com/jinzhu/copier"
	"saas/app/admin/pkg/errno"
	"saas/app/admin/pkg/utils"
	"saas/app/pkg/do"
	"saas/app/pkg/service/admin"

	"github.com/cloudwego/hertz/pkg/app"
	venue "saas/app/admin/idl_gen/model/admin/venue"
)

// VenueList .
// @router /api/admin/venue/list [POST]
func VenueList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.VenueListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}

	var venueListReq do.VenueListReq
	err = copier.Copy(&venueListReq, &req)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	userList, total, err := admin.NewVenue(ctx, c).List(&venueListReq)
	if err != nil {
		utils.SendResponse(c, errno.ConvertErr(err), nil, 0, "")
		return
	}
	utils.SendResponse(c, errno.Success, userList, int64(total), "")
	return
}