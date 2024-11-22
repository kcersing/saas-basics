namespace go schedule

include "../base/base.thrift"

struct CreateOrUpdateScheduleReq {
    1:optional i64 id  (api.raw = "id")
    2:optional string type  (api.raw = "type")
    3:optional i64 propertyId  (api.raw = "propertyId")
    4:optional i64 venueId  (api.raw = "venueId")
    5:optional i64 placeId  (api.raw = "placeId")
    6:optional i64 num  (api.raw = "num")
    7:optional string startTime  (api.raw = "startTime")
    8:optional double price  (api.raw = "price")
    9:optional string remark  (api.raw = "remark")
    10:optional i64 coachId  (api.raw = "coachId")
    11:optional i64 memberId  (api.raw = "memberId")
    12:optional i64 memberProductId  (api.raw = "memberProductId")
    13:optional i64 memberProductPropertyId  (api.raw = "memberProductPropertyId")
}

struct ScheduleListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 member (api.raw = "member")
    4:  optional list<i64> coach (api.raw = "coach")
    5:  optional list<i64> product (api.raw = "product")
    6:  optional i64 venueId (api.raw = "venueId")
    7:  optional list<i64> property (api.raw = "property")
    8:  optional string startTime  (api.raw = "startTime")
    9:  optional string type  (api.raw = "type")
    10: optional string place  (api.raw = "place")
    11: optional string endTime  (api.raw = "endTime")
}

struct ScheduleMemberListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 member (api.raw = "member")
    4:  optional i64 schedule (api.raw = "schedule")
    5:  optional string type  (api.raw = "type")
}

struct ScheduleCoachListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 coach (api.raw = "coach")
    4:  optional i64 schedule (api.raw = "schedule")
    5:  optional string type  (api.raw = "type")
}

struct SearchSubscribeByMemberReq{
    1:  optional i64	propertyId  (api.raw = "propertyId")
    2:  optional string	mobile      (api.raw = "mobile")
    3:  optional i64	venue       (api.raw = "venue")
    4:  optional string	Type        (api.raw = "type")
}

struct CreateMemberReq{
    1:  optional list<i64> memberProductPropertyId (api.raw = "memberProductPropertyId")
    2:  optional i64 schedule (api.raw = "schedule")
    3:  optional string remark (api.raw = "remark")
}

struct ScheduleInfo  {
	1:i64 Id  (api.raw = "id")

	2:string Type      (api.raw = "type")
	3:i64 PropertyId   (api.raw = "property_id")
	4:i64 VenueId      (api.raw = "venue_id")
	5:i64 PlaceID      (api.raw = "place_id")

	6:i64 Num          (api.raw = "num")
	7:i64 NumSurplus                    (api.raw = "num_surplus")
	8:string Date                       (api.raw = "data")
	9:string StartTime                  (api.raw = "start_time")
	10:string EndTime                   (api.raw = "end_time")
	11:double Price                     (api.raw = "price")
	12:string Name                      (api.raw = "name")
	13:string Remark                    (api.raw = "remark")
	14:i64 CoachID                      (api.raw = "coach_id")
	15:i64 MemberID                     (api.raw = "member_id")
	16:i64 MemberProductID              (api.raw = "member_product_id")
	17:i64 MemberProductPropertyID      (api.raw = "member_product_property_id")
	18:i64 Status                       (api.raw = "status")
	19:string PropertyName                (api.raw = "property_name")
	20:string VenueName                   (api.raw = "venue_name")
	21:string PlaceName                   (api.raw = "place_name")
	22:string CoachName                   (api.raw = "coach_name")
	23:string MemberName                  (api.raw = "member_name")
	24:string MemberProductName           (api.raw = "member_product_name")
	25:string MemberProductPropertyName   (api.raw = "member_product_property_name")

	26:list<ScheduleMemberInfo> ScheduleMember (api.raw = "member_course_record")
	27:list<ScheduleCoachInfo> ScheduleCoach  (api.raw = "coach_course_record")

	28:string CreatedAt  (api.raw = "created_at")
	29:string UpdatedAt  (api.raw = "updated_at")
}

struct ScheduleMemberInfo  {
	1:i64 Id                    (api.raw = "id")
	2:i64 MemberId              (api.raw = "memberId")
	3:i64 VenueId               (api.raw = "venueId")
	4:i64 PlaceID               (api.raw = "placeId")
	5:i64 PropertyId            (api.raw = "propertyId")
	6:i64 ScheduleId            (api.raw = "scheduleId")
	7:string ScheduleName         (api.raw = "scheduleName")
	8:string Type                 (api.raw = "type")
	9:string CreatedAt            (api.raw = "createdAt")
	10:string UpdatedAt            (api.raw = "updatedAt")
	11:string StartTime            (api.raw = "startTime")
	12:string EndTime              (api.raw = "endTime")
	13:string SignStartTime        (api.raw = "signStartTime")
	14:string SignEndTime          (api.raw = "signEndTime")
	15:i64 Status                (api.raw = "status")
	16:i64 MemberProductId       (api.raw = "memberProductId")
	17:i64 MemberProductItemId   (api.raw = "memberProductItemId")

	18:string VenueName                  (api.raw = "venueName")
	19:string MemberName                 (api.raw = "memberName")
	20:string MemberProductName          (api.raw = "memberProductName")
	21:string MemberProductPropertyName  (api.raw = "memberProductPropertyName")
	22:string Gender                     (api.raw = "gender")
	23:i64 Birthday                    (api.raw = "birthday")
	24:string Mobile                     (api.raw = "mobile")
}

struct ScheduleCoachInfo  {
	1:i64 ID              (api.raw = "id")
	2:i64 CoachId         (api.raw = "coach_id")
	3:i64 VenueId         (api.raw = "venue_id")
	4:i64 PlaceID         (api.raw = "place_id")
	5:i64 PropertyId      (api.raw = "property_id")
	6:i64 ScheduleId      (api.raw = "schedule_id")
	7:string Type           (api.raw = "type")
	8:string CreatedAt      (api.raw = "createdAt")
	9:string UpdatedAt      (api.raw = "updatedAt")
	10:string Date           (api.raw = "date")
	11:string StartTime      (api.raw = "start_time")
	12:string EndTime        (api.raw = "end_time")
	13:string SignStartTime  (api.raw = "sign_start_time")
	14:string SignEndTime    (api.raw = "sign_end_time")
	15:i64 Status          (api.raw = "status")

	16:string ScheduleName  (api.raw = "schedule_name")
	17:string PropertyName  (api.raw = "property_name")
	18:string VenueName     (api.raw = "venue_name")
	19:string PlaceName     (api.raw = "place_name")
	20:string CoachName     (api.raw = "coach_name")
	21:string CoachAvatar   (api.raw = "coach_avatar")

	22:string Mobile                     (api.raw = "mobile")
	23:string MemberName                 (api.raw = "member_name")
	24:string MemberAvatar               (api.raw = "member_avatar")
	25:string MemberProductName          (api.raw = "member_product_name")
	26:string MemberProductPropertyName  (api.raw = "member_product_property_name")
	27:string Remark                     (api.raw = "remark")
	28:string MRemark                    (api.raw = "m_remark")
}

struct SubscribeByMember  {
	1:string Avatar                   (api.raw = "avatar")
	2:string Mobile                   (api.raw = "mobile")
	3:i64 MemberID                  (api.raw = "member_id")
	4:i64 MemberProductID           (api.raw = "member_product_id")
	5:i64 MemberProductPropertyId   (api.raw = "member_product_property_id")

	6:string MemberName                 (api.raw = "member_name")
	7:string MemberProductName          (api.raw = "member_product_name")
	8:string MemberProductPropertyName  (api.raw = "member_product_property_name")
}


struct ScheduleListResp {
    1: base.BaseResp resp
    2: optional list<ScheduleInfo> extra
}
struct ScheduleMemberListResp {
    1: base.BaseResp resp
    2: optional list<ScheduleMemberInfo> extra
}
struct ScheduleCoachListResp {
    1: base.BaseResp resp
    2: optional list<ScheduleCoachInfo> extra
}
struct ScheduleDateListResp {
  1: base.BaseResp resp
  2: optional map<string,list<ScheduleInfo>> extra
}
struct SearchSubscribeByMemberResp {
    1: base.BaseResp resp
    2: optional list<SubscribeByMember> extra
}
service ScheduleService {
    base.NilResponse CreateSchedule(1: CreateOrUpdateScheduleReq req)  (api.post = "/service/schedule/create")

    base.NilResponse UpdateSchedule(1: CreateOrUpdateScheduleReq req) (api.post = "/service/schedule/update")

    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/status")

    ScheduleListResp ScheduleList(1: ScheduleListReq req )(api.post = "/service/schedule/list")

    ScheduleDateListResp ScheduleDateList(1: ScheduleListReq req )(api.post = "/service/schedule/date-list")

    ScheduleInfo ScheduleInfo(1: base.IDReq req) (api.post = "/service/schedule/info")



    ScheduleMemberListResp MemberList(1: ScheduleMemberListReq req) (api.post = "/service/schedule/member-list")

    ScheduleMemberListResp CreateMember(1: CreateMemberReq req) (api.post = "/service/schedule/member-create")

    base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-member-status")

    SearchSubscribeByMemberResp SearchSubscribeByMember(1: SearchSubscribeByMemberReq req) (api.post = "/service/schedule/search-subscribe-by-member")

    ScheduleCoachListResp CoachList(1: ScheduleCoachListReq req) (api.post = "/service/schedule/coach-list")

    base.NilResponse UpdateCoachStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-coach-status")

}