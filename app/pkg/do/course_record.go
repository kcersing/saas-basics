package do

import "time"

type CourseRecord interface {
	ScheduleCreate(req ScheduleInfo) error
	ScheduleUpdate(req ScheduleInfo) error
	ScheduleDelete(id int64) error
	ScheduleList(req ScheduleListReq) (resp []*ScheduleInfo, total int, err error)
	ScheduleUpdateStatus(ID int64, status int8) error
	ScheduleInfo(ID int64) (roleInfo *ScheduleInfo, err error)

	UserCreate(req CourseRecordMemberInfo) error
	UserUpdate(req CourseRecordMemberInfo) error
	UserDelete(id int64) error
	UserList(req CourseRecordMemberListReq) (resp []*CourseRecordMemberInfo, total int, err error)
	UserUpdateStatus(ID int64, status int8) error
	UserInfo(ID int64) (roleInfo *CourseRecordMemberInfo, err error)

	CoachCreate(req CoachInfo) error
	CoachUpdate(req CoachInfo) error
	CoachDelete(id int64) error
	CoachList(req CoachListReq) (resp []*CoachInfo, total int, err error)
	CoachUpdateStatus(ID int64, status int8) error
	CoachInfo(ID int64) (roleInfo *CoachInfo, err error)
}

type ScheduleInfo struct {
	ID        int64     `json:"id,omitempty"`
	VenueId   int64     `json:"venue_id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ScheduleListReq struct {
}

type CourseRecordMemberInfo struct {
	ID                     int64     `json:"id,omitempty"`
	MemberId               int64     `json:"member_id,omitempty"`
	VenueId                int64     `json:"venue_id,omitempty"`
	CourseRecordScheduleId int64     `json:"course_record_schedule_id,omitempty"`
	Type                   int64     `json:"type,omitempty"`
	CreatedAt              time.Time `json:"createdAt,omitempty"`
	UpdatedAt              time.Time `json:"updatedAt,omitempty"`
	StartTime              time.Time `json:"start_time,omitempty"`
	EndTime                time.Time `json:"end_time,omitempty"`
	SignStartTime          time.Time `json:"sign_start_time,omitempty"`
	SignEndTime            time.Time `json:"sign_end_time,omitempty"`
	Status                 int64     `json:"status,omitempty"`
}

type CourseRecordMemberListReq struct {
}

type CoachInfo struct {
	ID                     int64     `json:"id,omitempty"`
	CoachId                int64     `json:"coach_id,omitempty"`
	VenueId                int64     `json:"venue_id,omitempty"`
	CourseRecordScheduleId int64     `json:"course_record_schedule_id,omitempty"`
	Type                   int64     `json:"type,omitempty"`
	CreatedAt              time.Time `json:"createdAt,omitempty"`
	UpdatedAt              time.Time `json:"updatedAt,omitempty"`
	StartTime              time.Time `json:"start_time,omitempty"`
	EndTime                time.Time `json:"end_time,omitempty"`
	SignStartTime          time.Time `json:"sign_start_time,omitempty"`
	SignEndTime            time.Time `json:"sign_end_time,omitempty"`
	Status                 int64     `json:"status,omitempty"`
}

type CoachListReq struct {
}
