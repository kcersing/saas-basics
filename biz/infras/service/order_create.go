package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	member2 "saas/biz/dal/db/ent/member"
	order2 "saas/biz/dal/db/ent/order"
	product2 "saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/user"
	"saas/biz/infras/do"
	"saas/idl_gen/model/order"
	"saas/pkg/enums"
	"saas/pkg/utils"
	"strconv"
	"time"
)

func (o Order) Buy(req *order.BuyReq) (orderOne *order.OrderInfo, err error) {

	product, err := o.db.Product.Query().Where(product2.ID(req.ProductId)).First(o.ctx)
	if err != nil {
		return nil, err
	}

	ok, err := o.db.Member.Query().Where(member2.ID(req.MemberId)).Exist(o.ctx)
	if !ok {
		return nil, errors.New("未找到会员信息")
	}

	orderTx := o.db.Order.
		Create().
		SetOrderSn(utils.CreateCn()).
		SetMemberID(req.MemberId).
		SetDevice(req.Device).
		SetVenueID(req.VenueId).
		SetProductType(product.Type).
		SetStatus(1).
		SetNature(enums.Buy)

	userId, exist := o.c.Get("userId")
	if exist || userId != nil {
		uId, ok := userId.(string)
		if ok {
			uid, _ := strconv.ParseInt(uId, 10, 64)
			create, _ := o.db.User.Query().Where(user.IDEQ(uid)).First(o.ctx)
			if create != nil {
				orderTx.SetOrderCreates(create)
			}
		}
	}

	one, err := orderTx.Save(o.ctx)

	if err != nil {
		return nil, errors.Wrap(err, "创建 Order 失败")
	}

	_, err = o.db.OrderItem.Create().
		SetOrder(one).
		SetName(product.Name).
		SetProductID(req.ProductId).
		SetNumber(req.Number).
		Save(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "创建 Order Item 失败")
		return nil, err
	}

	amount, err := o.db.OrderAmount.Create().
		SetOrder(one).
		SetTotal(req.Fee).
		SetResidue(req.Fee).
		Save(o.ctx)

	if err != nil {
		err = errors.Wrap(err, "创建Order Amount失败")
		return nil, err
	}

	orderOne = o.entOrderInfo(one)

	o.FinishOrder(do.FinishOrder{
		Sn:            one.OrderSn,
		Fee:           int64(amount.Total * 100),
		TransactionId: "123456789",
		Transaction:   nil,
	})

	return orderOne, nil
}
func (o Order) FinishOrder(req do.FinishOrder) (err error) {
	first, err := o.db.Order.Query().Where(order2.OrderSn(req.Sn)).First(o.ctx)
	if err != nil {
		return err
	}

	fee := float64(req.Fee / 100)
	amount, _ := first.QueryAmount().First(o.ctx)
	actual := amount.Actual + fee
	_, err = o.db.OrderAmount.UpdateOne(amount).
		SetActual(actual).
		Save(o.ctx)
	if err != nil {
		hlog.Info(req.Sn + "设置已付金额失败")
	}

	_, err = o.db.OrderPay.Create().
		SetOrder(first).
		SetPay(fee).
		SetPayAt(time.Now()).
		SetPayWay("wxPay").
		SetPaySn(req.Sn).
		SetPrepayID(req.TransactionId).
		SetPayExtra(req.Transaction).
		Save(o.ctx)
	if err != nil {
		hlog.Info(req.Sn + "OrderPay 创建失败")
	}
	if (amount.Total - fee) == 0 {

		_, err = o.db.Order.
			UpdateOne(first).
			SetCompletionAt(time.Now()).
			SetStatus(5).
			Save(o.ctx)
		if err != nil {
			hlog.Info(req.Sn + "完成时间更新失败")
		}

		go func() {
			NewMemberProduct(o.ctx, o.c).CreateMemberProduct(do.CreateMemberProductReq{
				MemberId:    first.MemberID,
				VenueId:     first.VenueID,
				Order:       first,
				OrderAmount: amount,
			})
		}()
	}

	return nil
}
