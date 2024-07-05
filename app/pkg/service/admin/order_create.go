package admin

import (
	"bytes"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/pkg/minio"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/contract"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/order"
	"saas/pkg/db/ent/product"
	"saas/pkg/db/ent/productproperty"
	"saas/pkg/db/ent/user"
	venue2 "saas/pkg/db/ent/venue"
	"saas/pkg/utils"
	"strconv"
	"sync"
	"time"
)

func (o Order) Create(req do.CreateOrder) (string, error) {

	one := &ent.Order{}
	mp := &ent.MemberProduct{}
	var err error

	userId, exist := o.c.Get("user_id")
	if !exist || userId == nil {
		return "", errors.Wrap(err, "Unauthorized")
	}

	uId, ok := userId.(string)
	if !ok {
		return "", errors.Wrap(err, "user id 转换失败")
	}
	uid, _ := strconv.ParseInt(uId, 10, 64)

	create, err := o.db.User.Query().Where(user.IDEQ(uid)).First(o.ctx)
	if err != nil {
		return "", errors.Wrap(err, "未找到创建人")
	}

	venue, err := o.db.Venue.Query().Where(venue2.IDEQ(req.Venue)).First(o.ctx)
	if err != nil {
		return "", errors.Wrap(err, "未找到场馆")
	}

	members, err := o.db.Member.Query().Where(member.IDEQ(req.Member)).First(o.ctx)

	if err != nil {
		return "", errors.Wrap(err, "未找到会员")
	}

	products, err := o.db.Product.Query().Where(product.IDEQ(req.Product)).First(o.ctx)
	if err != nil {
		return "", errors.Wrap(err, "未找到产品")
	}

	layout := time.DateOnly
	assignAt, err := time.Parse(layout, req.AssignAt)
	if err != nil {
		return "", errors.Wrap(err, "时间转换失败")
	}

	errChan := make(chan error, 99)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(5)

	tx, err := o.db.Tx(o.ctx)

	if err != nil {
		return "", errors.Wrap(err, "starting a transaction:")
	}

	one, err = tx.Order.Create().
		SetOrderSn(utils.CreateCn()).
		SetOrderVenues(venue).
		SetOrderMembers(members).
		SetOrderCreates(create).
		SetStatus(0).
		SetNature(req.NatureType).
		//SetSource(req.Source).
		//SetDevice(req.Device).
		Save(o.ctx)
	if err != nil {
		return "", rollback(tx, errors.Wrap(err, "创建 Order 失败"))
	}
	mp, err = tx.MemberProduct.Create().
		SetProductID(req.Product).
		// SetVenue (req.Venue).
		SetSn(utils.CreateCn()).
		SetMembers(members).
		SetVenueID(venue.ID).
		SetOrderID(one.ID).
		SetName(products.Name).
		SetStatus(0).
		Save(o.ctx)
	if err != nil {
		return "", rollback(tx, errors.Wrap(err, "创建会员产品失败"))
	}
	_, err = tx.Order.Update().
		Where(order.ID(one.ID)).
		SetMemberProductID(mp.ID).
		SetOrderVenues(venue).
		Save(o.ctx)

	if err != nil {
		return "", rollback(tx, errors.Wrap(err, "创建 Order 失败"))
	}

	go func() {
		_, err := tx.OrderItem.Create().
			SetOrder(one).
			SetProductID(products.ID).
			//SetData(fmt.Sprintf("%+v", req)).
			SetData(req).
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建 Order Item 失败")
			errChan <- err
		}

		wg.Done()
	}()

	go func() {
		_, err = tx.OrderAmount.Create().
			SetOrder(one).
			SetTotal(req.Total).
			SetResidue(req.Total).
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建Order Amount失败")
			errChan <- err
		}

		wg.Done()
	}()

	go func() {
		for _, v := range req.Staffs {
			_, err = tx.OrderSales.Create().
				SetOrder(one).
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
				if v.Property > 0 {
					p, err := o.db.ProductProperty.Query().
						Where(productproperty.IDEQ(v.Property)).
						First(o.ctx)
					if err != nil {
						err = errors.Wrap(err, "查询Product Property失败")
						errChan <- err
					} else {
						venues, err := p.QueryVenues().All(o.ctx)
						if err != nil {
							err = errors.Wrap(err, "查询Product Property venues失败")
							errChan <- err
						}

						_, err = tx.MemberProductProperty.Create().
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
							AddVenues(venues...).
							SetValidityAt(assignAt).
							SetStatus(0).
							Save(o.ctx)
						//tx.MemberProductProperty.Update().Where(memberproductproperty.ID(pp.ID)).SetSn(pp.ID).Save()
						if err != nil {
							err = errors.Wrap(err, "创建Member Product Property失败")
							errChan <- err
						}
					}
				}
			}
		}
		wg.Done()
	}()

	go func() {
		signImg := ""
		if req.SignImg != "" {
			filename := minio.NewFileName(0, time.Now().UnixMicro()) + ".png"
			buffer := new(bytes.Buffer)
			// 将字符串写入Buffer
			_, err = buffer.WriteString(req.SignImg)
			if err != nil {
				err = errors.Wrap(err, "写入失败")
				errChan <- err
			}
			uploadInfo, err := minio.PutToBucketByBuf(o.ctx, config.GlobalServerConfig.Minio.ImgBucketName, filename, buffer)

			if err != nil {
				err = errors.Wrap(err, "签名上传失败")
				errChan <- err
			}
			signImg = config.GlobalServerConfig.Minio.ImgBucketName + "/" + uploadInfo.Key
		}

		for i, v := range req.Contract {
			contracts, err := o.db.Contract.Query().
				Where(contract.IDEQ(v)).
				First(o.ctx)
			if err != nil {
				err = errors.Wrap(err, "合同 Contract 失败")
				errChan <- err
			}
			memberContract, err := tx.MemberContract.Create().
				SetMember(members).
				SetOrderID(one.ID).
				SetContractID(contracts.ID).
				SetMemberProduct(mp).
				SetName(contracts.Name).
				SetSign(signImg).
				SetStatus(0).
				Save(o.ctx)
			if err != nil {
				err = errors.Wrap(err, "创建Member Contract "+string(i)+" 失败")
				errChan <- err
			}

			_, err = tx.MemberContractContent.Create().
				SetContract(memberContract).
				SetContent(contracts.Content).
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
		return "", rollback(tx, result)
	default:
	}

	if err = tx.Commit(); err != nil {
		return "", err
	}
	return one.OrderSn, nil

}
