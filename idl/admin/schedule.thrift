namespace go admin.schedule

include "../base/base.thrift"
include "../base/data.thrift"


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

struct ListScheduleReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 member (api.raw = "member")
    4:  optional list<i64> coach (api.raw = "coach")
    5:  optional list<i64> product (api.raw = "product")+
    6:  optional i64 venueId (api.raw = "venueId")
    7:  optional list<i64> property (api.raw = "property")
    8:  optional string startTime  (api.raw = "startTime")
    9:  optional string type  (api.raw = "type")
}
struct ScheduleMemberReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 member (api.raw = "member")
    4:  optional i64 schedule (api.raw = "schedule")
    5:  optional string type  (api.raw = "type")
}
struct ScheduleCoachReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 member (api.raw = "coach")
    4:  optional i64 schedule (api.raw = "schedule")
    5:  optional string type  (api.raw = "type")
}

struct SearchSubscribeByMemberReq{
    1:  optional i64	propertyId  (api.raw = "propertyId")
    2:  optional string	mobile      (api.raw = "mobile")
    3:  optional i64	venue       (api.raw = "venue")
}
struct MemberSubscribeReq{
    1:  optional list<i64> memberProductPropertyId (api.raw = "memberProductPropertyId")
    2:  optional i64 schedule (api.raw = "schedule")
    3:  optional string remark (api.raw = "remark")
}

service ScheduleService {
    base.NilResponse CreateSchedule(1: CreateOrUpdateScheduleReq req)  (api.post = "/api/admin/schedule/create")

    base.NilResponse UpdateSchedule(1: CreateOrUpdateScheduleReq req) (api.post = "/api/admin/schedule/update")

    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/schedule/status")

    base.NilResponse ListSchedule(1: ListScheduleReq req )(api.post = "/api/admin/schedule/list")

    base.NilResponse DateListSchedule(1: ListScheduleReq req )(api.post = "/api/admin/schedule/date-list")

    base.NilResponse GetScheduleById(1: base.IDReq req) (api.post = "/api/admin/schedule/info")

    base.NilResponse GetScheduleMemberList(1: ScheduleMemberReq req) (api.post = "/api/admin/schedule/schedule-member-list")

    base.NilResponse SearchSubscribeByMember(1: SearchSubscribeByMemberReq req) (api.post = "/api/admin/schedule/search-subscribe-by-member")

    base.NilResponse MemberSubscribe(1: MemberSubscribeReq req) (api.post = "/api/admin/schedule/member-subscribe")

    base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/schedule/schedule-member-status")


    base.NilResponse GetScheduleCoachList(1: ScheduleMemberReq req) (api.post = "/api/admin/schedule/schedule-coach-list")

    base.NilResponse UpdateCoachStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/schedule/schedule-coach-status")

}