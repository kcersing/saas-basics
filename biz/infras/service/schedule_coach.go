package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/schedulecoach"
	"saas/biz/dal/db/ent/usertimeperiod"
	"saas/idl_gen/model/schedule"
	"time"
)

func (s Schedule) CreateScheduleUserTimePeriod(req schedule.UserTimePeriodReq) error {
	startDate, err := time.Parse(time.DateTime, req.StartDate)

	if err != nil {
		return errors.New("日期类型传值错误")
	}
	endDate, err := time.Parse(time.DateTime, req.EndDate)

	if err != nil {
		return errors.New("日期类型传值错误")
	}
	var usc []*ent.UserTimePeriodCreate
	if len(req.UserId) > 0 {
		for _, v := range req.UserId {
			//Before
			//
			for {
				date := startDate
				usc = append(usc, s.db.UserTimePeriod.
					Create().
					SetDate(date).
					SetUserID(v).
					SetVenueID(req.VenueId).
					SetPeriod(*req.Period))

				date = startDate.Add(1)
				if date.After(endDate) {
					continue
				}
			}

		}
	}

	err = s.db.UserTimePeriod.CreateBulk(usc...).Exec(s.ctx)
	hlog.Info(err)

	return nil

}

func (s Schedule) UpdateScheduleUserTimePeriod(req schedule.UpdateUserTimePeriodReq) error {
	date, err := time.Parse(time.DateTime, req.Date)

	if err != nil {
		return errors.New("日期类型传值错误")
	}
	_, err = s.db.UserTimePeriod.
		Create().
		SetUserID(req.UserId).
		SetVenueID(req.VenueId).
		SetPeriod(*req.Period).
		SetDate(date).
		Save(s.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s Schedule) ScheduleCoachList(req schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int, err error) {
	var predicates []predicate.ScheduleCoach

	//if req.StartTime != "" {
	//	startTime, _ := time.Parse(time.DateOnly, req.StartTime)
	//	//大于
	//	predicates = append(predicates, schedulemember.StartTimeLTE(startTime))
	//	//小于
	//	predicates = append(predicates, schedulemember.EndTimeGTE(startTime.Add(7*24*time.Hour)))
	//}
	if req.CoachId > 0 {
		predicates = append(predicates, schedulecoach.CoachID(req.CoachId))
	}
	if req.ScheduleId > 0 {
		predicates = append(predicates, schedulecoach.ScheduleID(req.ScheduleId))
	}
	if req.Type != "" {
		predicates = append(predicates, schedulecoach.Type(req.Type))

	}
	lists, err := s.db.ScheduleCoach.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Coach list failed")
		return resp, total, err
	}
	for _, v := range lists {
		resp = append(resp, s.entScheduleCoachInfo(v))

	}

	total, _ = s.db.ScheduleCoach.Query().Where(predicates...).Count(s.ctx)
	return
}

func (s Schedule) UpdateScheduleCoachStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) ScheduleCoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error) {
	//TODO implement me
	panic("implement me")
}
func (s Schedule) UserTimePeriod(req schedule.UserPeriodReq) (resp *schedule.UserTimePeriodInfo, err error) {
	startTime, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		return nil, errors.New("日期类型传值错误")
	}

	var predicates []predicate.UserTimePeriod
	if req.VenueId > 0 {
		predicates = append(predicates, usertimeperiod.VenueID(req.VenueId))
	}

	predicates = append(predicates, usertimeperiod.UserID(req.ID))
	predicates = append(predicates, usertimeperiod.Date(startTime))
	first, err := s.db.Debug().UserTimePeriod.
		Query().
		Where(predicates...).
		First(s.ctx)

	if err != nil {
		return nil, err
	}
	resp = &schedule.UserTimePeriodInfo{
		Date:   first.Date.Format(time.DateTime),
		Period: &first.Period,
		UserId: first.UserID,
	}
	return
}
