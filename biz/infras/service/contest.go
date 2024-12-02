package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/contest"
	"saas/pkg/db/ent"
	contest2 "saas/pkg/db/ent/contest"
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

	signStartAt, _ := time.Parse(time.DateOnly, *req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateOnly, *req.SignEndAt)
	startAt, _ := time.Parse(time.DateOnly, *req.StartAt)
	endAt, _ := time.Parse(time.DateOnly, *req.EndAt)

	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	cont := tx.Contest.Create().
		SetName(*req.Name).
		SetDetail(*req.Detail).
		SetSignNumber(*req.SignNumber).
		SetSignStartAt(signStartAt).
		SetSignEndAt(signEndAt).
		SetNumber(*req.Number).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetPic(*req.Pic).
		SetSponsor(*req.Sponsor).
		SetFee(*req.Fee).
		SetIsCancel(*req.IsCancel).
		SetCancelTime(*req.CancelTime).
		SetSignFields(*req.SignFields)

	_, err = cont.Save(c.ctx)

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (c Contest) UpdateContest(req contest.ContestInfo) error {
	signStartAt, _ := time.Parse(time.DateOnly, *req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateOnly, *req.SignEndAt)
	startAt, _ := time.Parse(time.DateOnly, *req.StartAt)
	endAt, _ := time.Parse(time.DateOnly, *req.EndAt)

	tx, err := c.db.Tx(c.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	_, err = tx.Contest.Update().
		Where(contest2.IDEQ(*req.ID)).
		SetName(*req.Name).
		SetDetail(*req.Detail).
		SetSignNumber(*req.SignNumber).
		SetSignStartAt(signStartAt).
		SetSignEndAt(signEndAt).
		SetNumber(*req.Number).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetPic(*req.Pic).
		SetSponsor(*req.Sponsor).
		SetFee(*req.Fee).
		SetIsCancel(*req.IsCancel).
		SetCancelTime(*req.CancelTime).
		SetSignFields(*req.SignFields).
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
	if *req.Name != "" {
		predicates = append(predicates, contest2.NameEQ(*req.Name))
	}
	if *req.Condition != 0 {
		predicates = append(predicates, contest2.Condition(*req.Condition))
	}

	if *req.SignStartAt != "" && *req.SignEndAt != "" {
		signStartAt, _ := time.Parse(time.DateOnly, *req.SignStartAt)
		signEndAt, _ := time.Parse(time.DateOnly, *req.SignEndAt)

		predicates = append(predicates, contest2.SignStartAtGT(signStartAt))
		predicates = append(predicates, contest2.SignEndAtEQ(signEndAt))
	}
	if *req.StartAt != "" && *req.EndAt != "" {
		startAt, _ := time.Parse(time.DateOnly, *req.StartAt)
		endAt, _ := time.Parse(time.DateOnly, *req.EndAt)

		predicates = append(predicates, contest2.StartAtGT(startAt))
		predicates = append(predicates, contest2.EndAtEQ(endAt))
	}

	lists, err := c.db.Contest.Query().Where(predicates...).
		Order(ent.Desc(contest2.FieldID)).
		Offset(int(*req.Page-1) * int(*req.PageSize)).
		Limit(int(*req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Member info failed")
		return resp, 0, err
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
	signStartAt := contestEnt.SignStartAt.Format(time.DateTime)
	startAt := contestEnt.SignStartAt.Format(time.DateTime)
	signEndAt := contestEnt.SignStartAt.Format(time.DateTime)
	endAt := contestEnt.SignStartAt.Format(time.DateTime)

	pic := minio.URLconvert(c.ctx, c.c, contestEnt.Pic)
	resp.Pic = &pic
	resp.ID = &contestEnt.ID
	resp.Name = &contestEnt.Name
	resp.Detail = &contestEnt.Detail
	resp.SignNumber = &contestEnt.SignNumber
	resp.SignStartAt = &signStartAt
	resp.SignEndAt = &signEndAt
	resp.Number = &contestEnt.Number
	resp.StartAt = &startAt
	resp.EndAt = &endAt
	resp.Sponsor = &contestEnt.Sponsor
	resp.Fee = &contestEnt.Fee
	resp.IsCancel = &contestEnt.IsCancel
	resp.CancelTime = &contestEnt.CancelTime
	resp.SignFields = &contestEnt.SignFields

	c.cache.SetWithTTL("contestInfo"+strconv.Itoa(int(*resp.ID)), &resp, 1, 1*time.Hour)

	return

}

func (c Contest) UpdateContestStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) CreateParticipant(req contest.ParticipantInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) UpdateParticipant(req contest.ParticipantInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) DeleteParticipant(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Contest) ParticipantInfo(id int64) (resp *contest.ParticipantInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Contest) ParticipantList(req contest.ParticipantListReq) (resp []*contest.ParticipantInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Contest) UpdateParticipantStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
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
