package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"

	"saas/biz/infras/do"
	"saas/config"

	"saas/idl_gen/model/schedule"
)

type Schedule struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (s Schedule) DeleteSchedule(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) ScheduleList(req schedule.ScheduleListReq) (resp []*schedule.ScheduleInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) ScheduleDateList(req schedule.ScheduleListReq) (map[string][]*schedule.ScheduleInfo, int, error) {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) UpdateScheduleStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func NewSchedule(ctx context.Context, c *app.RequestContext) do.Schedule {
	return &Schedule{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
