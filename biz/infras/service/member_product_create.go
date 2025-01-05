package service

import (
	product2 "saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/productcourses"
	"saas/biz/infras/do"
	"saas/pkg/utils"
)

func (m MemberProduct) CreateMemberProduct(req do.CreateMemberProductReq) error {

	product, _ := m.db.Product.Query().Where(product2.ID(req.ProductId)).First(m.ctx)

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

	all, _ := m.db.ProductCourses.Query().Where(productcourses.ProductID(req.ProductId)).All(m.ctx)

	if all != nil {
		go func() {
			for _, v := range all {
				m.db.MemberProductCourses.Create().
					SetMemberProductID(mp.ID).
					SetName(v.Name).
					SetType(v.Type).
					SetNumber(v.Number).
					SetCoursesID(v.ID).
					Save(m.ctx)
			}
		}()
	}

	return err
}
