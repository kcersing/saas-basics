namespace go schedule

include "../base/base.thrift"
struct ScheduleInfo {
	1:i64 id  (api.raw = "id")

	2:string type      (api.raw = "type")

	4:i64 venueId      (api.raw = "venueId")
	5:i64 placeId      (api.raw = "placeId")

	6:i64 num          (api.raw = "num")
	7:i64 numSurplus                    (api.raw = "numSurplus")
	8:string date                       (api.raw = "data")
	9:string startTime                  (api.raw = "startTime")
	10:string endTime                   (api.raw = "endTime")
	11:double price                     (api.raw = "price")
	12:string name                      (api.raw = "name")
	13:string remark                    (api.raw = "remark")
	14:i64 coachId                      (api.raw = "coachId")
	15:i64 memberId                     (api.raw = "memberId")
	16:i64 memberProductId              (api.raw = "memberProductId")

	18:i64 status                       (api.raw = "status")

	20:string venueName                   (api.raw = "venueName")
	21:string placeName                   (api.raw = "placeName")
	22:string coachName                   (api.raw = "coachName")
	23:string memberName                  (api.raw = "memberName")
	24:string memberProductName           (api.raw = "memberProductName")

	26:list<ScheduleMemberInfo> memberCourseRecord (api.raw = "memberCourseRecord")
	27:list<ScheduleCoachInfo> coachCourseRecord  (api.raw = "coachCourseRecord")

	28:string createdAt  (api.raw = "createdAt")
	29:string updatedAt  (api.raw = "updatedAt")
    }
struct ScheduleMemberInfo {
	1:i64 id                    (api.raw = "id")
	2:i64 memberId              (api.raw = "memberId")
	3:i64 venueId               (api.raw = "venueId")
	4:i64 placeID               (api.raw = "placeId")
	6:i64 scheduleId            (api.raw = "scheduleId")
	7:string scheduleName         (api.raw = "scheduleName")
	8:string type                 (api.raw = "type")
	9:string createdAt            (api.raw = "createdAt")
	10:string updatedAt            (api.raw = "updatedAt")
	11:string startTime            (api.raw = "startTime")
	12:string endTime              (api.raw = "endTime")
	13:string signStartTime        (api.raw = "signStartTime")
	14:string signEndTime          (api.raw = "signEndTime")
	15:i64 status                (api.raw = "status")
	16:i64 memberProductId       (api.raw = "memberProductId")
	17:i64 memberProductItemId   (api.raw = "memberProductItemId")

	18:string venueName                  (api.raw = "venueName")
	19:string memberName                 (api.raw = "memberName")
	20:string memberProductName          (api.raw = "memberProductName")
	22:string gender                     (api.raw = "gender")
	23:i64 birthday                    (api.raw = "birthday")
	24:string mobile                     (api.raw = "mobile")
}
struct ScheduleCoachInfo{
	1:i64 id              (api.raw = "id")
	2:i64 coachId         (api.raw = "coachId")
	3:i64 venueId         (api.raw = "venueId")
	4:i64 placeId         (api.raw = "placeId")
	5:i64 productId      (api.raw = "productId")
	6:i64 scheduleId      (api.raw = "scheduleId")
	7:string type           (api.raw = "type")
	8:string createdAt      (api.raw = "createdAt")
	9:string updatedAt      (api.raw = "updatedAt")
	10:string date           (api.raw = "date")
	11:string startTime      (api.raw = "startTime")
	12:string endTime        (api.raw = "endTime")
	13:string signStartTime  (api.raw = "signStartTime")
	14:string signEndTime    (api.raw = "signEndTime")
	15:i64 status          (api.raw = "status")

	16:string scheduleName  (api.raw = "scheduleName")
	17:string productName  (api.raw = "productName")
	18:string venueName     (api.raw = "venueName")
	19:string placeName     (api.raw = "placeName")
	20:string coachName     (api.raw = "coachName")
	21:string coachAvatar   (api.raw = "coachAvatar")

	22:string mobile                     (api.raw = "mobile")
	23:string memberName                 (api.raw = "memberName")
	24:string memberAvatar               (api.raw = "memberAvatar")
	25:string memberProductName          (api.raw = "memberProductName")
	27:string remark                     (api.raw = "remark")
	28:string mRemark                    (api.raw = "mRemark")
}
struct CreateOrUpdateScheduleReq {
    1:optional i64 id=0  (api.raw = "id")
    /**类别[一对一:1;一对多:2团课:3]*/
    2:optional i64 type=1  (api.raw = "type")
    4:optional i64 venueId =0  (api.raw = "venueId")
    /**场地*/
    5:optional i64 placeId =0  (api.raw = "placeId")
    /**人数*/
    6:optional i64 num=0   (api.raw = "num")
    /**开始时间*/
    7:optional string startTime =""  (api.raw = "startTime")
    /**价格*/
    8:optional double price=0   (api.raw = "price")
    /**备注*/
    9:optional string remark =""   (api.raw = "remark")
    /**教练ID*/
    10:optional i64 coachId =0  (api.raw = "coachId")
    /**产品ID*/
    11:optional i64 productId =0  (api.raw = "productId")
     /**会员ID*/
    13:optional i64 memberId =0  (api.raw = "memberId")
    /**会员产品ID*/
    14:optional i64 memberProductId =0  (api.raw = "memberProductId")
}


struct ScheduleListReq {
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize = 100 (api.raw = "pageSize")
    3:  optional i64 member = 0 (api.raw = "member")
    4:  optional list<i64> coach = 0 (api.raw = "coach")
    5:  optional list<i64> product = 0 (api.raw = "product")
    6:  optional i64 venueId = 0 (api.raw = "venueId")
    8:  optional string startTime =""  (api.raw = "startTime")
    9:  optional i64 type =1  (api.raw = "type")
}
struct ScheduleMemberListReq {
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize= 100 (api.raw = "pageSize")
    3:  optional i64 memberId = 0 (api.raw = "memberId")
    4:  optional i64 scheduleId = 0(api.raw = "scheduleId")
    5:  optional i64 type =1  (api.raw = "type")
}
struct ScheduleCoachListReq{
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize = 100 (api.raw = "pageSize")
    3:  optional i64 coachId = 0 (api.raw = "coachId")
    4:  optional i64 scheduleId = 0(api.raw = "scheduleId")
    5:  optional i64 type=1  (api.raw = "type")
}

//struct SearchSubscribeByMemberReq{
//    1:  optional i64	memberProductId=0  (api.raw = "memberProductId")
//    2:  optional string	mobile  ="" (api.raw = "mobile")
//    3:  optional i64	venue  = 0  (api.raw = "venue")
//}
struct MemberSubscribeReq{
    1:  optional list<i64> memberProductId =0 (api.raw = "memberProductId")
    2:  optional i64 scheduleId =0 (api.raw = "scheduleId")
    3:  optional string remark ="" (api.raw = "remark")
}

struct UserTimePeriodReq{
    /**时间如2024-05-16*/
    1:  optional string date ="" (api.raw = "date")
    /**时间段*/
    2:  optional string period ="" (api.raw = "period")
    3:  optional i64 userId =0 (api.raw = "userId")
}

service ScheduleService {

    base.NilResponse UpdateScheduleUserTimePeriod(1: UserTimePeriodReq req)  (api.post = "/service/schedule/create-user-time-period")

    base.NilResponse CreateSchedule(1: CreateOrUpdateScheduleReq req)  (api.post = "/service/schedule/create")

    base.NilResponse UpdateSchedule(1: CreateOrUpdateScheduleReq req) (api.post = "/service/schedule/update")

    base.NilResponse UpdateScheduleStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/status")

    base.NilResponse ScheduleList(1: ScheduleListReq req )(api.post = "/service/schedule/list")

    base.NilResponse ScheduleDateList(1: ScheduleListReq req )(api.post = "/service/schedule/date-list")

    base.NilResponse GetScheduleInfo(1: base.IDReq req) (api.post = "/service/schedule/info")



    base.NilResponse  CreateMemberSubscribe(1: MemberSubscribeReq req) (api.post = "/service/schedule/create-member-subscribe")
    base.NilResponse GetScheduleMemberList(1: ScheduleMemberListReq req) (api.post = "/service/schedule/schedule-member-list")
//    base.NilResponse SearchSubscribeByMember(1: SearchSubscribeByMemberReq req) (api.post = "/service/schedule/search-subscribe-by-member")
    base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-member-status")
    base.NilResponse ScheduleMemberInfo(1: base.IDReq req) (api.post = "/service/schedule/schedule-member-info")


    base.NilResponse GetScheduleCoachList(1: ScheduleCoachListReq req) (api.post = "/service/schedule/schedule-coach-list")
    base.NilResponse UpdateCoachStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-coach-status")
    base.NilResponse ScheduleCoachInfo(1: base.IDReq req) (api.post = "/service/schedule/schedule-coach-info")



}