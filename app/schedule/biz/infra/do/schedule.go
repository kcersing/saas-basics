package do

type Schedule interface {
	CreateSchedule(req CreateOrUpdateScheduleReq) error
	UpdateSchedule(req CreateOrUpdateScheduleReq) error
	//ScheduleDelete(id int64) error
	ScheduleList(req ScheduleListReq) (resp []*ScheduleInfo, total int, err error)
	ScheduleDateList(req ScheduleListReq) (map[string][]*ScheduleInfo, int, error)
	UpdateScheduleStatus(ID int64, status int64) error
	ScheduleInfo(ID int64) (roleInfo *ScheduleInfo, err error)

	CreateMember(req ScheduleMemberCreate) error
	UpdateMember(req ScheduleMemberInfo) error
	DeleteMember(id int64) error

	MemberList(req ScheduleMemberListReq) (resp []*ScheduleMemberInfo, total int, err error)
	UpdateMemberStatus(ID int64, status int64) error
	MemberInfo(ID int64) (roleInfo *ScheduleMemberInfo, err error)
	SearchSubscribeByMember(req SearchSubscribeByMemberReq) (list []SubscribeByMember, total int64, err error)

	CreateCoach(req ScheduleCoachInfo) error
	UpdateCoach(req ScheduleCoachInfo) error
	DeleteCoach(id int64) error

	CoachList(req ScheduleCoachListReq) (resp []*ScheduleCoachInfo, total int, err error)
	UpdateCoachStatus(ID int64, status int64) error
	CoachInfo(ID int64) (roleInfo *ScheduleCoachInfo, err error)
}
