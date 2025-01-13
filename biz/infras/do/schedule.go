package do

import (
	"saas/biz/dal/db/ent"
	"saas/idl_gen/model/schedule"
	"time"
)

type Schedule interface {
	CreateScheduleUserTimePeriod(req schedule.UserTimePeriodReq) error
	UpdateScheduleUserTimePeriod(req schedule.UpdateUserTimePeriodReq) error

	CreateScheduleCourse(req schedule.CreateOrUpdateScheduleCourseReq) error
	CreateScheduleLessons(req schedule.CreateOrUpdateScheduleLessonsReq) error

	//UpdateSchedule(req schedule.CreateOrUpdateScheduleReq) error

	DeleteSchedule(id int64) error
	ScheduleList(req schedule.ScheduleListReq, isSubList bool) (resp []*schedule.ScheduleInfo, total int, err error)
	ScheduleDateList(req schedule.ScheduleListReq) (map[string][]*schedule.ScheduleInfo, int, error)
	UpdateScheduleStatus(ID int64, status int64) error
	ScheduleInfo(ID int64) (resp *schedule.ScheduleInfo, err error)

	ScheduleLessonsPublish(ids []int64) error
	CreateMemberSubscribeLessons(req schedule.MemberSubscribeReq) error
	//UpdateScheduleMember(req schedule.MemberSubscribeReq) error
	//DeleteScheduleMember(id int64) error
	ScheduleMemberList(req schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int, err error)
	UpdateScheduleMemberStatus(ID int64, status int64) error
	ScheduleMemberInfo(ID int64) (roleInfo *schedule.ScheduleMemberInfo, err error)
	//SearchSubscribeByMember(req schedule.SearchSubscribeByMemberReq) (list []schedule.SubscribeByMember, total int64, err error)
	UserTimePeriod(req schedule.UserPeriodReq) (resp *schedule.UserTimePeriodInfo, err error)
	CreateScheduleMemberCourse(req CreateScheduleMemberCourse) error

	//CoachCreate(req schedule.ScheduleCoachInfo) error
	//CoachUpdate(req schedule.ScheduleCoachInfo) error
	//CoachDelete(id int64) error
	ScheduleCoachList(req schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int, err error)

	ScheduleCoachPeriodList(req schedule.UserPeriodReq) (resp []*schedule.ScheduleCoachPeriod, total int, err error)

	UpdateScheduleCoachStatus(ID int64, status int64) error
	ScheduleCoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error)

	DeleteUserTimePeriod(ID int64) error
}

type CreateScheduleMemberCourse struct {
	MemberId        int64
	VenueId         int64
	StartTime       time.Time
	MemberProductId int64
	ProductId       int64
	Type            string
	One             *ent.Schedule
	PlaceId         int64
}
