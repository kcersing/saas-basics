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
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/venue"
)

type Venue struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (v Venue) Create(logsReq *do.VenueInfo) error {
	//TODO implement me
	panic("implement me")
}

func (v Venue) List(req *do.VenueListReq) (list []*do.VenueInfo, total int, err error) {
	var predicates []predicate.Venue

	if req.Name != "" {
		predicates = append(predicates, venue.NameEQ(req.Name))
	}

	venueList, err := v.db.Venue.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(v.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user list failed")
		return list, total, err
	}
	// copy to UserInfo struct
	err = copier.Copy(&list, &venueList)
	if err != nil {
		err = errors.Wrap(err, "copy user info failed")
		return list, 0, err
	}
	total, _ = v.db.Venue.Query().Where(predicates...).Count(v.ctx)
	return
}

func (v Venue) UpdateVenueStatus(id int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (v Venue) DeleteAll() error {
	//TODO implement me
	panic("implement me")
}

func NewVenue(ctx context.Context, c *app.RequestContext) do.Venue {
	return &Venue{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
	}
}
