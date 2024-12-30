package do

import "saas/idl_gen/model/schedule"

type Schedule interface {
	CreateSchedule(req schedule.CreateOrUpdateScheduleReq) error
	UpdateSchedule(req schedule.CreateOrUpdateScheduleReq) error
	DeleteSchedule(id int64) error
	ScheduleList(req schedule.ListScheduleReq) (resp []*schedule.ScheduleInfo, total int, err error)
	ScheduleDateList(req schedule.ListScheduleReq) (map[string][]*schedule.ScheduleInfo, int, error)
	UpdateScheduleStatus(ID int64, status int64) error
	ScheduleInfo(ID int64) (roleInfo *schedule.ScheduleInfo, err error)

	MemberCreate(req schedule.MemberSubscribeReq) error
	MemberUpdate(req schedule.MemberSubscribeReq) error
	MemberDelete(id int64) error
	MemberList(req schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int, err error)
	UpdateMemberStatus(ID int64, status int64) error
	MemberInfo(ID int64) (roleInfo *schedule.ScheduleMemberInfo, err error)
	//SearchSubscribeByMember(req schedule.SearchSubscribeByMemberReq) (list []schedule.SubscribeByMember, total int64, err error)

	//CoachCreate(req schedule.ScheduleCoachInfo) error
	//CoachUpdate(req schedule.ScheduleCoachInfo) error
	//CoachDelete(id int64) error
	CoachList(req schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int, err error)
	CoachUpdateStatus(ID int64, status int64) error
	CoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error)
}
