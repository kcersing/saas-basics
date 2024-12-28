package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/entrylogs"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/infras/do"
	"saas/config"
	"saas/idl_gen/model/entry"
	"time"
)

type EntryLogs struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewEntryLogs(ctx context.Context, c *app.RequestContext) do.EntryLogs {
	return &EntryLogs{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (e EntryLogs) Create(logsReq *entry.EntryInfo) error {
	//TODO implement me
	panic("implement me")
}

func (e EntryLogs) List(req *entry.EntryListReq) (list []*entry.EntryInfo, total int, err error) {
	var predicates []predicate.EntryLogs

	if req.VenueId > 0 {
		predicates = append(predicates, entrylogs.VenueID(req.VenueId))
	}
	if req.MemberId > 0 {
		predicates = append(predicates, entrylogs.MemberID(req.MemberId))
	}
	if req.MemberProductId > 0 {
		predicates = append(predicates, entrylogs.MemberProductID(req.MemberProductId))
	}

	if req.UserId > 0 {
		predicates = append(predicates, entrylogs.UserID(req.UserId))
	}

	//if *req.EntryTime != "" {
	//	predicates = append(predicates, entrylogs.EntryTimeEQ(req.EntryTime))
	//}
	//if *req.LeavingTime != "" {
	//	predicates = append(predicates, entrylogs.LeavingTimeGT(req.LeavingTime))
	//}
	predicates = append(predicates, entrylogs.Delete(0))
	lists, err := e.db.EntryLogs.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(e.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Entry list failed")
		return list, total, err
	}

	for _, v := range lists {
		l := entEntryInfo(v)
		list = append(list, l)

	}
	total, _ = e.db.EntryLogs.Query().Where(predicates...).Count(e.ctx)
	return
}

func entEntryInfo(v *ent.EntryLogs) *entry.EntryInfo {
	entryTime := v.EntryTime.Format(time.DateTime)
	leavingTime := v.LeavingTime.Format(time.DateTime)
	return &entry.EntryInfo{
		ID:              v.ID,
		EntryTime:       entryTime,
		LeavingTime:     leavingTime,
		MemberId:        v.MemberID,
		MemberProductId: v.MemberProductID,
		UserId:          v.UserID,
		VenueId:         v.VenueID,
	}
}
