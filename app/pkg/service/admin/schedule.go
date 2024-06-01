package admin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/schedule"
	"time"
)

type Schedule struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	Cbs   *casbin.Enforcer
	db    *ent.Client
	cache *ristretto.Cache
}

func NewSchedule(ctx context.Context, c *app.RequestContext) do.Schedule {
	return &Schedule{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
		Cbs:   infras.CasbinEnforcer(),
	}
}

func (s Schedule) ScheduleUpdate(req do.CreateOrUpdateScheduleReq) error {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) ScheduleDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) ScheduleList(req do.ScheduleListReq) (resp []*do.ScheduleInfo, total int, err error) {
	var predicates []predicate.Schedule

	if req.StartTime != "" {
		startTime, _ := time.Parse(time.DateTime, req.StartTime)
		//大于
		predicates = append(predicates, schedule.StartTimeGTE(startTime))
		//小于
		predicates = append(predicates, schedule.EndTimeLTE(startTime.Add(7*24*time.Hour)))
	}
	lists, err := s.db.Schedule.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Order list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Order info failed")
		return resp, 0, err
	}

	for i, v := range lists {
		resp[i].StartTime = v.StartTime.Format(time.DateTime)
		resp[i].EndTime = v.EndTime.Format(time.DateTime)
	}

	total, _ = s.db.Schedule.Query().Where(predicates...).Count(s.ctx)
	return
}
func (s Schedule) ScheduleDateList(req do.ScheduleListReq) (resp map[string][]*do.ScheduleInfo, total int, err error) {
	lists, total, err := s.ScheduleList(req)

	for _, v := range lists {
		resp["1"] = append(resp["1"], v)

	}
	return

}

func (s Schedule) ScheduleUpdateStatus(ID int64, status int64) error {
	_, err := s.db.Schedule.Update().Where(schedule.IDEQ(ID)).SetStatus(int64(status)).Save(s.ctx)
	return err
}

func (sc Schedule) ScheduleInfo(ID int64) (roleInfo *do.ScheduleInfo, err error) {
	//TODO implement me
	panic("implement me")
}
