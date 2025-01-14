package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"saas/biz/dal/db/ent"
	bootcamp2 "saas/biz/dal/db/ent/bootcamp"
	"saas/biz/dal/db/ent/bootcampparticipant"
	member2 "saas/biz/dal/db/ent/member"
	order2 "saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/infras/do"
	"saas/idl_gen/model/bootcamp"
	"saas/pkg/enums"
	"saas/pkg/utils"
	"strconv"
	"time"
)

func (b Bootcamp) CreateParticipantMember(req bootcamp.ParticipantInfo) (member *ent.Member, err error) {
	member, err = b.db.Member.Query().Where(member2.Mobile(req.Mobile)).First(b.ctx)

	if member == nil {
		member, err = b.db.Member.Create().SetMobile(req.Mobile).Save(b.ctx)
		//.SetName(req.Name)
		hlog.Info(err)
		if err != nil {
			return nil, err
		}
		_, err = b.db.MemberDetails.Create().SetMemberID(member.ID).Save(b.ctx)
	} else {

		first, _ := b.db.BootcampParticipant.Query().Where(
			bootcampparticipant.MemberID(member.ID),
			bootcampparticipant.BootcampID(req.BootcampId),
		).Exist(b.ctx)
		if first {
			return nil, errors.New("已报名")
		}
	}

	return member, nil
}
func (b Bootcamp) CreateParticipant(req bootcamp.ParticipantInfo) error {
	tx, err := b.db.Tx(b.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	member, err := b.CreateParticipantMember(req)

	if err != nil {
		return err
	}

	cont := tx.BootcampParticipant.Create().
		SetName(req.Name).
		SetBootcampID(req.BootcampId).
		SetMobile(req.Mobile).
		SetFields(req.Fields).
		SetMemberID(member.ID)
	_, err = cont.Save(b.ctx)

	if err = tx.Commit(); err != nil {
		return err
	}

	order, err := NewOrder(b.ctx, b.c).CreateParticipantOrder(do.CreateParticipantOrderReq{
		Member: member,
		Device: req.Device,
		NodeId: req.BootcampId,
		Fee:    0,
	})
	if err != nil {
		return err
	}
	if req.Device == "pc" {
		b.db.OrderPay.Create().
			SetOrderID(order.ID).
			SetPayWay("pc").
			SetRemission(0).
			SetPaySn(utils.CreateCn()).
			SetPayAt(time.Now()).
			SetPay(0).
			Save(b.ctx)
		b.db.Order.Update().
			Where(order2.IDEQ(order.ID)).
			SetStatus(5).
			SetCompletionAt(time.Now()).
			Save(b.ctx)
	}

	return nil
}

func (b Bootcamp) UpdateParticipant(req bootcamp.ParticipantInfo) error {
	tx, err := b.db.Tx(b.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	cont := tx.BootcampParticipant.Update().
		Where(bootcampparticipant.IDEQ(req.ID)).
		SetBootcampID(req.BootcampId).
		SetName(req.Name).
		SetMobile(req.Mobile).
		SetFields(req.Fields)
	_, err = cont.Save(b.ctx)
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (b Bootcamp) DeleteParticipant(id int64) error {
	_, err := b.db.BootcampParticipant.Update().Where(bootcampparticipant.IDEQ(id)).SetDelete(1).Save(b.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b Bootcamp) ParticipantInfo(id int64) (resp *bootcamp.ParticipantInfo, err error) {
	resp = new(bootcamp.ParticipantInfo)

	inter, exist := b.cache.Get("participantInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := inter.(*bootcamp.ParticipantInfo); ok {
			return u, nil
		}
	}
	participantEnt, _ := b.db.BootcampParticipant.Query().Where(bootcampparticipant.IDEQ(id)).First(b.ctx)

	resp = b.entParticipantInfo(participantEnt)

	b.cache.SetWithTTL("participantInfo"+strconv.Itoa(int(resp.ID)), &resp, 1, 1*time.Hour)
	return
}
func (b Bootcamp) entParticipantInfo(v *ent.BootcampParticipant) *bootcamp.ParticipantInfo {

	createdAt := v.CreatedAt.Unix()
	return &bootcamp.ParticipantInfo{
		ID:         v.ID,
		Name:       v.Name,
		Mobile:     v.Mobile,
		Fields:     v.Fields,
		BootcampId: v.BootcampID,
		CreatedAt:  strconv.FormatInt(createdAt, 10),
		OrderSn:    v.OrderSn,
		Fee:        v.Fee,
		Status:     v.Status,
	}
}
func (b Bootcamp) ParticipantList(req bootcamp.ParticipantListReq) (resp []*bootcamp.ParticipantInfo, total int, err error) {
	var predicates []predicate.BootcampParticipant
	if req.Name != "" {
		predicates = append(predicates, bootcampparticipant.NameEQ(req.Name))
	}
	if req.Mobile != "" {
		predicates = append(predicates, bootcampparticipant.Mobile(req.Mobile))
	}
	if req.Sn != "" {
		//predicates = append(predicates, bootcampparticipant.Mobile(req.Sn))
	}
	predicates = append(predicates, bootcampparticipant.BootcampID(req.BootcampId))
	predicates = append(predicates, bootcampparticipant.Delete(0))
	lists, err := b.db.BootcampParticipant.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(bootcampparticipant.FieldID)).
		Limit(int(req.PageSize)).All(b.ctx)
	if err != nil {
		err = errors.Wrap(err, "get bootcamp participant list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, b.entParticipantInfo(v))
	}

	total, _ = b.db.BootcampParticipant.Query().Where(predicates...).Count(b.ctx)
	return
}

func (b Bootcamp) UpdateParticipantStatus(ID int64, status int64) error {
	_, err := b.db.BootcampParticipant.Update().Where(bootcampparticipant.IDEQ(ID)).SetStatus(status).Save(b.ctx)
	return err
}

func (b Bootcamp) ParticipantListExport(req bootcamp.ParticipantListReq) (string, error) {
	bootcampe, _ := b.db.Bootcamp.Query().Where(bootcamp2.IDEQ(req.BootcampId)).First(b.ctx)
	exportFilePath, domain := utils.ExportFilePath(bootcampe.Name + "参赛人导出")
	resp, total, _ := b.ParticipantList(req)

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
	tale := []interface{}{"名称", "手机号", "费用", "状态", "报名时间"}
	//row = []interface{}{resp[idx].Name, resp[idx].Mobile, resp[idx].Fee, resp[idx].Status, resp[idx].CreatedAt}
	err = f.SetSheetRow("Sheet1", cell, &tale)
	if err != nil {
		return "", err
	}

	for idx, row := range resp {
		cell, err := excelize.CoordinatesToCellName(1, idx+2)
		if err != nil {
			return "", err
		}
		var createdAt int64
		createdAt, _ = strconv.ParseInt(row.CreatedAt, 10, 64)
		createdAtr := time.Unix(createdAt, 0).Format(time.DateTime)

		status := enums.ReturnBootcampParticipantStatusValues(row.Status)

		r := []interface{}{row.Name, row.Mobile, row.Fee, status, createdAtr}
		err = f.SetSheetRow("Sheet1", cell, &r)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	//Save spreadsheet by the given path.
	if err := f.SaveAs(exportFilePath); err != nil {
		fmt.Println(err)
	}
	return domain, nil
}
