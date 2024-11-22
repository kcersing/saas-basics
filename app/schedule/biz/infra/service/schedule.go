package service

import (
	"context"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"rpc_gen/kitex_gen/schedule"
	"schedule/biz/dal/cache"
	"schedule/biz/dal/mysql"
	"schedule/biz/dal/mysql/ent"
	"schedule/biz/dal/mysql/ent/predicate"
	schedule2 "schedule/biz/dal/mysql/ent/schedule"
	"schedule/biz/infra/do"
	"time"
)

type Schedule struct {
	ctx   context.Context
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func NewSchedule(ctx context.Context) do.Schedule {
	return &Schedule{
		ctx:   ctx,
		salt:  "",
		db:    mysql.DB,
		cache: cache.Cache,
	}
}

func (s Schedule) UpdateSchedule(req schedule.CreateOrUpdateScheduleReq) error {

	return nil
}

func (s Schedule) ScheduleList(req schedule.ScheduleListReq) (resp []*schedule.ScheduleInfo, total int64, err error) {
	var predicates []predicate.Schedule

	if *req.StartTime != "" {
		startTime, _ := time.Parse(time.DateOnly, *req.StartTime)
		//大于
		predicates = append(predicates, schedule2.StartTimeGTE(startTime))
		//小于
		predicates = append(predicates, schedule2.EndTimeLTE(startTime.Add(7*24*time.Hour)))
	}
	if *req.VenueId > 0 {
		predicates = append(predicates, schedule2.VenueID(*req.VenueId))
	}

	lists, err := s.db.Schedule.Query().Where(predicates...).
		Offset(int(*req.Page-1) * int(*req.PageSize)).
		Limit(int(*req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Schedule info failed")
		return resp, 0, err
	}

	for i, v := range lists {
		resp[i].StartTime = v.StartTime.Format("15:04")
		resp[i].EndTime = v.EndTime.Format("15:04")

		coach, _ := v.QueryCoachs().First(s.ctx)
		resp[i].CoachID = coach.CoachID
		resp[i].CoachName = coach.CoachName

		if v.Type == "course" {
			member, _ := v.QueryMembers().First(s.ctx)
			resp[i].MemberName = member.MemberName
			resp[i].MemberProductName = member.MemberProductName
			resp[i].MemberProductPropertyName = member.MemberProductPropertyName
		}
	}

	total, _ = s.db.Schedule.Query().Where(predicates...).Count(s.ctx)
	return
}

func (s Schedule) ScheduleDateList(req schedule.ScheduleListReq) (map[string][]*schedule.ScheduleInfo, int64, error) {
	*req.Page = 1
	*req.PageSize = 1000
	lists, total, err := s.ScheduleList(req)
	m := make(map[string][]*schedule.ScheduleInfo)
	for _, v := range lists {
		m[v.Date] = append(m[v.Date], v)
	}
	return m, total, err
}

func (s Schedule) UpdateScheduleStatus(ID int64, status int64) error {
	_, err := s.db.Schedule.Update().Where(schedule2.IDEQ(ID)).SetStatus(status).Save(s.ctx)
	return err
}

func (sc Schedule) ScheduleInfo(ID int64) (roleInfo *schedule.ScheduleInfo, err error) {
	//TODO implement me
	panic("implement me")
}
