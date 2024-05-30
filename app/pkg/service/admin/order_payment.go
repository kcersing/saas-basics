package admin

import (
	"github.com/pkg/errors"
	"saas/app/pkg/do"
	order2 "saas/pkg/db/ent/order"
)

func (o Order) UnifyPay(req do.UnifyPayReq) error {
	order, err := o.db.Order.Query().Where(
		order2.OrderSn(req.OrderSn),
		order2.StatusIn(0, 1),
	).
		First(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "获取订单失败")
		return err
	}
	amount, err := order.QueryAmount().First(o.ctx)

	if err != nil {
		err = errors.Wrap(err, "获取订单金额失败")
		return err
	}

	if (amount.Residue - req.Remission - req.Payment) < 0 {
		err = errors.Wrap(err, "订单支付异常")
		return err
	}

	_, err = o.db.OrderPay.Create().
		SetOrder(order).
		SetPay(req.Payment).
		SetRemission(req.Remission).
		SetNote(req.Note).
		SetPayWay("线下支付").
		Save(o.ctx)

	if err != nil {
		err = errors.Wrap(err, "创建支付信息失败")
		return err
	}

	_, err = o.db.OrderAmount.Update().Where().
		SetResidue(amount.Residue - req.Remission - req.Payment).
		SetRemission(amount.Remission + req.Remission).
		SetActual(amount.Actual + req.Payment).
		Save(o.ctx)

	if err != nil {
		return err
	}

	if (amount.Residue - req.Remission - req.Payment) == 0 {
		_, err = o.db.Order.Update().Where().SetStatus(2).Save(o.ctx)
		if err != nil {
			return err
		}
		_, err = o.db.MemberProduct.Update().Where().SetStatus(1).Save(o.ctx)
		if err != nil {
			return err
		}
		_, err = o.db.MemberProductProperty.Update().Where().SetStatus(1).Save(o.ctx)
		if err != nil {
			return err
		}
	}

	return nil

}

func (o Order) QRPay(req do.QRPayReq) (string, error) {
	return "htpp://127.0.0.1", nil
}
