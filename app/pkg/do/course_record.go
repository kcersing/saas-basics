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
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	VenueId   int64     `json:"venue_id"`
	PlaceID   int64     `json:"place_id"`
	Num       int64     `json:"num"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Price     float64   `json:"price"`
	CoachID   int64     `json:"coach_id"`

	MemberID            int64 `json:"member_id"`
	MemberProductID     int64 `json:"member_product_id"`
	MemberProductItemID int64 `json:"member_product_item_id"`
}

type ScheduleInfo struct {
	ID        int64     `json:"id"`
	VenueId   int64     `json:"venue_id"`
	Type      string    `json:"type"`
	PlaceID   int64     `json:"place_id"`
	Num       int64     `json:"num"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Price     float64   `json:"price"`

	MemberCourseRecord []CourseRecordMemberInfo `json:"member_course_record"`
	CoachCourseRecord  []CourseRecordCoachInfo  `json:"coach_course_record"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ScheduleListReq struct {
}

type CourseRecordMemberInfo struct {
	ID                     int64     `json:"id"`
	MemberId               int64     `json:"member_id"`
	VenueId                int64     `json:"venue_id"`
	CourseRecordScheduleId int64     `json:"course_record_schedule_id"`
	Type                   int64     `json:"type"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
	StartTime              time.Time `json:"start_time"`
	EndTime                time.Time `json:"end_time"`
	SignStartTime          time.Time `json:"sign_start_time"`
	SignEndTime            time.Time `json:"sign_end_time"`
	Status                 int64     `json:"status"`
}

type CourseRecordMemberListReq struct {
}

type CourseRecordCoachInfo struct {
	ID                     int64     `json:"id"`
	CoachId                int64     `json:"coach_id"`
	VenueId                int64     `json:"venue_id"`
	CourseRecordScheduleId int64     `json:"course_record_schedule_id"`
	Type                   int64     `json:"type"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
	StartTime              time.Time `json:"start_time"`
	EndTime                time.Time `json:"end_time"`
	SignStartTime          time.Time `json:"sign_start_time"`
	SignEndTime            time.Time `json:"sign_end_time"`
	Status                 int64     `json:"status"`
}

type CourseRecordCoachListReq struct {
}
