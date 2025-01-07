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
		SetDevice(req.Device).
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

	_, err = o.db.OrderAmount.Create().
		SetOrder(one).
		SetTotal(req.Fee).
		SetResidue(req.Fee).
		Save(o.ctx)

	if err != nil {
		err = errors.Wrap(err, "创建Order Amount失败")
		return nil, err
	}

	orderOne = o.entOrderInfo(one)
	return orderOne, nil
}
