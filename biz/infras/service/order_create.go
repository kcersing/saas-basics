package service

import (
	"github.com/pkg/errors"
	member2 "saas/biz/dal/db/ent/member"
	product2 "saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/user"
	"saas/idl_gen/model/order"
	"saas/pkg/enums"
	"saas/pkg/utils"
	"strconv"
	"sync"
)

func (o Order) Buy(req *order.BuyReq) (orderOne *order.OrderInfo, err error) {

	var productName string
	product, err := o.db.Product.Query().Where(product2.ID(req.ProductId)).First(o.ctx)
	if err != nil {
		return nil, err
	}
	if product != nil {
		productName = product.Name
	}

	ok, err := o.db.Member.Query().Where(member2.ID(req.MemberId)).Exist(o.ctx)
	if !ok {
		return nil, errors.New("未找到会员信息")
	}
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
		SetMemberID(req.MemberId).
		SetDevice(req.Device).
		SetVenueID(req.VenueId).
		SetProductType(enums.Contest).
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
		SetNature(enums.Buy).
		SetDevice(req.Device).Save(o.ctx)

	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "创建 Order 失败"))
	}

	go func() {
		tx.OrderItem.Create().
			SetOrder(one).
			SetName(productName).
			SetProductID(req.ProductId).
			SetNumber(req.Number).
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
