package service

import (
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/schedulecoach"
	"saas/biz/dal/db/ent/userscheduling"
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
	var usc []*ent.UserSchedulingCreate
	if len(req.UserId) > 0 {
		for _, v := range req.UserId {
			s.db.UserScheduling.Delete().Where(
				userscheduling.UserID(v),
				userscheduling.StartDate(startDate),
				userscheduling.EndDate(endDate),
			).Exec(s.ctx)
			usc = append(usc, s.db.UserScheduling.Create().SetStartDate(startDate).SetEndDate(endDate).SetUserID(v).SetPeriod(*req.Period))
		}
	}

	err = s.db.UserScheduling.CreateBulk(usc...).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil

}

func (s Schedule) UpdateScheduleUserTimePeriod(req schedule.UserTimePeriodReq) error {
	return s.CreateScheduleUserTimePeriod(req)
}

func (s Schedule) ScheduleCoachList(req schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int, err error) {
	var predicates []predicate.ScheduleCoach

	//if req.StartTime != "" {
	//	startTime, _ := time.Parse(time.DateOnly, req.StartTime)
	//	//大于
	//	predicates = append(predicates, schedulemember.StartTimeGTE(startTime))
	//	//小于
	//	predicates = append(predicates, schedulemember.EndTimeLTE(startTime.Add(7*24*time.Hour)))
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
	startTime, err := time.Parse(time.DateTime, req.Date)
	if err != nil {
		return nil, errors.New("日期类型传值错误")
	}
	first, err := s.db.UserScheduling.
		Query().
		Where(
			userscheduling.UserID(req.ID),
			userscheduling.StartDateGTE(startTime),
			userscheduling.EndDateLTE(startTime),
		).
		First(s.ctx)

	if err != nil {
		return nil, err
	}
	resp = &schedule.UserTimePeriodInfo{
		StartDate: first.StartDate.Format(time.DateTime),
		EndDate:   first.EndDate.Format(time.DateTime),
		Period:    &first.Period,
		UserId:    first.UserID,
	}
	return
}
