package service

import (
	member2 "saas/biz/dal/db/ent/member"
	product2 "saas/biz/dal/db/ent/product"
	"saas/biz/infras/do"
	"saas/pkg/utils"
)

func (m MemberProduct) CreateMemberProduct(req do.CreateMemberProductReq) error {

	product, _ := m.db.Product.Query().Where(product2.ID(req.ProductId)).First(m.ctx)

	member, _ := m.db.Member.Query().Where(member2.ID(req.MemberId)).First(m.ctx)

	price := float64(req.OrderAmount.Total) / float64(req.OrderItem.Number)
	mp, err := m.db.MemberProduct.
		Create().
		SetStatus(0).
		SetSn(utils.CreateCn()).
		SetType(product.Type).
		SetSubType(product.SubType).
		SetMemberID(req.MemberId).
		SetProductID(req.ProductId).
		SetVenueID(req.VenueId).
		SetOrderID(req.Order.ID).
		SetName(product.Name).
		SetPrice(price).
		SetFee(req.OrderAmount.Total).
		SetDuration(product.Duration).
		SetLength(product.Length).
		SetCount(product.Times).
		SetCountSurplus(product.Times).
		SetDeadline(product.Deadline).
		//SetValidityAt().
		//SetCancelAt().
		Save(m.ctx)
	if err != nil {
		return err
	}
	return err
}
