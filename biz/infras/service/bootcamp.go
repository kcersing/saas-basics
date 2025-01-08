package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	bootcamp2 "saas/biz/dal/db/ent/bootcamp"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/minio"

	"saas/biz/dal/db/ent/predicate"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/bootcamp"
	"strconv"
	"time"
)

type Bootcamp struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (b Bootcamp) CreateBootcamp(req bootcamp.BootcampInfo) error {

	signStartAt, _ := time.Parse(time.DateTime, req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateTime, req.SignEndAt)
	startAt, _ := time.Parse(time.DateTime, req.StartAt)
	endAt, _ := time.Parse(time.DateTime, req.EndAt)

	var createdId int64
	userId, exist := b.c.Get("userId")
	if exist || userId != nil {
		uId, ok := userId.(string)
		if ok {
			createdId, _ = strconv.ParseInt(uId, 10, 64)
		}
	}
	_, err := b.db.Bootcamp.Create().
		SetName(req.Name).
		SetDetail(req.Detail).
		SetSignNumber(req.SignNumber).
		SetSignStartAt(signStartAt).
		SetSignEndAt(signEndAt).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetPic(req.Pic).
		SetFee(req.Fee).
		SetIsCancel(req.IsCancel).
		SetCancelTime(req.CancelTime).
		SetSignFields(req.SignFields).
		SetIsFee(req.IsFee).
		SetIsShow(req.IsShow).
		SetCreatedID(createdId).
		Save(b.ctx)
	if err != nil {
		err = errors.Wrap(err, err.Error())
		return err
	}

	return nil
}

func (b Bootcamp) UpdateBootcamp(req bootcamp.BootcampInfo) error {
	signStartAt, _ := time.Parse(time.DateTime, req.SignStartAt)
	signEndAt, _ := time.Parse(time.DateTime, req.SignEndAt)
	startAt, _ := time.Parse(time.DateTime, req.StartAt)
	endAt, _ := time.Parse(time.DateTime, req.EndAt)

	_, err := b.db.Bootcamp.Update().
		Where(bootcamp2.IDEQ(req.ID)).
		SetName(req.Name).
		SetDetail(req.Detail).
		SetSignNumber(req.SignNumber).
		SetSignStartAt(signStartAt).
		SetSignEndAt(signEndAt).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetPic(req.Pic).
		SetFee(req.Fee).
		SetIsCancel(req.IsCancel).
		SetCancelTime(req.CancelTime).
		SetSignFields(req.SignFields).
		SetIsFee(req.IsFee).
		Save(b.ctx)

	if err != nil {
		err = errors.Wrap(err, err.Error())
		return err
	}

	return nil
}

func (b Bootcamp) DeleteBootcamp(id int64) error {
	_, err := b.db.Bootcamp.Update().Where(bootcamp2.IDEQ(id)).SetDelete(1).Save(b.ctx)
	return err
}

func (b Bootcamp) BootcampList(req bootcamp.BootcampListReq) (resp []*bootcamp.BootcampInfo, total int, err error) {
	var predicates []predicate.Bootcamp
	if req.Name != "" {
		predicates = append(predicates, bootcamp2.NameEQ(req.Name))
	}
	if req.Condition != 0 {
		predicates = append(predicates, bootcamp2.Condition(req.Condition))
	}

	if req.SignStartAt != "" && req.SignEndAt != "" {
		signStartAt, _ := time.Parse(time.DateTime, req.SignStartAt)
		signEndAt, _ := time.Parse(time.DateTime, req.SignEndAt)

		predicates = append(predicates, bootcamp2.SignStartAtLTE(signStartAt))
		predicates = append(predicates, bootcamp2.SignEndAtGTE(signEndAt))
	}
	if req.StartAt != "" && req.EndAt != "" {
		startAt, _ := time.Parse(time.DateTime, req.StartAt)
		endAt, _ := time.Parse(time.DateTime, req.EndAt)

		predicates = append(predicates, bootcamp2.StartAtLTE(startAt))
		predicates = append(predicates, bootcamp2.EndAtGTE(endAt))
	}
	predicates = append(predicates, bootcamp2.Delete(0))
	lists, err := b.db.Bootcamp.Query().Where(predicates...).
		Order(ent.Desc(bootcamp2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(b.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, b.entBootcampInfo(v))
	}

	total, _ = b.db.Bootcamp.Query().Where(predicates...).Count(b.ctx)
	return
}

func (b Bootcamp) BootcampInfo(id int64) (resp *bootcamp.BootcampInfo, err error) {
	resp = new(bootcamp.BootcampInfo)

	bootcampInter, exist := b.cache.Get("bootcampInfo" + strconv.Itoa(int(id)))
	if exist {
		if u, ok := bootcampInter.(*bootcamp.BootcampInfo); ok {
			return u, nil
		}
	}
	bootcampEnt, err := b.db.Bootcamp.Query().Where(bootcamp2.IDEQ(id)).First(b.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Bootcamp failed")
		return resp, err
	}
	resp = b.entBootcampInfo(bootcampEnt)
	b.cache.SetWithTTL("bootcampInfo"+strconv.Itoa(int(resp.ID)), &resp, 1, 1*time.Hour)

	return
}

func (b Bootcamp) entBootcampInfo(v *ent.Bootcamp) *bootcamp.BootcampInfo {
	signStartAt := v.SignStartAt.Unix()
	startAt := v.SignStartAt.Unix()
	signEndAt := v.SignStartAt.Unix()
	endAt := v.SignStartAt.Unix()
	createdAt := v.CreatedAt.Unix()
	pic := minio.URLconvert(b.ctx, b.c, v.Pic)
	var createdName string
	first, _ := b.db.User.Query().Where(user.IDEQ(v.CreatedID)).First(b.ctx)
	if first != nil {
		createdName = first.Name
	}

	return &bootcamp.BootcampInfo{
		Pic:         pic,
		ID:          v.ID,
		Name:        v.Name,
		Detail:      v.Detail,
		SignNumber:  v.SignNumber,
		SignStartAt: strconv.FormatInt(signStartAt, 10),
		SignEndAt:   strconv.FormatInt(signEndAt, 10),
		StartAt:     strconv.FormatInt(startAt, 10),
		EndAt:       strconv.FormatInt(endAt, 10),
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

func (b Bootcamp) UpdateBootcampStatus(ID int64, status int64) error {
	_, err := b.db.Bootcamp.Update().Where(bootcamp2.IDEQ(ID)).SetStatus(status).Save(b.ctx)
	return err
}

func (b Bootcamp) UpdateBootcampShow(ID int64, status int64) error {
	_, err := b.db.Bootcamp.Update().Where(bootcamp2.IDEQ(ID)).SetIsShow(status).Save(b.ctx)
	return err
}

func (b Bootcamp) DelBootcamp(ID int64) error {
	_, err := b.db.Bootcamp.Update().Where(bootcamp2.IDEQ(ID)).SetDelete(1).Save(b.ctx)
	return err
}

func NewBootcamp(ctx context.Context, c *app.RequestContext) do.Bootcamp {
	return &Bootcamp{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
