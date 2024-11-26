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
	SearchSubscribeByMember(req SearchSubscribeByMemberReq) (list []SubscribeByMember, total int64, err error)

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

	Num                       int64   `json:"num"`
	NumSurplus                int64   `json:"num_surplus"`
	Date                      string  `json:"data"`
	StartTime                 string  `json:"start_time"`
	EndTime                   string  `json:"end_time"`
	Price                     float64 `json:"price"`
	Name                      string  `json:"name"`
	Remark                    string  `json:"remark"`
	CoachID                   int64   `json:"coach_id"`
	MemberID                  int64   `json:"member_id"`
	MemberProductID           int64   `json:"member_product_id"`
	MemberProductPropertyID   int64   `json:"member_product_property_id"`
	Status                    int64   `json:"status"`
	PropertyName              string  `json:"property_name"`
	VenueName                 string  `json:"venue_name"`
	PlaceName                 string  `json:"place_name"`
	CoachName                 string  `json:"coach_name"`
	MemberName                string  `json:"member_name"`
	MemberProductName         string  `json:"member_product_name"`
	MemberProductPropertyName string  `json:"member_product_property_name"`

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
	ID                  int64  `json:"id"`
	MemberId            int64  `json:"member_id"`
	VenueId             int64  `json:"venue_id"`
	PlaceID             int64  `json:"place_id"`
	PropertyId          int64  `json:"property_id"`
	ScheduleId          int64  `json:"schedule_id"`
	ScheduleName        string `json:"schedule_name"`
	Type                string `json:"type"`
	CreatedAt           string `json:"createdAt"`
	UpdatedAt           string `json:"updatedAt"`
	StartTime           string `json:"start_time"`
	EndTime             string `json:"end_time"`
	SignStartTime       string `json:"sign_start_time"`
	SignEndTime         string `json:"sign_end_time"`
	Status              int64  `json:"status"`
	MemberProductID     int64  `json:"member_product_id"`
	MemberProductItemID int64  `json:"member_product_property_id"`

	VenueName                 string `json:"venue_name"`
	MemberName                string `json:"member_name"`
	MemberProductName         string `json:"member_product_name"`
	MemberProductPropertyName string `json:"member_product_property_name"`
	Gender                    string `json:"gender"`
	Birthday                  int64  `json:"birthday"`
	Mobile                    string `json:"mobile"`
}
type ScheduleMemberCreate struct {
	MemberProductPropertyID []int64 `json:"memberProductPropertyId"`
	Schedule                int64   `json:"schedule"`
	Remark                  string  `json:"remark"`
}
type ScheduleMemberListReq struct {
	Type     string `json:"type"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Member   int64  `json:"member"`
	Schedule int64  `json:"schedule"`
}

type ScheduleCoachInfo struct {
	ID            int64  `json:"id"`
	CoachId       int64  `json:"coach_id"`
	VenueId       int64  `json:"venue_id"`
	PlaceID       int64  `json:"place_id"`
	PropertyId    int64  `json:"property_id"`
	ScheduleId    int64  `json:"schedule_id"`
	Type          string `json:"type"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Date          string `json:"date"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	SignStartTime string `json:"sign_start_time"`
	SignEndTime   string `json:"sign_end_time"`
	Status        int64  `json:"status"`

	ScheduleName string `json:"schedule_name"`
	PropertyName string `json:"property_name"`
	VenueName    string `json:"venue_name"`
	PlaceName    string `json:"place_name"`
	CoachName    string `json:"coach_name"`
	CoachAvatar  string `json:"coach_avatar"`

	Mobile                    string `json:"mobile"`
	MemberName                string `json:"member_name"`
	MemberAvatar              string `json:"member_avatar"`
	MemberProductName         string `json:"member_product_name"`
	MemberProductPropertyName string `json:"member_product_property_name"`
	Remark                    string `json:"remark"`
	MRemark                   string `json:"m_remark"`
}

type ScheduleCoachListReq struct {
	Type     string `json:"type"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Coach    int64  `json:"coach"`
	Schedule int64  `json:"schedule"`
}

type SubscribeByMember struct {
	Avatar                  string `json:"avatar"`
	Mobile                  string `json:"mobile"`
	MemberID                int64  `json:"member_id"`
	MemberProductID         int64  `json:"member_product_id"`
	MemberProductPropertyId int64  `json:"member_product_property_id"`

	MemberName                string `json:"member_name"`
	MemberProductName         string `json:"member_product_name"`
	MemberProductPropertyName string `json:"member_product_property_name"`
}
type SearchSubscribeByMemberReq struct {
	PropertyId int64  `json:"propertyId"`
	Mobile     string `json:"mobile"`
	Venue      int64  `json:"venue"`
	Type       string `json:"type"`
}
