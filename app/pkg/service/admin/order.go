package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/order"
	"saas/pkg/db/ent/predicate"
	"time"
)

type Order struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (o Order) Update(req do.OrderInfo) error {
	_, err := o.db.Order.Update().
		Where(order.IDEQ(req.ID)).
		SetVenueID(req.VenueId).
		//SetSource(req.Source).
		//SetDevice(req.Device).
		Save(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Order failed")
		return err
	}

	return nil
}

func (o Order) List(req do.OrderListReq) (resp []*do.OrderInfo, total int, err error) {
	var predicates []predicate.Order
	if req.OrderSn != "" {
		predicates = append(predicates, order.OrderSnEQ(req.OrderSn))
	}
	lists, err := o.db.Order.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Order list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Order info failed")
		return resp, 0, err
	}

	for i, v := range lists {
		ven, _ := v.QueryOrderVenues().First(o.ctx)
		resp[i].VenueName = ven.Name
		amount, _ := v.QueryAmount().First(o.ctx)
		var amo do.OrderAmount
		err = copier.Copy(&amo, &amount)
		if err != nil {
			err = errors.Wrap(err, "copy Order info failed")
			return resp, 0, err
		}
		resp[i].OrderAmount = amo
		item, _ := v.QueryItem().First(o.ctx)
		var ite do.OrderItem
		err = copier.Copy(&ite, &item)
		if err != nil {
			err = errors.Wrap(err, "copy Order info failed")
			return resp, 0, err
		}
		resp[i].OrderItem = ite
		itemProductInfo, err := NewProduct(o.ctx, o.c).Info(ite.ProductId)
		if err == nil {
			resp[i].OrderItem.ProductName = itemProductInfo.Name
		} else {
			hlog.Error(err)
		}
		member, _ := v.QueryOrderMembers().First(o.ctx)
		resp[i].MemberName = member.Name
		cre, _ := v.QueryOrderCreates().First(o.ctx)
		resp[i].CreateName = cre.Nickname

		sales, _ := v.QuerySales().All(o.ctx)
		var sale []do.OrderSales
		err = copier.Copy(&sale, &sales)
		if err != nil {
			err = errors.Wrap(err, "copy Order Sales failed")
			return resp, 0, err
		}
		resp[i].OrderSales = sale
		for i2, orderSale := range resp[i].OrderSales {
			info, err := NewUser(o.ctx, o.c).Info(orderSale.SalesId)
			if err == nil {
				resp[i].OrderSales[i2].SalesName = info.Nickname
			} else {
				hlog.Error(err)
			}
		}
		resp[i].CreatedAt = v.CreatedAt.Format(time.DateTime)
		resp[i].CompletionAt = v.CreatedAt.Format(time.DateTime)

		//	OrderPay     []OrderPay   `json:"order_pay"`
	}

	total, _ = o.db.Order.Query().Where(predicates...).Count(o.ctx)
	return
}

func (o Order) UpdateStatus(id int64, status int64) error {
	_, err := o.db.Order.Update().Where(order.IDEQ(id)).SetStatus(int64(status)).Save(o.ctx)
	return err
}

func (o Order) Info(id int64) (info *do.OrderInfo, err error) {
	ret, err := o.db.Order.Query().
		Where(order.IDEQ(id)).
		Limit(1).
		First(o.ctx)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	info = &do.OrderInfo{
		ID: ret.ID,
	}
	return
}
func (o Order) GetBySnOrder(sn string) (info *do.OrderInfo, err error) {
	ret, err := o.db.Order.Query().
		Where(order.OrderSnEQ(sn)).
		Limit(1).
		All(o.ctx)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("OrdersNoFound")
	}
	row := ret[0]
	info = &do.OrderInfo{
		ID: row.ID,
	}
	return
}

func NewOrder(ctx context.Context, c *app.RequestContext) do.Order {
	return &Order{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
