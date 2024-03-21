package do

import "time"

type CourseRecord interface {
	ScheduleCreate(req CreateOrUpdateScheduleReq) error
	ScheduleUpdate(req CreateOrUpdateScheduleReq) error
	ScheduleDelete(id int64) error
	ScheduleList(req ScheduleListReq) (resp []*ScheduleInfo, total int, err error)
	ScheduleUpdateStatus(ID int64, status int64) error
	ScheduleInfo(ID int64) (roleInfo *ScheduleInfo, err error)

	MemberCreate(req CourseRecordMemberInfo) error
	MemberUpdate(req CourseRecordMemberInfo) error
	MemberDelete(id int64) error
	MemberList(req CourseRecordMemberListReq) (resp []*CourseRecordMemberInfo, total int, err error)
	UpdateMemberStatus(ID int64, status int64) error
	MemberInfo(ID int64) (roleInfo *CourseRecordMemberInfo, err error)

	CoachCreate(req CourseRecordCoachInfo) error
	CoachUpdate(req CourseRecordCoachInfo) error
	CoachDelete(id int64) error
	CoachList(req CourseRecordCoachListReq) (resp []*CourseRecordCoachInfo, total int, err error)
	CoachUpdateStatus(ID int64, status int64) error
	CoachInfo(ID int64) (roleInfo *CourseRecordCoachInfo, err error)
}

type CreateOrUpdateScheduleReq struct {
	ID        int64     `json:"id,omitempty"`
	Type      string    `json:"type,omitempty"`
	VenueId   int64     `json:"venue_id,omitempty"`
	PlaceID   int64     `json:"place_id,omitempty"`
	Num       int64     `json:"num,omitempty"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Price     float64   `json:"price,omitempty"`
	CoachID   int64     `json:"coach_id,omitempty"`

	MemberID            int64 `json:"member_id,omitempty"`
	MemberProductID     int64 `json:"member_product_id,omitempty"`
	MemberProductItemID int64 `json:"member_product_item_id,omitempty"`
}

type ScheduleInfo struct {
	ID        int64     `json:"id,omitempty"`
	VenueId   int64     `json:"venue_id,omitempty"`
	Type      string    `json:"type,omitempty"`
	PlaceID   int64     `json:"place_id,omitempty"`
	Num       int64     `json:"num,omitempty"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Price     float64   `json:"price,omitempty"`

	MemberCourseRecord []CourseRecordMemberInfo `json:"member_course_record,omitempty"`
	CoachCourseRecord  []CourseRecordCoachInfo  `json:"coach_course_record,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

type CourseRecordCoachInfo struct {
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

type CourseRecordCoachListReq struct {
}
