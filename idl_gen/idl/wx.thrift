namespace go wx

include "../base/base.thrift"
include "member_product.thrift"
include "member.thrift"
include "schedule.thrift"
include "order.thrift"
include "product.thrift"
include "contest.thrift"
include "community.thrift"
include "venue.thrift"
include "bootcamp.thrift"
/**
 * 微信小程序服务
 */

struct MemberRegisterReq {
    1:  optional string mobile="" (api.raw = "mobile")
    2:  optional string captcha="" (api.raw = "captcha")
    4:  optional i64 venueId=0 (api.raw = "venueId")
}
struct MemberCaptchaReq {
    1:  optional string mobile="" (api.raw = "mobile")
    4:  optional i64 venueId=0 (api.raw = "venueId")
}
struct MemberLoginReq {
    1:  optional string mobile="" (api.raw = "mobile")
    2:  optional string captcha="" (api.raw = "captcha")
    4:  optional i64 venueId=0 (api.raw = "venueId")
}

service WxService {

   /**会员注册*/
   base.NilResponse MemberRegister(1: MemberRegisterReq req) (api.post = "/service/wx/member/register")
   /**会员验证码*/
   base.NilResponse MemberCaptcha(1: MemberCaptchaReq req) (api.post = "/service/wx/member/captcha")
   /**会员登录*/
   base.NilResponse MemberLogin(1: MemberLoginReq req) (api.post = "/service/wx/member/login")
   /**会员登出*/
   base.NilResponse MemberLogout(1: base.IDReq req) (api.post = "/service/wx/member/logout")
   /**会员详情*/
   base.NilResponse MemberInfo(1: base.IDReq req) (api.post = "/service/wx/member/info")
   /**会员产品详情*/
   base.NilResponse MemberProductInfo(1: base.IDReq req) (api.post = "/service/wx/member/product-info")
   /**会员产品列表*/
   base.NilResponse MemberProductList(1:  member_product.MemberProductListReq req) (api.post = "/service/wx/member/product-list")
   /**会员参赛列表*/
   base.NilResponse MemberContestList (1: member.MemberContestListReq req) (api.post = "/service/wx/member/contest-list")
   /**会员训练营报名列表*/
   base.NilResponse MemberBootcampList (1: member.MemberBootcampListReq req) (api.post = "/service/wx/member/bootcamp-list")
   /**会员社群报名列表*/
   base.NilResponse MemberCommunityList (1: member.MemberCommunityListReq req) (api.post = "/service/wx/member/community-list")
   /**会员订单列表*/
   base.NilResponse MemberOrderList (1: MemberOrderListReq req) (api.post = "/service/wx/member/order-list")
   /**会员合同列表*/
   base.NilResponse MemberContractList(1: member.MemberContractListReq req) (api.post = "/service/wx/member/contract-list")
   /**会员课程列表*/
   base.NilResponse ScheduleMemberList(1: schedule.ScheduleMemberListReq req) (api.post = "/service/wx/member/schedule-member-list")



   /**比赛列表*/
   base.NilResponse ContestList (1: contest.ContestListReq req) (api.post = "/service/wx/member/contest-list")
   /**比赛详情*/
   base.NilResponse ContestInfo (1: base.IDReq req) (api.post = "/service/wx/member/contest-info")
   /**比赛报名*/
   base.NilResponse JoinContest (1: JoinContestReq req) (api.post = "/service/wx/member/join-contest")


   /**训练营列表*/
   base.NilResponse BootcampList (1: bootcamp.BootcampListReq req) (api.post = "/service/wx/member/bootcamp-list")
   /**训练营详情*/
   base.NilResponse BootcampInfo (1: base.IDReq req) (api.post = "/service/wx/member/bootcamp-info")
   /**训练营报名*/
   base.NilResponse JoinBootcamp (1: JoinBootcampReq req) (api.post = "/service/wx/member/join-bootcamp")

   /**社群列表*/
   base.NilResponse CommunityList (1: community.CommunityListReq req) (api.post = "/service/wx/member/community-list")
   /**社群详情*/
   base.NilResponse CommunityInfo (1: base.IDReq req) (api.post = "/service/wx/member/community-info")
   /**社群报名*/
   base.NilResponse JoinCommunity (1: JoinCommunityReq req) (api.post = "/service/wx/member/join-community")

   /**场馆列表*/
   base.NilResponse VenueList (1: venue.VenueListReq req) (api.post = "/service/wx/member/venue-list")
   /**场馆详情*/
   base.NilResponse VenueInfo (1: base.IDReq req) (api.post = "/service/wx/member/venue-info")

   /**扫码入场*/
//   base.NilResponse ScanEnterQR (1: ScanQREnterReq req) (api.post = "/service/wx/scan-QR-Enter")
      base.NilResponse ScanQR (1: ScanQRReq req) (api.post = "/service/wx/member/scan-QR")
   /**激活*/
   base.NilResponse Activation (1: base.IDReq req) (api.post = "/service/wx/member/activation")


   /**会员私教课约课*/
   base.NilResponse CreateMemberScheduleCourse(1: schedule.CreateOrUpdateScheduleCourseReq req) (api.post = "/service/wx/member/create-member-schedule-course")
    /**排团教课*/
    base.NilResponse CreateMemberScheduleCourseLessons(1: CreateMemberScheduleLessonsReq req) (api.post = "/service/wx/member/create-member-schedule-lessons")

   /**产品详情*/
   base.NilResponse ProductInfo(1: base.IDReq req) (api.post = "/service/wx/member/product-info")
   /**产品列表*/
   base.NilResponse ProductList(1: product.ProductListReq req) (api.post = "/service/wx/member/product-list")
   /**下单*/
   base.NilResponse Buy(1: order.BuyReq req) (api.post = "/service/wx/member/buy")

   /**教练列表*/
   base.NilResponse CoachList(1: CoachListReq req) (api.post = "/service/wx/member/coach-list")
   /**教练详情*/
   base.NilResponse CoachInfo(1: base.IDReq req) (api.post = "/service/wx/member/coach-info")

   /**场地列表*/
   base.NilResponse PlaceList(1: venue.VenuePlaceListReq req) (api.post = "/service/wx/member/place-list")
   /**场地详情*/
   base.NilResponse PlaceInfo(1: base.IDReq req) (api.post = "/service/wx/member/place-info")
   /**场地预约*/
   base.NilResponse CreatePlaceSchedule(1: CreatePlaceScheduleReq req) (api.post = "/service/wx/member/create-place-schedule")

   /**我的会员*/
   base.NilResponse MyMember(1: MyMemberReq req) (api.post = "/service/wx/staff/my-member")
   /**课程列表*/
   base.NilResponse CoachScheduleList(1: schedule.ScheduleCoachListReq req) (api.post = "/service/wx/staff/schedule-member-list")
   /**会员签到*/
   base.NilResponse SignMemberSchedule(1: SignMemberScheduleReq req) (api.post = "/service/wx/staff/sign-member-schedule")
   /**教练签到*/
   base.NilResponse SignStaffSchedule(1: SignStaffScheduleReq req) (api.post = "/service/wx/staff/sign-staff-schedule")
   /**教练验证码*/
   base.NilResponse StaffCaptcha(1: StaffCaptchaReq req) (api.post = "/service/wx/staff/captcha")
   /**教练登录*/
   base.NilResponse StaffLogin(1: StaffLoginReq req) (api.post = "/service/wx/staff/login")
   /**会员登出*/
   base.NilResponse StaffLogout(1: base.IDReq req) (api.post = "/service/wx/staff/logout")
}

struct StaffCaptchaReq {
    1:  optional string mobile="" (api.raw = "mobile")
    2:  optional string captcha="" (api.raw = "captcha")
    4:  optional i64 venueId=0 (api.raw = "venueId")
}
struct StaffLoginReq {
    1:  optional string mobile="" (api.raw = "mobile")
    2:  optional string captcha="" (api.raw = "captcha")
    4:  optional i64 venueId=0 (api.raw = "venueId")
}

struct CreatePlaceScheduleReq {
    1:  optional i64 memberId=0 (api.raw = "memberId")
    2:  optional i64 placeId=0 (api.raw = "placeId")
}

struct CoachListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional i64 status=0 (api.raw = "status")
    4:  optional i64 venueId=0 (api.raw = "venueId")
    /**职能*/
    5: optional list<string> functions="" (api.raw = "functions")
    /**标签*/
    6: optional list<i64> tagId=0 (api.raw = "tagId" )
}

struct JoinContestReq {
    1:  optional i64 contestId=1 (api.raw = "contestId")
    2:  optional string fields="" (api.raw = "fields")
    3:  optional string name="" (api.raw = "name")
    4:  optional string mobile="" (api.raw = "mobile")
}

struct JoinBootcampReq{
    1:  optional i64 bootcampId=0 (api.raw = "bootcampId")
    2:  optional string fields="" (api.raw = "fields")
    3:  optional string name="" (api.raw = "name")
    4:  optional string mobile="" (api.raw = "mobile")
}

struct JoinCommunityReq{
    1:  optional i64 communityId=0 (api.raw = "communityId")
    2:  optional string fields="" (api.raw = "fields")
    3:  optional string name="" (api.raw = "name")
}

struct ScanQRReq{
    1:  optional i64 memberId=0 (api.raw = "memberId")
    2:  optional i64 venueId=0 (api.raw = "venueId")
}

struct CreateMemberScheduleLessonsReq{
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 scheduleId=0 (api.raw = "scheduleId")
    3:optional base.Seat seats  (api.raw = "seat")
    4:optional i64 memberProductId    (api.raw = "memberProductId")
}

struct ProductListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional i64 status=0 (api.raw = "status")
    4:  optional i64 venueId=0 (api.raw = "venueId")
}
struct MemberOrderListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    6: optional list<i64> venueId=0 (api.raw = "venueId")
    7: optional list<i64> status=0 (api.raw = "status")
    /**订单完成时间*/
    9: optional string startCompletionAt="" (api.raw = "startCompletionAt")
    10: optional string endCompletionAt="" (api.raw = "endCompletionAt")
    /**产品类型*/
    11: optional string productType="" (api.raw = "productType")
    /**订单业务*/
    12: optional string nature="" (api.raw = "nature")
    /**产品名称*/
    13: optional string name ="" (api.raw = "name ")
    14: optional string memberName="" (api.raw = "memberName")
    15: optional i64 memberId=0 (api.raw = "memberId")
}

struct MyMemberReq{
  1:optional i64 staffId =0  (api.raw = "staffId")
  2:optional i64 venueId =0  (api.raw = "venueId")
}
struct SignMemberScheduleReq{
  1:optional i64 staffId =0  (api.raw = "staffId")
  2:optional i64 memberId =0  (api.raw = "memberId")
  3:optional i64 scheduleId =0  (api.raw = "scheduleId")
}
struct SignStaffScheduleReq{
  1:optional i64 staffId =0  (api.raw = "staffId")
  2:optional i64 scheduleId =0  (api.raw = "scheduleId")
}