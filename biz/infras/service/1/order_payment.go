package admin

import (
	"fmt"
	"github.com/pkg/errors"
	"saas/app/pkg/do"
	"saas/pkg/db/ent/memberproduct"
	"saas/pkg/db/ent/memberproductproperty"
	order2 "saas/pkg/db/ent/order"
	"strconv"
	"time"
)

func (o Order) UnifyPay(req do.UnifyPayReq) error {
	tx, err := o.db.Tx(o.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	order, err := tx.Order.Query().Where(
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
		err = errors.Wrap(errors.New("金额不能为负数 :"+fmt.Sprintf("%+v", amount.Residue-req.Remission-req.Payment)), "订单支付异常")
		return err
	}

	userId, exist := o.c.Get("user_id")
	if !exist || userId == nil {
		return errors.Wrap(errors.New("Unauthorized"), "Unauthorized")
	}

	uId, ok := userId.(string)
	if !ok {
		return errors.Wrap(errors.New("user id 转换失败"), "user id 转换失败")
	}
	uid, _ := strconv.ParseInt(uId, 10, 64)

	_, err = tx.OrderPay.Create().
		SetOrder(order).
		SetPay(req.Payment).
		SetRemission(req.Remission).
		SetNote(req.Note).
		SetPayWay("线下支付").
		SetCreateID(uid).
		Save(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "创建支付信息失败")
		return rollback(tx, errors.Wrap(err, "创建支付信息失败"))
	}

	_, err = tx.OrderAmount.Update().Where().
		SetResidue(amount.Residue - req.Remission - req.Payment).
		SetRemission(amount.Remission + req.Remission).
		SetActual(amount.Actual + req.Payment).
		Save(o.ctx)

	if err != nil {
		return rollback(tx, errors.Wrap(err, "订单金额修改失败"))
	}

	if (amount.Residue - req.Remission - req.Payment) == 0 {

		_, err = tx.Order.Update().Where(order2.IDEQ(order.ID)).SetStatus(2).SetCompletionAt(time.Now()).Save(o.ctx)

		if err != nil {
			return rollback(tx, errors.Wrap(err, "会员产品属性修改失败"))
		}
		_, err = tx.MemberProduct.Update().
			Where(memberproduct.IDEQ(order.MemberProductID)).
			SetPrice(amount.Actual + req.Payment).
			SetStatus(1).
			Save(o.ctx)

		if err != nil {
			return rollback(tx, errors.Wrap(err, "会员产品属性修改失败"))
		}
		_, err = tx.MemberProductProperty.Update().Where(memberproductproperty.MemberProductID(order.MemberProductID)).SetStatus(1).Save(o.ctx)
		if err != nil {
			return rollback(tx, errors.Wrap(err, "会员产品属性修改失败"))
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil

}

func (o Order) QRPay(req do.QRPayReq) (string, error) {
	return "htpp://127.0.0.1", nil
}
