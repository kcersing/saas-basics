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
		SetStatus(1).
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
		SetCancelAt(cancelAt)

	switch product.SubType {

	case enums.CoursePackage:
		// "私教课包"
		mpEnt.
			SetNumber(product.Times).
			SetNumberSurplus(product.Times).
			SetIsCourse(product.IsCourse)

	case enums.CourseOne:
		//一对一私教课
		mpEnt.
			SetNumber(orderItem.Number).
			SetNumberSurplus(orderItem.Number).
			SetIsCourse(0)
	case enums.CourseMore:
		//一对多私教课
		mpEnt.
			SetNumber(orderItem.Number).
			SetNumberSurplus(orderItem.Number).
			SetIsCourse(0)
	case enums.CardSub:
		//次卡
		mpEnt.
			SetNumber(product.Times).
			SetNumberSurplus(product.Times).
			SetIsCourse(0)
	default:

	}

	mp, err := mpEnt.Save(m.ctx)
	if err != nil {
		return err
	}

	go func() {
		contract, _ := product.QueryContracts().First(m.ctx)
		if contract != nil {
			hlog.Info(mp.Sn + "查找合同失败")
		} else {
			memberContract, err := m.db.MemberContract.Create().
				SetContractID(contract.ID).
				SetVenueID(mp.VenueID).
				SetProductID(product.ID).
				SetName(contract.Name).
				SetMemberID(mp.MemberID).
				SetMemberProductID(mp.ID).
				SetOrderID(req.Order.ID).
				Save(m.ctx)
			if err != nil {
				hlog.Info(mp.Sn + "创建合同失败")
			}
			_, err = m.db.MemberContractContent.Create().
				SetMemberContractID(memberContract.ID).
				SetContent(contract.Content).
				SetContractID(memberContract.ID).Save(m.ctx)
			if err != nil {
				hlog.Info(mp.Sn + "创建合同内容失败")
			}
		}

	}()

	all, _ := m.db.ProductCourses.Query().Where(productcourses.ProductID(orderItem.ProductID)).All(m.ctx)

	if all != nil {
		go func() {
			for _, v := range all {
				_, err := m.db.MemberProductCourses.Create().
					SetMemberProductID(mp.ID).
					SetName(v.Name).
					SetType(v.Type).
					SetCoursesID(v.CoursesID).
					Save(m.ctx)
				if err != nil {
					hlog.Info(mp.Sn + "创建附属课程失败")
				}
			}
		}()
	}

	return err
}
