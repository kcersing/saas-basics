package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"os"
	"saas/biz/dal/db/ent/member"
	order2 "saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/orderitem"
	"saas/biz/dal/db/ent/predicate"
	"saas/idl_gen/model/order"
	"saas/pkg/consts"
	"saas/pkg/enums"
	"strconv"
	"time"
)

func (o Order) List(req *order.ListOrderReq) (resp []*order.OrderInfo, total int, err error) {
	var predicates []predicate.Order
	if req.OrderSn != "" {
		predicates = append(predicates, order2.OrderSnEQ(req.OrderSn))
	}
	if req.Mobile != "" {
		predicates = append(predicates, order2.HasOrderMembersWith(member.MobileEQ(req.Mobile)))
	}
	if req.MemberName != "" {
		predicates = append(predicates, order2.HasOrderMembersWith(member.NameEQ(req.MemberName)))
	}
	if req.Name != "" {
		predicates = append(predicates, order2.HasItemWith(orderitem.NameEQ(req.Name)))
	}
	if len(req.Status) > 0 {
		predicates = append(predicates, order2.StatusIn(req.Status...))
	}
	if req.ProductType != "" {
		predicates = append(predicates, order2.ProductTypeEQ(req.ProductType))
	}
	if req.Nature != "" {
		predicates = append(predicates, order2.NatureEQ(req.Nature))
	}
	if len(req.VenueId) > 0 {
		predicates = append(predicates, order2.VenueIDIn(req.VenueId...))
	}

	if req.StartCompletionAt != "" && req.EndCompletionAt != "" {
		signStartAt, _ := time.Parse(time.DateTime, req.StartCompletionAt)
		signEndAt, _ := time.Parse(time.DateTime, req.EndCompletionAt)

		predicates = append(predicates, order2.CompletionAtGTE(signStartAt))
		predicates = append(predicates, order2.CompletionAtLTE(signEndAt))
	}

	//lt：less than 小于
	//le：less than or equal to 小于等于
	//eq：equal to 等于
	//ne：not equal to 不等于
	//ge：greater than or equal to 大于等于
	//gt：greater than 大于

	lists, err := o.db.Order.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Order list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, o.entOrderInfo(v))
	}

	total, _ = o.db.Order.Query().Where(predicates...).Count(o.ctx)
	return

}

func (o Order) OrderListExport(req *order.ListOrderReq) (string, error) {
	exportFilePath := consts.ExportFilePath + time.Now().Format(time.DateOnly) + "/"
	if err := os.MkdirAll(exportFilePath, 0o777); err != nil {
		panic(err)
	}

	resp, total, _ := o.List(req)

	if total == 0 {
		return "", errors.New("暂无数据")
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return "", err
	}
	tale := []interface{}{
		"订单编号",
		"姓名",
		"手机号",
		"名称",
		"报名费",
		"实付金额",
		"订单状态",
		"下单时间",
		"支付时间",
		"支付方式",
	}

	err = f.SetSheetRow("Sheet1", cell, &tale)
	if err != nil {
		return "", err
	}

	for idx, row := range resp {
		cell, err := excelize.CoordinatesToCellName(1, idx+2)
		if err != nil {
			return "", err
		}

		var payWay, PayAt string
		if len(row.OrderPay) > 0 {
			payWay = enums.ReturnOrderPayWayValues(row.OrderPay[0].PayWay)
			PayAt = row.OrderPay[0].PayAt
		}

		status := enums.ReturnOrderStatusValues(row.Status)

		r := []interface{}{
			row.OrderSn,
			row.MemberName,
			row.MemberMobile,
			row.OrderItem.Name,
			row.OrderAmount.Total,
			row.OrderAmount.Actual,
			status,
			row.CompletionAt,
			PayAt,
			payWay,
		}

		err = f.SetSheetRow("Sheet1", cell, &r)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}
	ing := strconv.FormatInt(time.Now().Unix(), 10)
	files := exportFilePath + "订单列表导出" + ing + ".xlsx"
	//Save spreadsheet by the given path.
	if err := f.SaveAs(files); err != nil {
		fmt.Println(err)
	}
	return files, nil
}
