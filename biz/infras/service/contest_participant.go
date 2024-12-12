package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"saas/biz/dal/db/ent"
	contest2 "saas/biz/dal/db/ent/contest"
	"saas/biz/dal/db/ent/contestparticipant"
	member2 "saas/biz/dal/db/ent/member"
	order2 "saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/infras/do"
	"saas/idl_gen/model/contest"
	"saas/pkg/enums"
	"saas/pkg/utils"
	"strconv"
	"time"
)

func (c Contest) CreateParticipantMember(req contest.ParticipantInfo) (member *ent.Member, err error) {
	member, err = c.db.Member.Query().Where(member2.Mobile(req.Mobile)).First(c.ctx)

	if member == nil {
		member, err = c.db.Member.Create().SetName(req.Name).SetMobile(req.Mobile).Save(c.ctx)
		hlog.Info(err)
		if err != nil {
			return nil, err
		}
		_, err = c.db.MemberDetails.Create().SetMemberID(member.ID).Save(c.ctx)
	} else {

		first, _ := c.db.ContestParticipant.Query().Where(
			contestparticipant.MemberID(member.ID),
			contestparticipant.ContestID(req.ContestId),
		).Exist(c.ctx)
		if first {
			return nil, errors.New("已报名")
		}
	}

	return member, nil
}
func (c Contest) CreateParticipant(req contest.ParticipantInfo) error {

	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	member, err := c.CreateParticipantMember(req)

	if err != nil {
		return err
	}

	cont := tx.ContestParticipant.Create().
		SetName(req.Name).
		SetContestID(req.ContestId).
		SetMobile(req.Mobile).
		SetFields(req.Fields).
		SetMemberID(member.ID)
	_, err = cont.Save(c.ctx)

	if err = tx.Commit(); err != nil {
		return err
	}

	order, err := NewOrder(c.ctx, c.c).CreateParticipantOrder(do.CreateParticipantOrderReq{
		Member: member,
		Device: req.Device,
		NodeId: req.ContestId,
		Fee:    0,
	})
	if err != nil {
		return err
	}
	if req.Device == "pc" {
		c.db.OrderPay.Create().
			SetOrderID(order.ID).
			SetPayWay("pc").
			SetRemission(0).
			SetPaySn(utils.CreateCn()).
			SetPayAt(time.Now()).
			SetPay(0).
			Save(c.ctx)
		c.db.Order.Update().
			Where(order2.IDEQ(order.ID)).
			SetStatus(5).
			SetCompletionAt(time.Now()).
			Save(c.ctx)
	}

	return nil
}

func (c Contest) UpdateParticipant(req contest.ParticipantInfo) error {
	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	cont := tx.ContestParticipant.Update().
		Where(contestparticipant.IDEQ(req.ID)).
		SetContestID(req.ContestId).
		SetName(req.Name).
		SetMobile(req.Mobile).
		SetFields(req.Fields)
	_, err = cont.Save(c.ctx)
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (c Contest) DeleteParticipant(id int64) error {
	_, err := c.db.ContestParticipant.Update().Where(contestparticipant.IDEQ(id)).SetDelete(1).Save(c.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c Contest) ParticipantInfo(id int64) (resp *contest.ParticipantInfo, err error) {
	resp = new(contest.ParticipantInfo)

	inter, exist := c.cache.Get("participantInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := inter.(*contest.ParticipantInfo); ok {
			return u, nil
		}
	}
	participantEnt, _ := c.db.ContestParticipant.Query().Where(contestparticipant.IDEQ(id)).First(c.ctx)

	resp = c.entParticipantInfo(participantEnt)

	c.cache.SetWithTTL("participantInfo"+strconv.Itoa(int(resp.ID)), &resp, 1, 1*time.Hour)
	return
}

func (c Contest) entParticipantInfo(v *ent.ContestParticipant) *contest.ParticipantInfo {

	createdAt := v.CreatedAt.Unix()
	return &contest.ParticipantInfo{
		ID:        v.ID,
		Name:      v.Name,
		Mobile:    v.Mobile,
		Fields:    v.Fields,
		ContestId: v.ContestID,
		CreatedAt: strconv.FormatInt(createdAt, 10),
		OrderSn:   v.OrderSn,
		Fee:       v.Fee,
		Status:    v.Status,
	}
}

func (c Contest) ParticipantList(req contest.ParticipantListReq) (resp []*contest.ParticipantInfo, total int, err error) {
	var predicates []predicate.ContestParticipant
	if req.Name != "" {
		predicates = append(predicates, contestparticipant.NameEQ(req.Name))
	}
	if req.Mobile != "" {
		predicates = append(predicates, contestparticipant.Mobile(req.Mobile))
	}
	if req.Sn != "" {
		//predicates = append(predicates, contestparticipant.Mobile(req.Sn))
	}
	predicates = append(predicates, contestparticipant.ContestID(req.ContestId))
	predicates = append(predicates, contestparticipant.Delete(0))
	lists, err := c.db.ContestParticipant.Query().Where(predicates...).
		Order(ent.Desc(contestparticipant.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get contest participant list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, c.entParticipantInfo(v))
	}

	total, _ = c.db.ContestParticipant.Query().Where(predicates...).Count(c.ctx)
	return
}

func (c Contest) UpdateParticipantStatus(ID int64, status int64) error {
	_, err := c.db.ContestParticipant.Update().Where(contestparticipant.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Contest) ParticipantListExport(req contest.ParticipantListReq) (string, error) {

	conteste, _ := c.db.Contest.Query().Where(contest2.IDEQ(req.ContestId)).First(c.ctx)
	exportFilePath, domain := utils.ExportFilePath(conteste.Name + "参赛人导出")
	resp, total, _ := c.ParticipantList(req)

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

		status := enums.ReturnContestParticipantStatusValues(row.Status)

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
