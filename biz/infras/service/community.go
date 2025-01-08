package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/db"

	"saas/biz/dal/cache"

	"saas/biz/dal/db/ent"
	community2 "saas/biz/dal/db/ent/community"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/minio"

	"saas/biz/dal/db/ent/predicate"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/community"
	"strconv"
	"time"
)

type Community struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (c Community) CreateCommunity(req community.CommunityInfo) error {

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
	_, err := c.db.Community.Create().
		SetName(req.Name).
		SetDetail(req.Detail).
		SetSignNumber(req.SignNumber).
		SetSignStartAt(signStartAt).
		SetSignEndAt(signEndAt).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetPic(req.Pic).
		SetIsCancel(req.IsCancel).
		SetCancelTime(req.CancelTime).
		SetSignFields(req.SignFields).
		SetIsShow(req.IsShow).
		SetCreatedID(createdId).
		Save(c.ctx)
	if err != nil {
		err = errors.Wrap(err, err.Error())
		return err
	}

	return nil
}

func (c Community) UpdateCommunity(req community.CommunityInfo) error {
	signStartAt, _ := time.Parse(time.DateTime, req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateTime, req.SignEndAt)
	startAt, _ := time.Parse(time.DateTime, req.StartAt)
	endAt, _ := time.Parse(time.DateTime, req.EndAt)

	_, err := c.db.Community.Update().
		Where(community2.IDEQ(req.ID)).
		SetName(req.Name).
		SetDetail(req.Detail).
		SetSignNumber(req.SignNumber).
		SetSignStartAt(signStartAt).
		SetSignEndAt(signEndAt).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetPic(req.Pic).
		SetIsCancel(req.IsCancel).
		SetCancelTime(req.CancelTime).
		SetSignFields(req.SignFields).
		Save(c.ctx)

	if err != nil {
		err = errors.Wrap(err, err.Error())
		return err
	}

	return nil
}

func (c Community) DeleteCommunity(id int64) error {
	_, err := c.db.Community.Update().Where(community2.IDEQ(id)).SetDelete(1).Save(c.ctx)
	return err
}

func (c Community) CommunityList(req community.CommunityListReq) (resp []*community.CommunityInfo, total int, err error) {
	var predicates []predicate.Community
	if req.Name != "" {
		predicates = append(predicates, community2.NameEQ(req.Name))
	}
	if req.Condition != 0 {
		predicates = append(predicates, community2.Condition(req.Condition))
	}

	if req.SignStartAt != "" && req.SignEndAt != "" {
		signStartAt, _ := time.Parse(time.DateTime, req.SignStartAt)
		signEndAt, _ := time.Parse(time.DateTime, req.SignEndAt)

		predicates = append(predicates, community2.SignStartAtLTE(signStartAt))
		predicates = append(predicates, community2.SignEndAtGTE(signEndAt))
	}
	if req.StartAt != "" && req.EndAt != "" {
		startAt, _ := time.Parse(time.DateTime, req.StartAt)
		endAt, _ := time.Parse(time.DateTime, req.EndAt)

		predicates = append(predicates, community2.StartAtLTE(startAt))
		predicates = append(predicates, community2.EndAtLTE(endAt))
	}
	predicates = append(predicates, community2.Delete(0))
	lists, err := c.db.Community.Query().Where(predicates...).
		Order(ent.Desc(community2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, c.entCommunityInfo(v))
	}

	total, _ = c.db.Community.Query().Where(predicates...).Count(c.ctx)
	return
}

func (c Community) CommunityInfo(id int64) (resp *community.CommunityInfo, err error) {
	resp = new(community.CommunityInfo)

	communityInter, exist := c.cache.Get("communityInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := communityInter.(*community.CommunityInfo); ok {
			return u, nil
		}
	}
	communityEnt, err := c.db.Community.Query().Where(community2.IDEQ(id)).First(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Community failed")
		return resp, err
	}
	resp = c.entCommunityInfo(communityEnt)
	c.cache.SetWithTTL("communityInfo"+strconv.Itoa(int(resp.ID)), &resp, 1, 1*time.Hour)

	return
}

func (c Community) entCommunityInfo(v *ent.Community) *community.CommunityInfo {
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

	return &community.CommunityInfo{
		Pic:         pic,
		ID:          v.ID,
		Name:        v.Name,
		Detail:      v.Detail,
		SignNumber:  v.SignNumber,
		SignStartAt: strconv.FormatInt(signStartAt, 10),
		SignEndAt:   strconv.FormatInt(signEndAt, 10),
		StartAt:     strconv.FormatInt(startAt, 10),
		EndAt:       strconv.FormatInt(endAt, 10),
		IsCancel:    v.IsCancel,
		CancelTime:  v.CancelTime,
		SignFields:  v.SignFields,
		CreatedAt:   strconv.FormatInt(createdAt, 10),
		Condition:   v.Condition,
		CreatedId:   v.CreatedID,
		CreatedName: createdName,
		IsShow:      v.IsShow,
	}
}

func (c Community) UpdateCommunityStatus(ID int64, status int64) error {
	_, err := c.db.Community.Update().Where(community2.IDEQ(ID)).SetStatus(status).Save(c.ctx)
	return err
}

func (c Community) UpdateCommunityShow(ID int64, status int64) error {
	_, err := c.db.Community.Update().Where(community2.IDEQ(ID)).SetIsShow(status).Save(c.ctx)
	return err
}

func (c Community) DelCommunity(ID int64) error {
	_, err := c.db.Community.Update().Where(community2.IDEQ(ID)).SetDelete(1).Save(c.ctx)
	return err
}

func NewCommunity(ctx context.Context, c *app.RequestContext) do.Community {
	return &Community{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
