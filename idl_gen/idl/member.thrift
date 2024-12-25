namespace go member

include "../base/base.thrift"

// Create or update user information request | 创建或更新用户信息
struct CreateOrUpdateMemberReq {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string name="" (api.raw = "name")
    3: optional i64	mobileAscription=0 (api.raw = "mobileAscription")
    4:  optional string mobile="" (api.raw = "mobile")
    5:  optional string gender="未知" (api.raw = "gender")
    6:  optional string birthday="" (api.raw = "birthday")
    7:  optional i64 source=0 (api.raw = "source")
    8:  optional i64 grade=0 (api.raw = "grade")
    9: optional i64  intention=0 (api.raw = "intention")
}

struct MemberInfo {
     1:optional i64 id=0  (api.raw = "id")
     2:optional	string name ="" (api.raw = "name")
     3:optional i64 condition=0 (api.raw = "condition")
     4:optional i64 status=0 (api.raw = "status")
     5:optional	string username = "" (api.raw = "username")
     6:optional	string mobile ="" (api.raw = "mobile")
     7:optional	string avatar ="" (api.raw = "avatar")
     8:optional string createdAt ="" (api.raw = "createdAt")
     9:optional string updatedAt  =""(api.raw = "updatedAt")

      10:optional string conditionName=""(api.raw = "conditionName")

     250:optional MemberProfile profile={} (api.raw = "profile")
     251:optional MemberDetails details={} (api.raw = "details")
     252:optional MemberPrivacy privacy={} (api.raw = "privacy")
}

struct MemberProfile{
    1:  optional i64 mobileAscription=0 (api.raw = "mobileAscription")
    2:  optional string fatherName="" (api.raw = "fatherName")
    3:  optional string motherName="" (api.raw = "motherName")
    4:  optional i64 grade=0 (api.raw = "grade")
    5:  optional i64 intention=0 (api.raw = "intention")
    6:  optional i64 source=0 (api.raw = "source")
    7:  optional string gradeName="" (api.raw = "gradeName")
    8:  optional string intentionName="" (api.raw = "intentionName")
    9:  optional string sourceName="" (api.raw = "sourceName")

    10:	optional string email="" (api.raw = "email")
    11:	optional string gender = "" (api.raw = "gender")
    12: optional i64 age  =0 (api.raw = "age")
    13: optional string wecom ="" (api.raw = "wecom")
    14: optional string birthday=""   (api.raw = "birthday")

     250:optional i64 id=0  (api.raw = "id")
}

struct MemberDetails {
	 //消费总金额
     2:optional double moneySum=0 (api.raw = "moneySum")
     //首次的产品
     3:optional i64 productId =0  (api.raw = "productId")
     4:optional string productName =""(api.raw = "productName")
     //首次消费场馆
     5:optional i64 productVenue=0 (api.raw = "productVenue")
     6:optional string productVenueName =""(api.raw = "productVenueName")
     //进馆总次数
     7:optional i64 entrySum=0 (api.raw = "entrySum")
     //最后一次进馆时间
     8:optional string entryLastTime="" (api.raw = "entryLastTime")
     //进馆最后期限时间
     9:optional string entryDeadlineTime =""(api.raw = "entryDeadlineTime")
     //最后一次上课时间
     10:optional string classLastTime="" (api.raw = "classLastTime")
     //关联员工
     11:optional i64 relationUid =0  (api.raw = "relationUid")
     12:optional string relationUname="" (api.raw = "relationUname")
     //关联会员
     13:optional i64 relationMid=0   (api.raw = "relationMid")
     14:optional string relationMname="" (api.raw = "relationMname")

     15:optional i64 CreatedId=0   (api.raw = "createdId")
     16:optional string CreatedName="" (api.raw = "createdName")

}



struct MemberPrivacy {
	//身份证号
     1: string	identityCard (api.raw = "identityCard")
        //正面
     2: string	faceIdentityCard (api.raw = "faceIdentityCard")
        //反面
     3: string backIdentityCard (api.raw = "backIdentityCard")
        //人脸照片
     4: string facePic (api.raw = "facePic")
        //特征值
     5: string faceEigenvalue (api.raw = "faceEigenvalue")
        //人脸更新时间
     6: string faceUpdateTime (api.raw = "faceUpdateTime")
     7: string facePicUpdatedTime (api.raw = "facePicUpdatedTime")
}



// Get user list request | 获取用户列表请求参数
struct MemberListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string name = "" (api.raw = "name")
    4:  optional string mobile = "" (api.raw = "mobile")
    5:  optional i64 source = 0 (api.raw = "source")
    6:  optional i64 intention = 0 (api.raw = "intention")
    7:  optional i64 createdId = 0 (api.raw = "createdId")
    8:  optional string startCreatedAt = "" (api.raw = "startCreatedAt")
    9:  optional string endCreatedAt = "" (api.raw = "endCreatedAt")
}

struct MemberSearchReq {
    1:  string value="" (api.raw = "value")
    2:  string option="" (api.raw = "option")
}

struct MemberContractListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional i64 memberId=0 (api.raw = "memberId")
    4:  optional i64 venueId =0 (api.raw = "venueId")
    5:  optional i64 contractId =0 (api.raw = "contractId")
}

struct MemberContractInfo{
    1:  optional string	name (api.raw = "name")
    2:  optional i64 memberId (api.raw = "memberId")
    3:  optional string	memberName (api.raw = "memberName")
    4:  optional i64 venueId (api.raw = "venueId")
    5:  optional string	venueName (api.raw = "venueName")
    6:  optional i64 memberProductId (api.raw = "memberProductId")
    7:  optional string	memberProductName (api.raw = "memberProductName")
    8:  optional i64 contractId (api.raw = "contractId")
    10:  optional string sign (api.raw = "sign")
    11:  optional string signImg (api.raw = "signImg")
    12:  optional string content (api.raw = "content")
}


struct UpdateMemberFollowReq{
    1:  optional i64 memberId=0 (api.raw = "memberId")
    2:  optional i64 followId=0 (api.raw = "followId")
}



service MemberService {


  // 新增用户
  base.NilResponse CreateMember(1: CreateOrUpdateMemberReq req) (api.post = "/service/member/create")

  // 更新用户
  base.NilResponse UpdateMember(1: CreateOrUpdateMemberReq req) (api.post = "/service/member/update")

  // 获取用户基本信息
  base.NilResponse MemberInfo(1: base.IDReq req) (api.post = "/service/member/info")

  // 获取用户列表
  base.NilResponse MemberFullList(1: MemberListReq req) (api.post = "/service/member/full-list")

  // 获取用户列表
  base.NilResponse MemberPotentialList(1: MemberListReq req) (api.post = "/service/member/potential-list")

  // 更新用户状态
  base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/member/status")

  // 更新用户状态
  base.NilResponse UpdateMemberFollow(1: UpdateMemberFollowReq req) (api.post = "/service/member/update-follow")

  base.NilResponse MemberContractList(1: MemberContractListReq req) (api.post = "/service/member/contract-list")

}
