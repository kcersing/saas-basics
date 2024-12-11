package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	order2 "saas/biz/dal/db/ent/order"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/order"
	"time"
)

type Order struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewOrder(ctx context.Context, c *app.RequestContext) do.Order {
	return &Order{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (o Order) Info(id int64) (info *order.OrderInfo, err error) {
	first, err := o.db.Order.Query().
		Where(order2.IDEQ(id)).
		Limit(1).
		First(o.ctx)
	if err != nil {
		return nil, err
	}

	info = o.entOrderInfo(first)
	return
}
func (o Order) entOrderInfo(v *ent.Order) *order.OrderInfo {
	completionAt := v.CompletionAt.Format(time.DateTime)
	createdAt := v.CreatedAt.Format(time.DateTime)
	updatedAt := v.UpdatedAt.Format(time.DateTime)

	var memberName, memberMobile string
	m, _ := o.db.Member.Query().Where(member2.IDEQ(v.MemberID)).First(o.ctx)

	if m != nil {
		memberName = m.Name
		memberMobile = m.Mobile
	}

	return &order.OrderInfo{
		ID:           v.ID,
		OrderSn:      v.OrderSn,
		Status:       v.Status,
		Source:       v.Source,
		Device:       v.Device,
		Nature:       v.Nature,
		ProductType:  v.ProductType,
		VenueId:      v.VenueID,
		MemberId:     v.MemberID,
		CreateId:     v.CreatedID,
		CompletionAt: completionAt,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		OrderAmount:  o.entOrderAmount(v),
		OrderItem:    o.entOrderItem(v),
		OrderPay:     o.entOrderPay(v),
		OrderSales:   o.entOrderSales(v),
		MemberName:   memberName,
		MemberMobile: memberMobile,
	}
}
func (o Order) entOrderSales(v *ent.Order) (sales []*order.OrderSales) {
	all, err := v.QuerySales().All(o.ctx)
	if err != nil {
		return nil
	}
	for _, item := range all {
		sales = append(sales, &order.OrderSales{
			ID:      item.ID,
			SalesId: item.SalesID,
			Ratio:   item.Ratio,
			OrderId: item.OrderID,
		})
	}
	return
}
func (o Order) entOrderPay(v *ent.Order) (pays []*order.OrderPay) {

	all, err := v.QueryPay().All(o.ctx)
	if err != nil {
		return nil
	}
	for _, item := range all {

		pays = append(pays, &order.OrderPay{
			ID:        item.ID,
			OrderId:   item.OrderID,
			Pay:       item.Pay,
			Remission: item.Remission,
			PayWay:    item.PayWay,
			PaySn:     item.PaySn,
			PrepayId:  item.PrepayID,
			//PayExtra:  item.PayExtra,
			Note:      item.Note,
			CreatedAt: item.CreatedAt.Format(time.DateTime),
			UpdatedAt: item.UpdatedAt.Format(time.DateTime),
		},
		)
	}

	return
}

func (o Order) entOrderItem(v *ent.Order) *order.OrderItem {
	first, err := v.QueryItem().First(o.ctx)
	if err != nil {
		return nil
	}

	return &order.OrderItem{
		ID:                   first.ID,
		ProductId:            first.ProductID,
		RelatedUserProductId: first.RelatedUserProductID,
		ContestId:            first.ContestID,
		BootcampId:           first.BootcampID,
		OrderId:              first.OrderID,
		Name:                 first.Name,
		//Data:                 data,
		CreatedAt: first.CreatedAt.Format(time.DateTime),
		UpdatedAt: first.UpdatedAt.Format(time.DateTime),
	}
}

func (o Order) entOrderAmount(v *ent.Order) *order.OrderAmount {

	first, err := v.QueryAmount().First(o.ctx)
	if err != nil {
		return nil
	}
	return &order.OrderAmount{
		ID:        first.ID,
		Total:     first.Total,
		Actual:    first.Actual,
		Residue:   first.Residue,
		Remission: first.Remission,
		OrderId:   first.OrderID,
		CreatedAt: first.CreatedAt.Format(time.DateTime),
		UpdatedAt: first.UpdatedAt.Format(time.DateTime),
	}
}
