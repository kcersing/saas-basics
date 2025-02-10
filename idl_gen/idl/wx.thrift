namespace go wx

include "../base/base.thrift"
include "member_product.thrift"
include "member.thrift"
include "schedule.thrift"

service WxMemberService {

   /**会员注册*/
   base.NilResponse MemberRegister(1: MemberRegisterReq req) (api.post = "/service/wx/member-register")
   /**会员验证码*/
   base.NilResponse MemberCaptcha(1: MemberCaptchaReq req) (api.post = "/service/wx/member-captcha")
   /**会员登录*/
   base.NilResponse MemberLogin(1: MemberLoginReq req) (api.post = "/service/wx/member-login")
   /**会员登出*/
   base.NilResponse MemberLogout(1: base.IDReq req) (api.post = "/service/wx/member-logout")
   /**会员详情*/
   base.NilResponse MemberInfo(1: base.IDReq req) (api.post = "/service/wx/member-info")
   /**会员产品详情*/
   base.NilResponse MemberProductInfo(1: base.IDReq req) (api.post = "/service/wx/member-product-info")
   /**会员产品列表*/
   base.NilResponse MemberProductList(1:  member_product.MemberProductListReq req) (api.post = "/service/wx/member-product-list")
   /**会员参赛列表*/
   base.NilResponse MemberContestList (1: member.MemberContestListReq req) (api.post = "/service/wx/member-contest-list")
   /**会员训练营报名列表*/
   base.NilResponse MemberBootcampList (1: member.MemberBootcampListReq req) (api.post = "/service/wx/member-bootcamp-list")
   /**会员社群报名列表*/
   base.NilResponse MemberCommunityList (1: member.MemberCommunityListReq req) (api.post = "/service/wx/member-community-list")
   /**会员订单列表*/
   base.NilResponse MemberOrderList (1: member.MemberCommunityListReq req) (api.post = "/service/wx/member-order-list")
   /**会员合同列表*/
   base.NilResponse MemberContractList(1: member.MemberContractListReq req) (api.post = "/service/wx/member-contract-list")
   /**会员课程列表*/
   base.NilResponse ScheduleMemberList(1: schedule.ScheduleMemberListReq req) (api.post = "/service/wx/member-schedule-member-list")



   /**比赛列表*/
   base.NilResponse ContestList (1: ContestListReq req) (api.post = "/service/wx/contest-list")
   /**比赛详情*/
   base.NilResponse ContestInfo (1: base.IDReq req) (api.post = "/service/wx/contest-info")
   /**比赛报名*/
   base.NilResponse JoinContest (1: JoinContestReq req) (api.post = "/service/wx/join-contest")


   /**训练营列表*/
   base.NilResponse BootcampList (1: BootcampListReq req) (api.post = "/service/wx/bootcamp-list")
   /**训练营详情*/
   base.NilResponse BootcampInfo (1: base.IDReq req) (api.post = "/service/wx/bootcamp-info")
   /**训练营报名*/
   base.NilResponse JoinBootcamp (1: JoinBootcampReq req) (api.post = "/service/wx/join-bootcamp")

   /**社群列表*/
   base.NilResponse CommunityList (1: CommunityListReq req) (api.post = "/service/wx/community-list")
   /**社群详情*/
   base.NilResponse CommunityInfo (1: base.IDReq req) (api.post = "/service/wx/community-info")
   /**社群报名*/
   base.NilResponse JoinCommunity (1: JoinCommunityReq req) (api.post = "/service/wx/join-community")

   /**场馆列表*/
   base.NilResponse VenueList (1: VenueListReq req) (api.post = "/service/wx/venue-list")
   /**场馆详情*/
   base.NilResponse VenueInfo (1: base.IDReq req) (api.post = "/service/wx/venue-info")

   /**扫码入场*/
   base.NilResponse ScanEnterQR (1: ScanQREnterReq req) (api.post = "/service/wx/scan-QR-Enter")
   /**激活*/
   base.NilResponse Activation (1: base.IDReq req) (api.post = "/service/wx/activation")


   /**会员约课*/
   base.NilResponse CreateMemberSchedule(1: CreateMemberScheduleReq req) (api.post = "/service/wx/create-member-schedule")
   /**产品详情*/
   base.NilResponse ProductInfo(1: base.IDReq req) (api.post = "/service/wx/product-info")
   /**产品列表*/
   base.NilResponse ProductList(1:  ProductListReq req) (api.post = "/service/wx/product-list")
   /**下单*/
   base.NilResponse CreateOrder(1: CreateOrderReq req) (api.post = "/service/wx/create-order")

   /**教练列表*/
   base.NilResponse CoachList(1: CoachListReq req) (api.post = "/service/wx/coach-list")
   /**教练详情*/
   base.NilResponse CoachInfo(1: base.IDReq req) (api.post = "/service/wx/coach-info")

   /**场地列表*/
   base.NilResponse PlaceList(1: PlaceListReq req) (api.post = "/service/wx/place-list")
   /**场地详情*/
   base.NilResponse PlaceInfo(1: base.IDReq req) (api.post = "/service/wx/place-info")
   /**场地预约*/
   base.NilResponse CreatePlaceSchedule(1: CreatePlaceScheduleReq req) (api.post = "/service/wx/create-place-schedule")

}

service WxStaffService {

   /**我的会员*/
   base.NilResponse MyMember(1: MyMemberReq req) (api.post = "/service/wx/staff-my-member")
   /**课程列表*/
   base.NilResponse ScheduleMemberList(1: schedule.ScheduleMemberListReq req) (api.post = "/service/wx/staff-schedule-member-list")
   /**会员签到*/
   base.NilResponse SignMemberSchedule(1: SignMemberScheduleReq req) (api.post = "/service/wx/sign-member-schedule")
   /**教练签到*/
   base.NilResponse SignStaffSchedule(1: SignStaffScheduleReq req) (api.post = "/service/wx/sign-staff-schedule")


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