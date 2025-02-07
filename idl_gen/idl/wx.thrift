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

}