package service

import (
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent/user"
	"saas/biz/infras/do"
	"saas/idl_gen/model/order"
	"saas/pkg/enums"
	"saas/pkg/utils"
	"strconv"
	"sync"
)

func (o Order) CreateSmsOrder(req do.CreateSmsOrderReq) (orderOne *order.OrderInfo, err error) {

	errChan := make(chan error, 2)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(2)

	tx, err := o.db.Tx(o.ctx)

	if err != nil {
		return nil, errors.Wrap(err, "starting a transaction:")
	}

	orderTx := tx.Order.
		Create().
		SetOrderSn(utils.CreateCn()).
		SetProductType(enums.Sms).
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

	one, err := orderTx.
		SetStatus(0).
		Save(o.ctx)

	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "创建 Order 失败"))
	}
	number := strconv.FormatInt(req.Number, 10)
	go func() {
		tx.OrderItem.Create().
			SetOrder(one).
			SetName("短信[" + number + "条]").
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建 Order Item 失败")
			errChan <- err
		}

		wg.Done()
	}()

	go func() {
		tx.OrderAmount.Create().
			SetOrder(one).
			SetTotal(req.Fee).
			SetResidue(req.Fee).
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建Order Amount失败")
			errChan <- err
		}

		wg.Done()
	}()

	wg.Wait()
	select {
	case result := <-errChan:
		return nil, rollback(tx, result)
	default:
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	orderOne = o.entOrderInfo(one)
	return orderOne, nil
}
