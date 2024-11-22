package do

import "rpc_gen/kitex_gen/schedule"

type Schedule interface {
	CreateSchedule(req schedule.CreateOrUpdateScheduleReq) error
	UpdateSchedule(req schedule.CreateOrUpdateScheduleReq) error
	//ScheduleDelete(id int64) error
	ScheduleList(req schedule.ScheduleListReq) (resp []*schedule.ScheduleInfo, total int64, err error)
	ScheduleDateList(req schedule.ScheduleListReq) (map[string][]*schedule.ScheduleInfo, int64, error)
	UpdateScheduleStatus(ID int64, status int64) error
	ScheduleInfo(ID int64) (roleInfo *schedule.ScheduleInfo, err error)

	CreateMember(req schedule.CreateMemberReq) error
	UpdateMember(req schedule.ScheduleMemberInfo) error
	DeleteMember(id int64) error

	MemberList(req schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int64, err error)
	UpdateMemberStatus(ID int64, status int64) error
	MemberInfo(ID int64) (roleInfo *schedule.ScheduleMemberInfo, err error)
	SearchSubscribeByMember(req schedule.SearchSubscribeByMemberReq) (list []*schedule.SubscribeByMember, total int64, err error)

	CreateCoach(req schedule.ScheduleCoachInfo) error
	UpdateCoach(req schedule.ScheduleCoachInfo) error
	DeleteCoach(id int64) error

	CoachList(req schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int64, err error)
	UpdateCoachStatus(ID int64, status int64) error
	CoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error)
}
