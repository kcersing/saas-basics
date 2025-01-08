package service

import (
	"errors"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/userscheduling"
	"saas/idl_gen/model/schedule"
	"time"
)

func (s Schedule) CreateScheduleUserTimePeriod(req schedule.UserTimePeriodReq) error {
	startDate, err := time.Parse(time.DateOnly, req.StartDate)

	if err != nil {
		return errors.New("日期类型传值错误")
	}
	endDate, err := time.Parse(time.DateOnly, req.EndDate)

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
	//TODO implement me
	panic("implement me")
}

func (s Schedule) UpdateScheduleCoachStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) ScheduleCoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error) {
	//TODO implement me
	panic("implement me")
}
