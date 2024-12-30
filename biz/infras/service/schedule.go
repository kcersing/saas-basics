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

func (s Schedule) UpdateSchedule(req schedule.CreateOrUpdateScheduleReq) error {
	//TODO implement me
	panic("implement me")
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

func (s Schedule) ScheduleInfo(ID int64) (roleInfo *schedule.ScheduleInfo, err error) {
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

func (s Schedule) entScheduleMemberInfo(entSchedule *ent.ScheduleMember) (info *schedule.ScheduleMemberInfo) {
	//TODO implement me
	panic("implement me")
}
func (s Schedule) entScheduleCoachInfo(entSchedule *ent.ScheduleCoach) (info *schedule.ScheduleCoachInfo) {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) entScheduleInfo(entSchedule *ent.Schedule) (info *schedule.ScheduleInfo) {
	//TODO implement me
	panic("implement me")
}
