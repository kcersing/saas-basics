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
)

type Order struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (o Order) Create(req do.CreateOrder) error {

	hlog.Info("=============")
	hlog.Info(req)
	return nil
	//one := &ent.Order{}
	//var err error
	//
	//venue, err := o.db.Venue.Query().Where(venue2.IDEQ(req.VenueId)).First(o.ctx)
	//if err != nil {
	//	return errors.Wrap(err, "未找到场馆")
	//}
	//
	//members, err := o.db.Member.Query().Where(member.IDEQ(req.MemberId)).First(o.ctx)
	//
	//if err != nil {
	//	return errors.Wrap(err, "未找到会员")
	//}
	//
	//create, err := o.db.User.Query().Where(user.IDEQ(req.CreateId)).First(o.ctx)
	//if err != nil {
	//	return errors.Wrap(err, "未找到创建人")
	//}
	//
	//errChan := make(chan error, 15)
	//defer close(errChan)
	//var wg sync.WaitGroup
	//wg.Add(4)
	//
	//go func() {
	//	one, err = o.db.Order.Create().
	//		SetOrderSn(utils.CreateCn()).
	//		SetOrderVenues(venue).
	//		SetOrderMembers(members).
	//		SetOrderCreates(create).
	//		SetStatus(0).
	//		//SetSource(req.Source).
	//		//SetDevice(req.Device).
	//		Save(o.ctx)
	//	if err != nil {
	//		err = errors.Wrap(err, "创建订单失败")
	//		errChan <- err
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	_, err = o.db.OrderAmount.Create().
	//		SetAufk(one).
	//		SetTotal(req.Total).
	//		Save(o.ctx)
	//	if err != nil {
	//		err = errors.Wrap(err, "创建Order Amount失败")
	//		errChan <- err
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	for _, v := range req.Sales {
	//		_, err = o.db.OrderSales.Create().
	//			SetAufk(one).
	//			SetSalesID(v).
	//			Save(o.ctx)
	//		if err != nil {
	//			err = errors.Wrap(err, "创建Order Sales失败")
	//			errChan <- err
	//		}
	//
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	layout := "2006-01-02 15:04:05" // Go中的时间格式化参考时间
	//	assignAt, err := time.Parse(layout, req.AssignAt)
	//	if err != nil {
	//		err = errors.Wrap(err, "时间转换失败")
	//		errChan <- err
	//	}
	//
	//	var buffer bytes.Buffer
	//	for _, num := range req.ContractId {
	//		buffer.WriteString(strconv.Itoa(int(num)))
	//		buffer.WriteByte(',')
	//	}
	//	str := buffer.String()
	//
	//	_, err = o.db.OrderItem.Create().
	//		SetAufk(one).
	//		SetProductID(req.ProductId).
	//		SetQuantity(req.Quantity).
	//		SetContractID(str).
	//		SetAssignAt(assignAt).
	//		Save(o.ctx)
	//	if err != nil {
	//		err = errors.Wrap(err, "创建Order Item失败")
	//		errChan <- err
	//	}
	//	wg.Done()
	//}()
	//
	//wg.Wait()
	//select {
	//case result := <-errChan:
	//	return result
	//default:
	//}
	//return nil
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
