package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	contest2 "saas/biz/dal/db/ent/contest"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/minio"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/contest"
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
	if exist || userId != nil {
		uId, ok := userId.(string)
		if ok {
			createdId, _ = strconv.ParseInt(uId, 10, 64)
		}
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

		predicates = append(predicates, contest2.SignStartAtLTE(signStartAt))
		predicates = append(predicates, contest2.SignEndAtGTE(signEndAt))
	}
	if req.StartAt != "" && req.EndAt != "" {
		startAt, _ := time.Parse(time.DateTime, req.StartAt)
		endAt, _ := time.Parse(time.DateTime, req.EndAt)

		predicates = append(predicates, contest2.StartAtLTE(startAt))
		predicates = append(predicates, contest2.EndAtGTE(endAt))
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
	_, err := c.db.Contest.Update().Where(contest2.IDEQ(ID)).SetDelete(1).Save(c.ctx)
	return err
}
