package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	product2 "saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/productcourses"
	"saas/biz/infras/do"
	"saas/pkg/enums"
	"saas/pkg/utils"
	"time"
)

func (m MemberProduct) CreateMemberProduct(req do.CreateMemberProductReq) error {
	orderItem, _ := req.Order.QueryItem().First(m.ctx)
	product, _ := m.db.Product.Query().Where(product2.ID(orderItem.ProductID)).First(m.ctx)

	validityAt := time.Now().Add(time.Duration(product.Deadline) * 24 * time.Hour)
	cancelAt := validityAt.Add(time.Duration(product.Duration) * 24 * time.Hour)
	price := req.OrderAmount.Total / float64(orderItem.Number)
	mpEnt := m.db.MemberProduct.
		Create().
		SetStatus(0).
		SetSn(utils.CreateCn()).
		SetType(product.Type).
		SetSubType(product.SubType).
		SetMemberID(req.MemberId).
		SetProductID(orderItem.ProductID).
		SetVenueID(req.VenueId).
		SetOrderID(req.Order.ID).
		SetName(product.Name).
		SetPrice(price).
		SetFee(req.OrderAmount.Total).
		SetDuration(product.Duration).
		SetLength(product.Length).
		SetDeadline(product.Deadline).
		SetValidityAt(validityAt).
		SetIsCourse(product.IsCourse).
		SetCancelAt(cancelAt)

	switch product.SubType {

	case enums.CoursePackage:
		// "私教课包"
		mpEnt.
			SetNumber(product.Times).
			SetNumberSurplus(product.Times)

	case enums.CourseOne:
		//一对一私教课
		mpEnt.
			SetNumber(orderItem.Number).
			SetNumberSurplus(orderItem.Number)
	case enums.CourseMore:
		//一对多私教课
		mpEnt.
			SetNumber(orderItem.Number).
			SetNumberSurplus(orderItem.Number)
	case enums.CardSub:
		//次卡
		mpEnt.
			SetNumber(product.Times).
			SetNumberSurplus(product.Times)
	default:

	}

	mp, err := mpEnt.Save(m.ctx)
	if err != nil {
		return err
	}

	all, _ := m.db.ProductCourses.Query().Where(productcourses.ProductID(orderItem.ProductID)).All(m.ctx)

	if all != nil {
		go func() {
			for _, v := range all {
				_, err := m.db.MemberProductCourses.Create().
					SetMemberProductID(mp.ID).
					SetName(v.Name).
					SetType(v.Type).
					SetNumber(v.Number).
					SetNumberSurplus(product.Times).
					SetCoursesID(v.ID).
					Save(m.ctx)
				if err != nil {
					hlog.Info(mp.Sn + "创建附属课程失败")
				}
			}
		}()
	}

	return err
}
