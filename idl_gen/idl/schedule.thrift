namespace go schedule

include "../base/base.thrift"

struct CreateOrUpdateScheduleReq {
    1:optional i64 id=0  (api.raw = "id")
    2:optional string type=""  (api.raw = "type")
    3:optional i64 propertyId=0   (api.raw = "propertyId")
    4:optional i64 venueId =0  (api.raw = "venueId")
    5:optional i64 placeId =0  (api.raw = "placeId")
    6:optional i64 num=0   (api.raw = "num")
    7:optional string startTime =""  (api.raw = "startTime")
    8:optional double price=0   (api.raw = "price")
    9:optional string remark =""   (api.raw = "remark")
    10:optional i64 coachId =0  (api.raw = "coachId")
    11:optional i64 memberId =0  (api.raw = "memberId")
    12:optional i64 memberProductId =0  (api.raw = "memberProductId")
}

struct ListScheduleReq {
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize = 100 (api.raw = "pageSize")
    3:  optional i64 member = 0 (api.raw = "member")
    4:  optional list<i64> coach = 0 (api.raw = "coach")
    5:  optional list<i64> product = 0 (api.raw = "product")+
    6:  optional i64 venueId = 0 (api.raw = "venueId")
    7:  optional list<i64> property = 0 (api.raw = "property")
    8:  optional string startTime =""  (api.raw = "startTime")
    9:  optional string type =""  (api.raw = "type")
}
struct ScheduleMemberReq {
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize= 100 (api.raw = "pageSize")
    3:  optional i64 memberId = 0 (api.raw = "memberId")
    4:  optional i64 scheduleId = 0(api.raw = "scheduleId")
    5:  optional string type =""  (api.raw = "type")
}
struct ScheduleCoachReq{
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize = 100 (api.raw = "pageSize")
    3:  optional i64 coachId = 0 (api.raw = "coachId")
    4:  optional i64 scheduleId = 0(api.raw = "scheduleId")
    5:  optional string type=""  (api.raw = "type")
}

struct SearchSubscribeByMemberReq{
    1:  optional i64	propertyId=0  (api.raw = "propertyId")
    2:  optional string	mobile  ="" (api.raw = "mobile")
    3:  optional i64	venue  = 0  (api.raw = "venue")
}
struct MemberSubscribeReq{
    1:  optional list<i64> memberProductId =0 (api.raw = "memberProductId")
    2:  optional i64 scheduleId =0 (api.raw = "scheduleId")
    3:  optional string remark ="" (api.raw = "remark")
}

service ScheduleService {
    base.NilResponse CreateSchedule(1: CreateOrUpdateScheduleReq req)  (api.post = "/service/schedule/create")

    base.NilResponse UpdateSchedule(1: CreateOrUpdateScheduleReq req) (api.post = "/service/schedule/update")

    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/status")

    base.NilResponse ListSchedule(1: ListScheduleReq req )(api.post = "/service/schedule/list")

    base.NilResponse DateListSchedule(1: ListScheduleReq req )(api.post = "/service/schedule/date-list")

    base.NilResponse GetScheduleById(1: base.IDReq req) (api.post = "/service/schedule/info")

    base.NilResponse GetScheduleMemberList(1: ScheduleMemberReq req) (api.post = "/service/schedule/schedule-member-list")

    base.NilResponse SearchSubscribeByMember(1: SearchSubscribeByMemberReq req) (api.post = "/service/schedule/search-subscribe-by-member")

    base.NilResponse MemberSubscribe(1: MemberSubscribeReq req) (api.post = "/service/schedule/member-subscribe")

    base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-member-status")


    base.NilResponse GetScheduleCoachList(1: ScheduleMemberReq req) (api.post = "/service/schedule/schedule-coach-list")

    base.NilResponse UpdateCoachStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-coach-status")

}