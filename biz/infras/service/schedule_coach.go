package service

import "saas/idl_gen/model/schedule"

func (s Schedule) CreateScheduleUserTimePeriod(req schedule.UserTimePeriodReq) error {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) UpdateScheduleUserTimePeriod(req schedule.UserTimePeriodReq) error {
	//TODO implement me
	panic("implement me")
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
