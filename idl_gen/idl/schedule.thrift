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
//	13:string remark                    (api.raw = "remark")
	14:i64 coachId                      (api.raw = "coachId")
//	15:i64 memberId                     (api.raw = "memberId")
//	16:i64 memberProductId              (api.raw = "memberProductId")
    /**状态 是 1未发布 2发布 3取消*/
	18:i64 status                       (api.raw = "status")
    19:i64 productId (api.raw = "productId")
	20:string venueName                   (api.raw = "venueName")
	21:string placeName                   (api.raw = "placeName")
	22:string coachName                   (api.raw = "coachName")
//	23:string memberName                  (api.raw = "memberName")
//	24:string memberProductName           (api.raw = "memberProductName")

    25:list<list<base.Seat>> seats  (api.raw = "seats")
	26:list<ScheduleMemberInfo> memberCourseRecord (api.raw = "memberCourseRecord")
	27:list<ScheduleCoachInfo> coachCourseRecord  (api.raw = "coachCourseRecord")

	28:string createdAt  (api.raw = "createdAt")
	29:string updatedAt  (api.raw = "updatedAt")

    }
struct ScheduleMemberInfo {
	1:i64 id                    (api.raw = "id")
	2:i64 memberId              (api.raw = "memberId")
	3:i64 venueId               (api.raw = "venueId")
	4:i64 placeId               (api.raw = "placeId")
  	5:string date           (api.raw = "date")
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
	17:string placeName                  (api.raw = "placeName")
	18:string venueName                  (api.raw = "venueName")
	19:string memberName                 (api.raw = "memberName")
	20:string memberProductName          (api.raw = "memberProductName")
	21:string  memberProductSn          (api.raw = "memberProductSn")
    22:i64 coachId         (api.raw = "coachId")
	23:string coachName     (api.raw = "coachName")
	24:string mobile                     (api.raw = "mobile")
	25: base.Seat seat  (api.raw = "seat")

	26:i64 productId      (api.raw = "productId")


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

	11:string startTime      (api.raw = "startTime")
	12:string endTime        (api.raw = "endTime")
	13:string signStartTime  (api.raw = "signStartTime")
	14:string signEndTime    (api.raw = "signEndTime")
	15:i64 status          (api.raw = "status")

	16:string scheduleName  (api.raw = "scheduleName")
	20:string coachName     (api.raw = "coachName")
//	21:string coachAvatar   (api.raw = "coachAvatar")
 	10:string date           (api.raw = "date")
	18:string venueName     (api.raw = "venueName")
	19:string placeName     (api.raw = "placeName")

	21:i64 memberId      (api.raw = "memberId")
	22:string memberMobile                     (api.raw = "memberMobile")
	23:string memberName                 (api.raw = "memberName")
//	24:string memberAvatar               (api.raw = "memberAvatar")
	25:string memberProductName          (api.raw = "memberProductName")
	26:i64 memberProductId      (api.raw = "memberProductId")
//	27:string remark                     (api.raw = "remark")
//	28:string mRemark                    (api.raw = "mRemark")
}
struct CreateOrUpdateScheduleCourseReq {
    1:optional i64 id=0  (api.raw = "id")
    /**类别[一对一:courseOne;一对多:courseMore]*/
    2:optional string type=""  (api.raw = "type")
     /**场馆*/
    3:optional i64 venueId =0  (api.raw = "venueId")
    /**会员ID*/
    4:optional i64 memberId =0  (api.raw = "memberId")
     /**会员IDs*/
    5:optional list<i64> memberIds =0  (api.raw = "memberIds")
    /**开始时间*/
    6:optional string startTime =""  (api.raw = "startTime")
    /**教练ID*/
    7:optional i64 coachId =0  (api.raw = "coachId")
    /**产品ID*/
    8:optional i64 productId =0  (api.raw = "productId")
    /**会员产品ID*/
    9:optional i64 memberProductId =0  (api.raw = "memberProductId")
}
struct CreateOrUpdateScheduleLessonsReq {
    /**产品*/
    1:optional i64 productId =0  (api.raw = "productId")
     /**上课时间*/
    2:optional string startTime =""  (api.raw = "startTime")
     /**教练*/
    3:optional i64 coachId =0  (api.raw = "coachId")
     /**场馆*/
    4:optional i64 venueId =0  (api.raw = "venueId")
    /**教室*/
    5:optional i64 placeId =0  (api.raw = "placeId")

    255:optional i64 id=0  (api.raw = "id")
}



struct ScheduleListReq {
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize = 100 (api.raw = "pageSize")
    3:  optional i64 member = 0 (api.raw = "member")
    4:  optional list<i64> coach = 0 (api.raw = "coach")
    5:  optional list<i64> product = 0 (api.raw = "product")
    6:  optional i64 venueId = 0 (api.raw = "venueId")
    7:  optional list<i64> status = 0 (api.raw = "status")
    8:  optional string startTime =""  (api.raw = "startTime")
    9:  optional string type =""  (api.raw = "type")
}
struct ScheduleMemberListReq {
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize= 100 (api.raw = "pageSize")
    3:  optional i64 memberId = 0 (api.raw = "memberId")
    4:  optional i64 scheduleId = 0(api.raw = "scheduleId")
    5:  optional string type =""  (api.raw = "type")
}
struct ScheduleCoachListReq{
    1:  optional i64 page = 1 (api.raw = "page")
    2:  optional i64 pageSize = 100 (api.raw = "pageSize")
    3:  optional i64 coachId = 0 (api.raw = "coachId")
    4:  optional i64 scheduleId = 0(api.raw = "scheduleId")
    5:  optional string type=""  (api.raw = "type")
}

//struct SearchSubscribeByMemberReq{
//    1:  optional i64	memberProductId=0  (api.raw = "memberProductId")
//    2:  optional string	mobile  ="" (api.raw = "mobile")
//    3:  optional i64	venue  = 0  (api.raw = "venue")
//}
struct MemberSubscribeReq{
    1:  optional i64 memberProductId =0 (api.raw = "memberProductId")
    2:  optional i64 scheduleId =0 (api.raw = "scheduleId")
//    3:  optional string remark ="" (api.raw = "remark")
    4:  optional i64 memberId =0 (api.raw = "memberId")
    5:  optional base.Seat seat ={} (api.raw = "seat")
}

struct UserTimePeriodReq{
    /**时间如2024-05-16*/
    1:  optional string startDate ="" (api.raw = "startDate")
    2:  optional string endDate ="" (api.raw = "endDate")
    /**时间段*/
    3:  optional base.Period period = {} (api.raw = "period")
    4:  optional list<i64> userId = {} (api.raw = "userId")
    5:  optional i64 venueId = 0 (api.raw = "venueId")
}
struct UserTimePeriodInfo{
    2:  optional string date ="" (api.raw = "date")
    3:  optional base.Period period = {} (api.raw = "period")
    4:  optional i64 userId = 0 (api.raw = "userId")
}
struct UserPeriodReq{
    /**员工ID*/
    1: optional i64 id = 0 (api.raw = "id")
    /**时间*/
    2: optional string date = "" (api.raw = "date")
    3: optional i64 venueId = 0 (api.raw = "venueId")
}
struct UpdateUserTimePeriodReq{
    /**员工ID*/
    1: optional i64 userId = 0 (api.raw = "userId")
    /**时间*/
    2: optional string date = "" (api.raw = "date")
    3: optional i64 venueId = 0 (api.raw = "venueId")
    4:  optional base.Period period = {} (api.raw = "period")
}

struct ScheduleCoachPeriod{
    1:optional string date = "" (api.raw = "date")
	2:optional i64 coachId = 0 (api.raw = "coachId")
	3:optional i64 venueId = 0 (api.raw = "venueId")
	4:optional string coachName = "" (api.raw = "coachName")
    5:optional list<base.List> tags = {}  (api.raw = "tags")
	6:optional base.Period period = {} (api.raw = "period")
	7:optional list<ScheduleCoachInfo> scheduleCoachList = {}  (api.raw = "scheduleCoachList")
}

service ScheduleService {
    /**添加教练时间段*/
    base.NilResponse CreateScheduleUserTimePeriod(1: UserTimePeriodReq req)  (api.post = "/service/schedule/create-user-time-period")
    /**更新教练时间段*/
    base.NilResponse UpdateScheduleUserTimePeriod(1: UpdateUserTimePeriodReq req)  (api.post = "/service/schedule/update-user-time-period")
    /**教练时间段*/
    base.NilResponse UserTimePeriod(1: UserPeriodReq req)  (api.post = "/service/schedule/user-time-period")

    /**教练课程时间段*/
    base.NilResponse ScheduleCoachPeriodList(1: UserPeriodReq req)  (api.post = "/service/schedule/coach-period-list")

    /**约私教课*/
    base.NilResponse CreateScheduleCourse(1: CreateOrUpdateScheduleCourseReq req)  (api.post = "/service/schedule/create-cours")
    /**排团教课*/
    base.NilResponse CreateScheduleLessons(1: CreateOrUpdateScheduleLessonsReq req) (api.post = "/service/schedule/create-lessons")
   /**团教课发布*/
    base.NilResponse ScheduleLessonsPublish(1: base.Ids req) (api.post = "/service/schedule/lessons-publish")

    // base.NilResponse UpdateScheduleLessons(1: CreateOrUpdateScheduleLessonsReq req) (api.post = "/service/schedule/update-lessons")
    base.NilResponse UpdateScheduleStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/status")
    /**课程列表*/
    base.NilResponse ScheduleList(1: ScheduleListReq req )(api.post = "/service/schedule/list")
    /**按时间-课程列表*/
    base.NilResponse ScheduleDateList(1: ScheduleListReq req )(api.post = "/service/schedule/date-list")

    base.NilResponse ScheduleInfo(1: base.IDReq req) (api.post = "/service/schedule/info")
    /**会员约团教课*/
    base.NilResponse CreateMemberSubscribeLessons(1: MemberSubscribeReq req) (api.post = "/service/schedule/create-member-subscribe-lessons")

    /**会员课程列表*/
    base.NilResponse ScheduleMemberList(1: ScheduleMemberListReq req) (api.post = "/service/schedule/schedule-member-list")
    //base.NilResponse SearchSubscribeByMember(1: SearchSubscribeByMemberReq req) (api.post = "/service/schedule/search-subscribe-by-member")
    base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-member-status")
    base.NilResponse ScheduleMemberInfo(1: base.IDReq req) (api.post = "/service/schedule/schedule-member-info")

    /**教练课程列表*/
    base.NilResponse ScheduleCoachList(1: ScheduleCoachListReq req) (api.post = "/service/schedule/schedule-coach-list")
    base.NilResponse UpdateScheduleCoachStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-coach-status")
    base.NilResponse ScheduleCoachInfo(1: base.IDReq req) (api.post = "/service/schedule/schedule-coach-info")



}