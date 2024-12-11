package service

import (
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent/member"
	order2 "saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/orderitem"
	"saas/biz/dal/db/ent/predicate"
	"saas/idl_gen/model/order"
	"time"
)

func (o Order) List(req *order.ListOrderReq) (resp []*order.OrderInfo, total int, err error) {
	var predicates []predicate.Order
	if req.OrderSn != "" {
		predicates = append(predicates, order2.OrderSnEQ(req.OrderSn))
	}
	if req.Mobile != "" {
		predicates = append(predicates, order2.HasOrderMembersWith(member.MobileEQ(req.Mobile)))
	}
	if req.MemberName != "" {
		predicates = append(predicates, order2.HasOrderMembersWith(member.NameEQ(req.MemberName)))
	}
	if req.Name != "" {
		predicates = append(predicates, order2.HasItemWith(orderitem.NameEQ(req.Name)))
	}
	if len(req.Status) > 0 {
		predicates = append(predicates, order2.StatusIn(req.Status...))
	}
	if req.ProductType != "" {
		predicates = append(predicates, order2.ProductTypeEQ(req.ProductType))
	}
	if req.Nature != "" {
		predicates = append(predicates, order2.NatureEQ(req.Nature))
	}
	if len(req.VenueId) > 0 {
		predicates = append(predicates, order2.VenueIDIn(req.VenueId...))
	}

	if req.StartCompletionAt != "" && req.EndCompletionAt != "" {
		signStartAt, _ := time.Parse(time.DateTime, req.StartCompletionAt)
		signEndAt, _ := time.Parse(time.DateTime, req.EndCompletionAt)

		predicates = append(predicates, order2.CompletionAtGTE(signStartAt))
		predicates = append(predicates, order2.CompletionAtLTE(signEndAt))
	}

	//lt：less than 小于
	//le：less than or equal to 小于等于
	//eq：equal to 等于
	//ne：not equal to 不等于
	//ge：greater than or equal to 大于等于
	//gt：greater than 大于

	lists, err := o.db.Order.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Order list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, o.entOrderInfo(v))
	}

	total, _ = o.db.Order.Query().Where(predicates...).Count(o.ctx)
	return

}
