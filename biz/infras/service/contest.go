package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/contest"
	"saas/pkg/db/ent"
	contest2 "saas/pkg/db/ent/contest"
	"saas/pkg/db/ent/contestparticipant"
	"saas/pkg/db/ent/predicate"
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

func (c Contest) CreateContest(req contest.ContestInfo) error {

	signStartAt, _ := time.Parse(time.DateOnly, req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateOnly, req.SignEndAt)
	startAt, _ := time.Parse(time.DateOnly, req.StartAt)
	endAt, _ := time.Parse(time.DateOnly, req.EndAt)

	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	cont := tx.Contest.Create().
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
		SetIsFee(req.IsFee)

	_, err = cont.Save(c.ctx)

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (c Contest) UpdateContest(req contest.ContestInfo) error {
	signStartAt, _ := time.Parse(time.DateOnly, req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateOnly, req.SignEndAt)
	startAt, _ := time.Parse(time.DateOnly, req.StartAt)
	endAt, _ := time.Parse(time.DateOnly, req.EndAt)

	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	_, err = tx.Contest.Update().
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

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (c Contest) DeleteContest(id int64) error {
	//TODO implement me
	panic("implement me")
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
		signStartAt, _ := time.Parse(time.DateOnly, req.SignStartAt)
		signEndAt, _ := time.Parse(time.DateOnly, req.SignEndAt)

		predicates = append(predicates, contest2.SignStartAtGT(signStartAt))
		predicates = append(predicates, contest2.SignEndAtEQ(signEndAt))
	}
	if req.StartAt != "" && req.EndAt != "" {
		startAt, _ := time.Parse(time.DateOnly, req.StartAt)
		endAt, _ := time.Parse(time.DateOnly, req.EndAt)

		predicates = append(predicates, contest2.StartAtGT(startAt))
		predicates = append(predicates, contest2.EndAtEQ(endAt))
	}

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
	signStartAt := v.SignStartAt.Format(time.DateTime)
	startAt := v.SignStartAt.Format(time.DateTime)
	signEndAt := v.SignStartAt.Format(time.DateTime)
	endAt := v.SignStartAt.Format(time.DateTime)
	pic := minio.URLconvert(c.ctx, c.c, v.Pic)
	return &contest.ContestInfo{
		Pic:         pic,
		ID:          v.ID,
		Name:        v.Name,
		Detail:      v.Detail,
		SignNumber:  v.SignNumber,
		SignStartAt: signStartAt,
		SignEndAt:   signEndAt,
		Number:      v.Number,
		StartAt:     startAt,
		EndAt:       endAt,
		Sponsor:     v.Sponsor,
		Fee:         v.Fee,
		IsCancel:    v.IsCancel,
		CancelTime:  v.CancelTime,
		SignFields:  v.SignFields,
		IsFee:       v.IsFee,
	}
}

func (c Contest) UpdateContestStatus(ID int64, status int64) error {
	_, err := c.db.Contest.Update().Where(contest2.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Contest) CreateParticipant(req contest.ParticipantInfo) error {

	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	cont := tx.ContestParticipant.Create().
		SetName(*req.Name).
		SetContestID(*req.ContestId).
		SetMobile(*req.Mobile).
		SetFields(*req.Fields)
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
		Where(contestparticipant.IDEQ(*req.ID)).
		SetContestID(*req.ContestId).
		SetName(*req.Name).
		SetMobile(*req.Mobile).
		SetFields(*req.Fields)
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
	participantEnt, err := c.db.ContestParticipant.Query().Where(contestparticipant.IDEQ(id)).First(c.ctx)

	resp.ID = &participantEnt.ID
	resp.Name = &participantEnt.Name
	resp.Mobile = &participantEnt.Mobile
	resp.Fields = &participantEnt.Fields
	resp.ContestId = &participantEnt.ContestID

	c.cache.SetWithTTL("participantInfo"+strconv.Itoa(int(*resp.ID)), &resp, 1, 1*time.Hour)
	return
}

func (c Contest) ParticipantList(req contest.ParticipantListReq) (resp []*contest.ParticipantInfo, total int, err error) {
	var predicates []predicate.ContestParticipant
	if req.Name != "" {
		predicates = append(predicates, contestparticipant.NameEQ(req.Name))
	}
	if req.Mobile != "" {
		predicates = append(predicates, contestparticipant.Mobile(req.Mobile))
	}

	predicates = append(predicates, contestparticipant.ContestID(req.ContestId))

	lists, err := c.db.ContestParticipant.Query().Where(predicates...).
		Order(ent.Desc(contestparticipant.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get contest participant list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, &contest.ParticipantInfo{
			ID:        &v.ID,
			Name:      &v.Name,
			Mobile:    &v.Mobile,
			Fields:    &v.Fields,
			ContestId: &v.ContestID,
		})
	}

	total, _ = c.db.ContestParticipant.Query().Where(predicates...).Count(c.ctx)
	return
}

func (c Contest) UpdateParticipantStatus(ID int64, status int64) error {
	_, err := c.db.ContestParticipant.Update().Where(contestparticipant.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
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
