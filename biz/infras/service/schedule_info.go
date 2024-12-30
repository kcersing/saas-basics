package service

import (
	"saas/biz/dal/db/ent"

	"saas/idl_gen/model/schedule"
)

func (s Schedule) ScheduleInfo(ID int64) (roleInfo *schedule.ScheduleInfo, err error) {
	//TODO implement me
	panic("implement me")
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
