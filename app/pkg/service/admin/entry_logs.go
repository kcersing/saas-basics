package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/entrylogs"
	"saas/pkg/db/ent/predicate"
)

type EntryLogs struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (e EntryLogs) Create(logsReq *do.EntryLogsInfo) error {
	//TODO implement me
	panic("implement me")
}

func (e EntryLogs) List(req *do.EntryLogsListReq) (list []*do.EntryLogsInfo, total int, err error) {
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
	if req.MemberPropertyId > 0 {
		predicates = append(predicates, entrylogs.MemberPropertyID(req.MemberPropertyId))
	}

	if req.UserId > 0 {
		predicates = append(predicates, entrylogs.UserID(req.UserId))
	}

	//if req.EntryTime != "" {
	//	predicates = append(predicates, entrylogs.EntryTimeEQ(req.EntryTime))
	//}
	//if req.LeavingTime != "" {
	//	predicates = append(predicates, entrylogs.LeavingTimeGT(req.LeavingTime))
	//}

	lists, err := e.db.EntryLogs.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(e.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Entry list failed")
		return list, total, err
	}

	err = copier.Copy(&list, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Entry info failed")
		return list, 0, err
	}

	total, _ = e.db.EntryLogs.Query().Where(predicates...).Count(e.ctx)
	return
}

func NewEntryLogs(ctx context.Context, c *app.RequestContext) do.EntryLogs {
	return &EntryLogs{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
