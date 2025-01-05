package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"

	"saas/biz/dal/db/ent"
	community2 "saas/biz/dal/db/ent/community"
	"saas/biz/dal/db/ent/communityparticipant"
	member2 "saas/biz/dal/db/ent/member"
	order2 "saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/infras/do"
	"saas/idl_gen/model/community"
	"saas/pkg/enums"
	"saas/pkg/utils"
	"strconv"
	"time"
)

func (c Community) CreateParticipantMember(req community.ParticipantInfo) (member *ent.Member, err error) {
	member, err = c.db.Member.Query().Where(member2.Mobile(req.Mobile)).First(c.ctx)

	if member == nil {
		member, err = c.db.Member.Create().SetMobile(req.Mobile).Save(c.ctx)
		//SetName(req.Name).
		hlog.Info(err)
		if err != nil {
			return nil, err
		}
		_, err = c.db.MemberDetails.Create().SetMemberID(member.ID).Save(c.ctx)
	} else {

		first, _ := c.db.CommunityParticipant.Query().Where(
			communityparticipant.MemberID(member.ID),
			communityparticipant.CommunityID(req.CommunityId),
		).Exist(c.ctx)
		if first {
			return nil, errors.New("已报名")
		}
	}

	return member, nil
}
func (c Community) CreateParticipant(req community.ParticipantInfo) error {
	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	member, err := c.CreateParticipantMember(req)

	if err != nil {
		return err
	}

	cont := tx.CommunityParticipant.Create().
		SetName(req.Name).
		SetCommunityID(req.CommunityId).
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
		NodeId: req.CommunityId,
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

func (c Community) UpdateParticipant(req community.ParticipantInfo) error {
	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	cont := tx.CommunityParticipant.Update().
		Where(communityparticipant.IDEQ(req.ID)).
		SetCommunityID(req.CommunityId).
		SetName(req.Name).
		SetMobile(req.Mobile).
		SetFields(req.Fields)
	_, err = cont.Save(c.ctx)
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (c Community) DeleteParticipant(id int64) error {
	_, err := c.db.CommunityParticipant.Update().Where(communityparticipant.IDEQ(id)).SetDelete(1).Save(c.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c Community) ParticipantInfo(id int64) (resp *community.ParticipantInfo, err error) {
	resp = new(community.ParticipantInfo)

	inter, exist := c.cache.Get("participantInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := inter.(*community.ParticipantInfo); ok {
			return u, nil
		}
	}
	participantEnt, _ := c.db.CommunityParticipant.Query().Where(communityparticipant.IDEQ(id)).First(c.ctx)

	resp = c.entParticipantInfo(participantEnt)

	c.cache.SetWithTTL("participantInfo"+strconv.Itoa(int(resp.ID)), &resp, 1, 1*time.Hour)
	return
}
func (c Community) entParticipantInfo(v *ent.CommunityParticipant) *community.ParticipantInfo {

	createdAt := v.CreatedAt.Unix()
	return &community.ParticipantInfo{
		ID:          v.ID,
		Name:        v.Name,
		Mobile:      v.Mobile,
		Fields:      v.Fields,
		CommunityId: v.CommunityID,
		CreatedAt:   strconv.FormatInt(createdAt, 10),
		OrderSn:     v.OrderSn,

		Status: v.Status,
	}
}
func (c Community) ParticipantList(req community.ParticipantListReq) (resp []*community.ParticipantInfo, total int, err error) {
	var predicates []predicate.CommunityParticipant
	if req.Name != "" {
		predicates = append(predicates, communityparticipant.NameEQ(req.Name))
	}
	if req.Mobile != "" {
		predicates = append(predicates, communityparticipant.Mobile(req.Mobile))
	}
	if req.Sn != "" {
		//predicates = append(predicates, communityparticipant.Mobile(req.Sn))
	}
	predicates = append(predicates, communityparticipant.CommunityID(req.CommunityId))
	predicates = append(predicates, communityparticipant.Delete(0))
	lists, err := c.db.CommunityParticipant.Query().Where(predicates...).
		Order(ent.Desc(communityparticipant.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get community participant list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, c.entParticipantInfo(v))
	}

	total, _ = c.db.CommunityParticipant.Query().Where(predicates...).Count(c.ctx)
	return
}

func (c Community) UpdateParticipantStatus(ID int64, status int64) error {
	_, err := c.db.CommunityParticipant.Update().Where(communityparticipant.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Community) ParticipantListExport(req community.ParticipantListReq) (string, error) {
	communitye, _ := c.db.Community.Query().Where(community2.IDEQ(req.CommunityId)).First(c.ctx)
	exportFilePath, domain := utils.ExportFilePath(communitye.Name + "参赛人导出")
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
	tale := []interface{}{"名称", "手机号", "状态", "报名时间"}
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

		status := enums.ReturnCommunityParticipantStatusValues(row.Status)

		r := []interface{}{row.Name, row.Mobile, status, createdAtr}
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
