package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
	"os"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	contest2 "saas/biz/dal/db/ent/contest"
	"saas/biz/dal/db/ent/contestparticipant"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/user"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/contest"
	"saas/pkg/consts"
	"saas/pkg/minio"
	"strconv"
	"time"
)

type Contest struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewContest(ctx context.Context, c *app.RequestContext) do.Contest {
	return &Contest{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
func (c Contest) CreateContest(req contest.ContestInfo) error {

	signStartAt, _ := time.Parse(time.DateTime, req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateTime, req.SignEndAt)
	startAt, _ := time.Parse(time.DateTime, req.StartAt)
	endAt, _ := time.Parse(time.DateTime, req.EndAt)

	var createdId int64
	userId, exist := c.c.Get("userId")
	if !exist || userId == nil {
		createdId = userId.(int64)
	}

	_, err := c.db.Contest.Create().
		SetName(req.Name).
		SetDetail(req.Detail).
		SetSignNumber(req.SignNumber).
		SetSignStartAt(signStartAt).
		SetSignEndAt(signEndAt).
		SetNumber(req.Number).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetPic(req.Pic).
		SetSponsor(req.Sponsor).
		SetFee(req.Fee).
		SetIsCancel(req.IsCancel).
		SetCancelTime(req.CancelTime).
		SetSignFields(req.SignFields).
		SetIsFee(req.IsFee).
		SetIsShow(req.IsShow).
		SetCreatedID(createdId).
		Save(c.ctx)
	if err != nil {
		err = errors.Wrap(err, err.Error())
		return err
	}

	return nil
}

func (c Contest) UpdateContest(req contest.ContestInfo) error {
	signStartAt, _ := time.Parse(time.DateTime, req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateTime, req.SignEndAt)
	startAt, _ := time.Parse(time.DateTime, req.StartAt)
	endAt, _ := time.Parse(time.DateTime, req.EndAt)

	_, err := c.db.Contest.Update().
		Where(contest2.IDEQ(req.ID)).
		SetName(req.Name).
		SetDetail(req.Detail).
		SetSignNumber(req.SignNumber).
		SetSignStartAt(signStartAt).
		SetSignEndAt(signEndAt).
		SetNumber(req.Number).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetPic(req.Pic).
		SetSponsor(req.Sponsor).
		SetFee(req.Fee).
		SetIsCancel(req.IsCancel).
		SetCancelTime(req.CancelTime).
		SetSignFields(req.SignFields).
		SetIsFee(req.IsFee).
		Save(c.ctx)

	if err != nil {
		err = errors.Wrap(err, err.Error())
		return err
	}

	return nil
}

func (c Contest) DeleteContest(id int64) error {
	_, err := c.db.Contest.Update().Where(contest2.IDEQ(id)).SetDelete(1).Save(c.ctx)
	return err
}

func (c Contest) ContestList(req contest.ContestListReq) (resp []*contest.ContestInfo, total int, err error) {
	var predicates []predicate.Contest
	if req.Name != "" {
		predicates = append(predicates, contest2.NameEQ(req.Name))
	}
	if req.Condition != 0 {
		predicates = append(predicates, contest2.Condition(req.Condition))
	}

	if req.SignStartAt != "" && req.SignEndAt != "" {
		signStartAt, _ := time.Parse(time.DateTime, req.SignStartAt)
		signEndAt, _ := time.Parse(time.DateTime, req.SignEndAt)

		predicates = append(predicates, contest2.SignStartAtGT(signStartAt))
		predicates = append(predicates, contest2.SignEndAtEQ(signEndAt))
	}
	if req.StartAt != "" && req.EndAt != "" {
		startAt, _ := time.Parse(time.DateTime, req.StartAt)
		endAt, _ := time.Parse(time.DateTime, req.EndAt)

		predicates = append(predicates, contest2.StartAtGT(startAt))
		predicates = append(predicates, contest2.EndAtEQ(endAt))
	}
	predicates = append(predicates, contest2.Delete(0))
	lists, err := c.db.Contest.Query().Where(predicates...).
		Order(ent.Desc(contest2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, c.entContestInfo(v))
	}

	total, _ = c.db.Contest.Query().Where(predicates...).Count(c.ctx)
	return
}

func (c Contest) ContestInfo(id int64) (resp *contest.ContestInfo, err error) {

	resp = new(contest.ContestInfo)

	contestInter, exist := c.cache.Get("contestInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := contestInter.(*contest.ContestInfo); ok {
			return u, nil
		}
	}
	contestEnt, err := c.db.Contest.Query().Where(contest2.IDEQ(id)).First(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get contest failed")
		return resp, err
	}
	resp = c.entContestInfo(contestEnt)
	c.cache.SetWithTTL("contestInfo"+strconv.Itoa(int(resp.ID)), &resp, 1, 1*time.Hour)

	return

}

func (c Contest) entContestInfo(v *ent.Contest) *contest.ContestInfo {
	signStartAt := v.SignStartAt.Unix()
	startAt := v.SignStartAt.Unix()
	signEndAt := v.SignStartAt.Unix()
	endAt := v.SignStartAt.Unix()
	createdAt := v.CreatedAt.Unix()
	pic := minio.URLconvert(c.ctx, c.c, v.Pic)
	var createdName string
	first, _ := c.db.User.Query().Where(user.IDEQ(v.CreatedID)).First(c.ctx)
	if first != nil {
		createdName = first.Name
	}

	return &contest.ContestInfo{
		Pic:         pic,
		ID:          v.ID,
		Name:        v.Name,
		Detail:      v.Detail,
		SignNumber:  v.SignNumber,
		SignStartAt: strconv.FormatInt(signStartAt, 10),
		SignEndAt:   strconv.FormatInt(signEndAt, 10),
		Number:      v.Number,
		StartAt:     strconv.FormatInt(startAt, 10),
		EndAt:       strconv.FormatInt(endAt, 10),
		Sponsor:     v.Sponsor,
		Fee:         v.Fee,
		IsCancel:    v.IsCancel,
		CancelTime:  v.CancelTime,
		SignFields:  v.SignFields,
		IsFee:       v.IsFee,
		CreatedAt:   strconv.FormatInt(createdAt, 10),
		Condition:   v.Condition,
		CreatedId:   v.CreatedID,
		CreatedName: createdName,
		IsShow:      v.IsShow,
	}
}

func (c Contest) UpdateContestStatus(ID int64, status int64) error {
	_, err := c.db.Contest.Update().Where(contest2.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Contest) UpdateContestShow(ID int64, status int64) error {
	_, err := c.db.Contest.Update().Where(contest2.IDEQ(ID)).SetIsShow(status).Save(c.ctx)
	return err
}

func (c Contest) DelContest(ID int64) error {
	_, err := c.db.Contest.Update().
		Where(contest2.IDEQ(ID)).SetDelete(1).Save(c.ctx)

	return err
}

func (c Contest) CreateParticipant(req contest.ParticipantInfo) error {

	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	cont := tx.ContestParticipant.Create().
		SetName(req.Name).
		SetContestID(req.ContestId).
		SetMobile(req.Mobile).
		SetFields(req.Fields)
	_, err = cont.Save(c.ctx)
	if err = tx.Commit(); err != nil {
		return err
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
	//TODO implement me
	panic("implement me")
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

func (c Contest) ParticipantListListExport(req contest.ParticipantListReq) (string, error) {

	exportFilePath := consts.ExportFilePath + time.Now().Format(time.DateOnly) + "/"
	if err := os.MkdirAll(exportFilePath, 0o777); err != nil {
		panic(err)
	}

	contest, _ := c.db.Contest.Query().Where(contest2.IDEQ(req.ContestId)).First(c.ctx)

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

		var status string
		switch row.Status {
		case 1:
			status = "报名中"
		case 2:
			status = "报名完成"
		case 3:
			status = "取消报名"
		case 4:
			status = "报名失效"
		case 5:
			status = "报名成功"
		default:
			status = "状态异常"
		}

		r := []interface{}{row.Name, row.Mobile, row.Fee, status, createdAtr}
		err = f.SetSheetRow("Sheet1", cell, &r)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}
	ing := strconv.FormatInt(time.Now().Unix(), 10)
	files := exportFilePath + contest.Name + "-参赛人导出" + ing + ".xlsx"
	//Save spreadsheet by the given path.
	if err := f.SaveAs(files); err != nil {
		fmt.Println(err)
	}
	return files, nil

}
