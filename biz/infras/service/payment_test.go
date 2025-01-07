package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"saas/biz/infras/do"
	"testing"
)

func TestMemberOrder(t *testing.T) {
	//context.Context, c *app.RequestContext
	c := context.Background()
	c2 := app.RequestContext{}

	r := do.CreateMemberProductReq{
		MemberId:    100012,
		ProductId:   100019,
		VenueId:     100002,
		Order:       nil,
		OrderAmount: nil,
		OrderItem:   nil,
	}
	NewMemberProduct(c, &c2).CreateMemberProduct(r)
}
