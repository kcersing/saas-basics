package do

import "saas/idl_gen/model/schedule"

type Schedule interface {
	CreateSchedule(req schedule.CreateOrUpdateScheduleReq) error
	UpdateSchedule(req schedule.CreateOrUpdateScheduleReq) error
	DeleteSchedule(id int64) error
	ScheduleList(req schedule.ScheduleListReq) (resp []*schedule.ScheduleInfo, total int, err error)
	ScheduleDateList(req schedule.ScheduleListReq) (map[string][]*schedule.ScheduleInfo, int, error)
	UpdateScheduleStatus(ID int64, status int64) error
	ScheduleInfo(ID int64) (roleInfo *schedule.ScheduleInfo, err error)

	CreateScheduleMember(req schedule.MemberSubscribeReq) error
	UpdateScheduleMember(req schedule.MemberSubscribeReq) error
	DeleteScheduleMember(id int64) error
	ScheduleMemberList(req schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int, err error)
	UpdateScheduleMemberStatus(ID int64, status int64) error
	ScheduleMemberInfo(ID int64) (roleInfo *schedule.ScheduleMemberInfo, err error)
	//SearchSubscribeByMember(req schedule.SearchSubscribeByMemberReq) (list []schedule.SubscribeByMember, total int64, err error)

	//CoachCreate(req schedule.ScheduleCoachInfo) error
	//CoachUpdate(req schedule.ScheduleCoachInfo) error
	//CoachDelete(id int64) error
	ScheduleCoachList(req schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int, err error)
	ScheduleCoachUpdateStatus(ID int64, status int64) error
	ScheduleCoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error)
}
