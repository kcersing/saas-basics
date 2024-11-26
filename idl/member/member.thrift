namespace go member

include "../base/base.thrift"

// Create or update user information request | 创建或更新用户信息
struct CreateOrUpdateMemberReq {
    1:  optional i64 id (api.raw = "id")
    2:  optional string avatar (api.raw = "avatar")
    4:  optional string mobile (api.raw = "mobile")
    5:  optional string email (api.raw = "email")
    6:  optional i64 status (api.raw = "status")
    7:  optional string name (api.raw = "name")
    8:  optional string gender (api.raw = "gender")
    9:  optional string wecom (api.raw = "wecom")
    10: optional i64 createId (api.raw = "createId")
    11: optional string birthday (api.raw = "birthday")
    12: optional string	Password (api.raw = "password")
}
struct MemberListResp {
    1: base.BaseResp resp
    2: optional list<MemberInfo> extra
}

struct MemberInfo {
 1: i64	Id  (api.raw = "id")
 2:	string Name  (api.raw = "name")
 3: i64 Condition (api.raw = "condition")
 4: i64 Status    (api.raw = "status")
 5:	string Nickname  (api.raw = "nickname")
 6:	string Mobile    (api.raw = "mobile")
 7:	string Email (api.raw = "email")
 8:	string Avatar    (api.raw = "avatar")
 9:	string Gender    (api.raw = "gender")
 10: i64 Age   (api.raw = "age")
 11: string Birthday  (api.raw = "birthday")
	//微信号
 12: string Wecom (api.raw = "wecom")

}

struct MemberPrivacy{
	//身份证号
 1: string	IdentityCard (api.raw = "identityCard")
	//正面
 2: string	FaceIdentityCard (api.raw = "faceIdentityCard")
	//反面
 3: string BackIdentityCard (api.raw = "backIdentityCard")
	//人脸照片
 4: string FacePic (api.raw = "facePic")
	//特征值
 5: string FaceEigenvalue (api.raw = "faceEigenvalue")
	//人脸更新时间
 6: string FaceUpdateTime (api.raw = "faceUpdateTime")
}


// Get user list request | 获取用户列表请求参数
struct MemberListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional string name (api.raw = "name")
    4:  optional string mobile (api.raw = "mobile")
}

struct MemberSearchReq {
    1:  string value (api.raw = "value")
    2:  string option (api.raw = "option")
}


service MemberService {

  // 新增用户
  base.NilResponse CreateMember(1: CreateOrUpdateMemberReq req) (api.post = "/service/member/create")

  // 更新用户
  base.NilResponse UpdateMember(1: CreateOrUpdateMemberReq req) (api.post = "/service/member/update")

  // 获取用户基本信息
  MemberInfo MemberInfo(1: base.IDReq req) (api.post = "/service/member/info")

  MemberPrivacy memberPrivacy(1: base.IDReq req) (api.post = "/service/member/privacy")

  // 获取用户列表
  MemberListResp MemberList(1: MemberListReq req) (api.post = "/service/member/list")

  // 更新用户状态
  base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/member/status")

  // 搜索会员
  MemberListResp MemberSearch(1: MemberSearchReq req) (api.post = "/service/member/search")

}
