package service

import (
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/schedulecoach"
	"saas/biz/dal/db/ent/usertimeperiod"
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/schedule"
	"time"
)

func (s Schedule) CreateScheduleUserTimePeriod(req schedule.UserTimePeriodReq) error {
	if req.VenueId == 0 {
		return errors.New("场馆ID不能为空")
	}
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
			for date := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location()); date.Before(endDate); date = date.Add(24 * time.Hour) {
				usc = append(usc, s.db.UserTimePeriod.
					Create().
					SetDate(date).
					SetUserID(v).
					SetVenueID(req.VenueId).
					SetPeriod(*req.Period))
			}
		}
	}

	err = s.db.UserTimePeriod.CreateBulk(usc...).Exec(s.ctx)
	if err != nil {
		return errors.New("值班信息创建失败")
	}

	return nil

}

func (s Schedule) UpdateScheduleUserTimePeriod(req schedule.UpdateUserTimePeriodReq) error {
	startTime, err := time.Parse(time.DateTime, req.Date)
	if err != nil {
		return errors.New("日期类型传值错误")
	}

	if req.VenueId == 0 {
		return errors.New("场馆ID不能为空")
	}

	date := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())

	_, err = s.db.UserTimePeriod.
		Update().
		Where(
			usertimeperiod.UserIDEQ(req.UserId),
			usertimeperiod.DateEQ(date),
			usertimeperiod.VenueIDEQ(req.VenueId),
		).
		SetPeriod(*req.Period).
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

		k := s.entScheduleCoachInfo(v)

		resp = append(resp, k)

	}

	total, _ = s.db.ScheduleCoach.Query().Where(predicates...).Count(s.ctx)
	return
}

func (s Schedule) ScheduleCoachPeriodList(req schedule.UserPeriodReq) (resp []*schedule.ScheduleCoachPeriod, total int, err error) {
	if req.VenueId == 0 {
		return nil, 0, errors.New("场馆ID不能为空")
	}
	startTime, err := time.Parse(time.DateTime, req.Date)
	if err != nil {
		return nil, 0, errors.New("日期类型传值错误")
	}
	date := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())

	userArr, err := s.db.UserTimePeriod.Query().Where(
		usertimeperiod.Date(date),
		usertimeperiod.VenueID(req.VenueId),
	).All(s.ctx)

	if userArr == nil {
		return nil, 0, errors.New("未查询到值班教练")
	}
	for _, v := range userArr {

		u, _ := v.QueryUsers().First(s.ctx)

		coach := &schedule.ScheduleCoachPeriod{
			Date:      date.Format(time.DateOnly),
			CoachId:   u.ID,
			VenueId:   u.DefaultVenueID,
			CoachName: u.Name,
		}
		var userTags []*base.List
		Tags, _ := u.QueryTags().All(s.ctx)
		if len(Tags) > 0 {
			for _, d := range Tags {
				userTags = append(userTags, &base.List{
					ID:   d.ID,
					Name: d.Title,
				})
			}
		}
		coach.Tags = userTags
		coach.Period = &v.Period
		coach.PeriodId = v.ID

		var predicates []predicate.ScheduleCoach
		predicates = append(predicates, schedulecoach.CoachID(u.ID))
		predicates = append(predicates, schedulecoach.Date(date))
		predicates = append(predicates, schedulecoach.StatusNotIn(0, 5))
		if req.Status > 0 {
			predicates = append(predicates, schedulecoach.Status(req.Status))
		}

		lists, _ := s.db.ScheduleCoach.Query().Where(predicates...).All(s.ctx)

		if len(lists) > 0 {
			for _, d := range lists {
				coach.ScheduleCoachList = append(coach.ScheduleCoachList, s.entScheduleCoachInfo(d))
			}
		}
		resp = append(resp, coach)
	}

	return resp, len(userArr), nil
}
func (s Schedule) UpdateScheduleCoachStatus(ID int64, status int64) error {
	_, err := s.db.ScheduleCoach.Update().Where(schedulecoach.ID(ID)).SetStatus(status).Save(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s Schedule) ScheduleCoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error) {
	first, err := s.db.ScheduleCoach.Query().Where(schedulecoach.ID(ID)).First(s.ctx)
	if err != nil {
		return nil, err
	}
	return s.entScheduleCoachInfo(first), nil
}
func (s Schedule) UserTimePeriod(req schedule.UserPeriodReq) (resp *schedule.UserTimePeriodInfo, err error) {
	if req.VenueId == 0 {
		return nil, errors.New("场馆ID不能为空")
	}
	startTime, err := time.Parse(time.DateTime, req.Date)
	if err != nil {
		return nil, errors.New("日期类型传值错误")
	}
	date := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())

	var predicates []predicate.UserTimePeriod

	predicates = append(predicates, usertimeperiod.VenueID(req.VenueId))
	predicates = append(predicates, usertimeperiod.UserID(req.ID))
	predicates = append(predicates, usertimeperiod.Date(date))
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
func (s Schedule) DeleteUserTimePeriod(ID int64) error {
	err := s.db.UserTimePeriod.DeleteOneID(ID).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}
