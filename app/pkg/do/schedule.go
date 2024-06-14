package do

import "time"

type Schedule interface {
	ScheduleCreate(req CreateOrUpdateScheduleReq) error
	ScheduleUpdate(req CreateOrUpdateScheduleReq) error
	ScheduleDelete(id int64) error
	ScheduleList(req ScheduleListReq) (resp []*ScheduleInfo, total int, err error)
	ScheduleDateList(req ScheduleListReq) (map[string][]*ScheduleInfo, int, error)
	ScheduleUpdateStatus(ID int64, status int64) error
	ScheduleInfo(ID int64) (roleInfo *ScheduleInfo, err error)

	MemberCreate(req ScheduleMemberCreate) error
	MemberUpdate(req ScheduleMemberInfo) error
	MemberDelete(id int64) error
	MemberList(req ScheduleMemberListReq) (resp []*ScheduleMemberInfo, total int, err error)
	UpdateMemberStatus(ID int64, status int64) error
	MemberInfo(ID int64) (roleInfo *ScheduleMemberInfo, err error)

	CoachCreate(req ScheduleCoachInfo) error
	CoachUpdate(req ScheduleCoachInfo) error
	CoachDelete(id int64) error
	CoachList(req ScheduleCoachListReq) (resp []*ScheduleCoachInfo, total int, err error)
	CoachUpdateStatus(ID int64, status int64) error
	CoachInfo(ID int64) (roleInfo *ScheduleCoachInfo, err error)
}

type CreateOrUpdateScheduleReq struct {
	ID                      int64   `json:"id"`
	Type                    string  `json:"type"`
	PropertyId              int64   `json:"propertyId"`
	VenueId                 int64   `json:"venueId"`
	PlaceID                 int64   `json:"placeId"`
	Num                     int64   `json:"num"`
	StartTime               string  `json:"startTime"`
	Price                   float64 `json:"price"`
	Remark                  string  `json:"remark"`
	CoachID                 int64   `json:"coachId"`
	MemberID                int64   `json:"memberId"`
	MemberProductID         int64   `json:"memberProductId"`
	MemberProductPropertyId int64   `json:"memberProductPropertyId"`
}

type ScheduleInfo struct {
	ID int64 `json:"id"`

	Type       string `json:"type"`
	PropertyId int64  `json:"property_id"`
	VenueId    int64  `json:"venue_id"`
	PlaceID    int64  `json:"place_id"`

	Num                 int64   `json:"num"`
	NumSurplus          int64   `json:"num_surplus"`
	Date                string  `json:"data"`
	StartTime           string  `json:"start_time"`
	EndTime             string  `json:"end_time"`
	Price               float64 `json:"price"`
	Name                string  `json:"name"`
	Remark              string  `json:"remark"`
	CoachID             int64   `json:"coach_id"`
	MemberID            int64   `json:"member_id"`
	MemberProductID     int64   `json:"member_product_id"`
	MemberProductItemID int64   `json:"member_product_property_id"`

	PropertyName          string `json:"property_name"`
	VenueName             string `json:"venue_name"`
	PlaceName             string `json:"place_name"`
	CoachName             string `json:"coach_name"`
	MemberName            string `json:"member_name"`
	MemberProductName     string `json:"member_product_name"`
	MemberProductItemName string `json:"member_product_property_name"`

	ScheduleMember []ScheduleMemberInfo `json:"member_course_record"`
	ScheduleCoach  []ScheduleCoachInfo  `json:"coach_course_record"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ScheduleListReq struct {
	PlaceID    int64  `json:"place"`
	PropertyId int64  `json:"property"`
	CoachId    int64  `json:"coach"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	VenueId    int64  `json:"venueId"`
	Page       int64  `json:"page"`
	PageSize   int64  `json:"pageSize"`
}

type ScheduleMemberInfo struct {
	ID                  int64     `json:"id"`
	MemberId            int64     `json:"member_id"`
	VenueId             int64     `json:"venue"`
	PlaceID             int64     `json:"place_id"`
	PropertyId          int64     `json:"property"`
	ScheduleScheduleId  int64     `json:"course_record_schedule_id"`
	Type                int64     `json:"type"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	StartTime           time.Time `json:"start_time"`
	EndTime             time.Time `json:"end_time"`
	SignStartTime       time.Time `json:"sign_start_time"`
	SignEndTime         time.Time `json:"sign_end_time"`
	Status              int64     `json:"status"`
	MemberProductID     int64     `json:"member_product_id"`
	MemberProductItemID int64     `json:"member_product_property_id"`

	VenueName             string `json:"venue_name"`
	MemberName            string `json:"member_name"`
	MemberProductName     string `json:"member_product_name"`
	MemberProductItemName string `json:"member_product_property_name"`
}
type ScheduleMemberCreate struct {
	Member   []int64 `json:"member"`
	Schedule int64   `json:"schedule"`
}
type ScheduleMemberListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	Member   int64 `json:"member"`
	Schedule int64 `json:"schedule"`
}

type ScheduleCoachInfo struct {
	ID                 int64     `json:"id"`
	CoachId            int64     `json:"coach_id"`
	VenueId            int64     `json:"venue"`
	PlaceID            int64     `json:"place_id"`
	PropertyId         int64     `json:"property"`
	ScheduleScheduleId int64     `json:"course_record_schedule_id"`
	Type               int64     `json:"type"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	StartTime          time.Time `json:"start_time"`
	EndTime            time.Time `json:"end_time"`
	SignStartTime      time.Time `json:"sign_start_time"`
	SignEndTime        time.Time `json:"sign_end_time"`
	Status             int64     `json:"status"`

	PropertyName string `json:"property_name"`
	VenueName    string `json:"venue_name"`
	PlaceName    string `json:"place_name"`
	CoachName    string `json:"coach_name"`
}

type ScheduleCoachListReq struct {
}
