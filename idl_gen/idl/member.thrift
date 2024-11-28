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
 7: string FacePicUpdatedTime (api.raw = "facePicUpdatedTime")
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
struct MemberNode {
	//消费总金额
 2: double MoneySum (api.raw = "moneySum")
	//首次的产品
 3: i64 ProductId   (api.raw = "productId")
 4: i64 ProductName (api.raw = "productName")
	//首次消费场馆
 5: i64 ProductVenue (api.raw = "productVenue")
 6: i64 ProductVenueName (api.raw = "productVenueName")
	//进馆总次数
 7: i64 EntrySum (api.raw = "entrySum")
	//最后一次进馆时间
 8: string EntryLastTime (api.raw = "entryLastTime")
	//进馆最后期限时间
 9: string EntryDeadlineTime (api.raw = "entryDeadlineTime")
	//最后一次上课时间
 10: string ClassLastTime (api.raw = "classLastTime")
	//关联员工
 11: i64 RelationUid   (api.raw = "relationUid")
 12: i64 RelationUname (api.raw = "relationUname")
	//关联会员
 13: i64 RelationMid   (api.raw = "relationMid")
 14: i64 RelationMname (api.raw = "relationMname")

 15: string CreatedAt (api.raw = "createdAt")
 16: string UpdatedAt (api.raw = "updatedAt")
}

struct MemberContractListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "memberId")
    4:  optional i64 venueId (api.raw = "venueId")
    5:  optional i64 contractId (api.raw = "contractId")
}
struct MemberProductListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "memberId")
    4:  optional string name (api.raw = "name")
    5:  optional i64 venueId (api.raw = "venueId")
    6:  optional i64 status (api.raw = "status")

}
struct MemberPropertyListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "memberId")
    4:  optional string type (api.raw = "type")
    5:  optional string name (api.raw = "name")
    6:  optional i64 venueId (api.raw = "venueId")
    7:  optional i64 status (api.raw = "status")
    8:  optional i64 memberProductId (api.raw = "memberProductId")
}

struct MemberProductSearchReq{
    1:  optional list<i64> members (api.raw = "members")
}
struct MemberPropertySearchReq{
    1:  optional list<i64> memberProductIds (api.raw = "memberProducts")
}

service MemberService {


  // 新增用户
  base.NilResponse CreateMember(1: CreateOrUpdateMemberReq req) (api.post = "/service/member/create")

  // 更新用户
  base.NilResponse UpdateMember(1: CreateOrUpdateMemberReq req) (api.post = "/service/member/update")

  // 获取用户基本信息
  base.NilResponse MemberInfo(1: base.IDReq req) (api.post = "/service/member/info")

  base.NilResponse memberPrivacy(1: base.IDReq req) (api.post = "/service/member/privacy")

  // 获取用户列表
  base.NilResponse MemberList(1: MemberListReq req) (api.post = "/service/member/list")

  // 更新用户状态
  base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/member/status")

  // 搜索会员
  base.NilResponse MemberSearch(1: MemberSearchReq req) (api.post = "/service/member/search")

  // 获取会员节点信息
  base.NilResponse MemberNode(1: base.IDReq req) (api.post = "/service/member/node")

  base.NilResponse MemberContractList(1: MemberContractListReq req) (api.post = "/service/member/contract-list")

  base.NilResponse MemberProductList(1: MemberProductListReq req) (api.post = "/service/member/product-list")

  base.NilResponse MemberPropertyList(1: MemberPropertyListReq req) (api.post = "/service/member/property-list")

  base.NilResponse MemberProductDetail(1: base.IDReq req) (api.post = "/service/member/product-detail")

  base.NilResponse MemberPropertyDetail(1: base.IDReq req) (api.post = "/service/member/property-detail")

  base.NilResponse MemberPropertyUpdate(1: MemberPropertyListReq req) (api.post = "/service/member/property-update")

  base.NilResponse MemberProductSearch(1: MemberProductSearchReq req) (api.post = "/service/member/search-product")

  base.NilResponse MemberPropertySearch(1: MemberPropertySearchReq req) (api.post = "/service/member/search-property")




}
