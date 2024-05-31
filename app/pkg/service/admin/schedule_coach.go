package admin

import (
	"saas/app/pkg/do"
)

func (c Schedule) CoachCreate(req do.ScheduleCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) CoachUpdate(req do.ScheduleCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) CoachDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) CoachList(req do.ScheduleCoachListReq) (resp []*do.ScheduleCoachInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) CoachUpdateStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) CoachInfo(ID int64) (roleInfo *do.ScheduleCoachInfo, err error) {
	//TODO implement me
	panic("implement me")
}
