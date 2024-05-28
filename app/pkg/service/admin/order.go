package admin

import (
	"context"
	"fmt"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/contract"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/order"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/product"
	"saas/pkg/db/ent/productproperty"
	venue2 "saas/pkg/db/ent/venue"
	"saas/pkg/utils"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	hlog.Info(fmt.Sprintf("%+v", req))
	hlog.Info("=============")
	return nil

	one := &ent.Order{}
	mp := &ent.MemberProduct{}
	var err error

	venue, err := o.db.Venue.Query().Where(venue2.IDEQ(req.Venue)).First(o.ctx)
	if err != nil {
		return errors.Wrap(err, "未找到场馆")
	}

	members, err := o.db.Member.Query().Where(member.IDEQ(req.Member)).First(o.ctx)

	if err != nil {
		return errors.Wrap(err, "未找到会员")
	}

	product, err := o.db.Product.Query().Where(product.IDEQ(req.Product)).First(o.ctx)
	if err != nil {
		return errors.Wrap(err, "未找到产品")
	}
	// create, err := o.db.User.Query().Where(user.IDEQ(req.Create)).First(o.ctx)
	// if err != nil {
	// 	return errors.Wrap(err, "未找到创建人")
	// }

	errChan := make(chan error, 15)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(4)

	layout := "2006-01-02 15:04:05"
	assignAt, err := time.Parse(layout, req.AssignAt)
	if err != nil {
		err = errors.Wrap(err, "时间转换失败")
		errChan <- err
	}

	go func() {
		one, err = o.db.Order.Create().
			SetOrderSn(utils.CreateCn()).
			SetOrderVenues(venue).
			SetOrderMembers(members).
			// SetOrderCreates(create).
			SetStatus(0).
			//SetSource(req.Source).
			//SetDevice(req.Device).
			// req.NatureType
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建 Order 失败")
			errChan <- err
		}
		wg.Done()
	}()
	go func() {

		_, err := o.db.OrderItem.Create().
			SetAufk(one).
			SetProductID(req.Product).
			// SetData(fmt.Sprintf("%+v", req)).
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建 Order Item 失败")
			errChan <- err
		}
		wg.Done()
	}()
	go func() {
		mp, err = o.db.MemberProduct.Create().
			SetProductID(req.Product).
			// SetVenue (req.Venue).
			SetSn(utils.CreateCn()).
			SetMembers(members).
			SetOrderID(one.ID).
			SetValidityAt(assignAt).
			SetName(product.Name).
			SetStatus(0).
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建会员产品失败")
			errChan <- err
		}
		wg.Done()
	}()

	go func() {
		_, err = o.db.OrderAmount.Create().
			SetAufk(one).
			SetTotal(req.Total).
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建Order Amount失败")
			errChan <- err
		}
		wg.Done()
	}()

	go func() {
		for _, v := range req.Staffs {
			_, err = o.db.OrderSales.Create().
				SetAufk(one).
				SetSalesID(v.Id).
				SetRatio(v.Ratio).
				SetMemberID(members.ID).
				Save(o.ctx)
			if err != nil {
				err = errors.Wrap(err, "创建Order Sales失败")
				errChan <- err
			}

		}
		wg.Done()
	}()

	go func() {

		var property []do.PropertyItem
		property = append(property, req.CardProperty...)
		property = append(property, req.CourseProperty...)
		property = append(property, req.ClassProperty...)

		if len(property) > 0 {
			for _, v := range property {
				p, err := o.db.ProductProperty.Query().
					Where(productproperty.IDEQ(v.Property)).
					First(o.ctx)

				if err != nil {
					err = errors.Wrap(err, "查询Product Property失败")
				}

				// venues, err := o.db.Venue.Query().
				// 	QueryPropertyVenues().
				// 	Where(productproperty.IDEQ(v.Property)).
				// 	All(o.ctx)

				// if err != nil {
				// 	err = errors.Wrap(err, "查询Product Property venues失败")

				// }

				_, err = o.db.MemberProductProperty.Create().
					SetMemberID(members.ID).
					SetOwner(mp).
					SetType(p.Type).
					SetPropertyID(p.ID).
					SetPrice(p.Price).
					SetDuration(p.Duration).
					SetLength(p.Length).
					SetName(p.Name).
					SetCount(v.Quantity).
					SetCountSurplus(v.Quantity).
					// AddVenues(&venues...).
					SetStatus(0).
					Save(o.ctx)
				if err != nil {
					err = errors.Wrap(err, "创建Member Product Property失败")
					errChan <- err
				}

			}
		}
		wg.Done()
	}()

	go func() {

		for i, v := range req.Contract {

			contract, err := o.db.Contract.Query().
				Where(contract.IDEQ(v)).
				First(o.ctx)
			if err != nil {
				err = errors.Wrap(err, "合同 Contract 失败")
				errChan <- err
			}
			memberContract, err := o.db.MemberContract.Create().
				SetMember(members).
				SetOrder(one).
				SetMemberProduct(mp).
				SetName(contract.Name).
				SetStatus(0).
				Save(o.ctx)
			if err != nil {
				err = errors.Wrap(err, "创建Member Contract "+string(i)+" 失败")
				errChan <- err
			}

			_, err = o.db.MemberContractContent.Create().
				SetContract(memberContract).
				SetContent(contract.Content).
				SetSignImg(req.SignImg).
				Save(o.ctx)

			if err != nil {
				err = errors.Wrap(err, "创建 Member Contract Content 失败")
				errChan <- err
			}
		}
		wg.Done()
	}()

	wg.Wait()
	select {
	case result := <-errChan:
		return result
	default:
	}
	return nil
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
