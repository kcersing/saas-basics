namespace go member

include "../base/base.thrift"

struct MemberNode {
 1: string FacePicUpdatedTime (api.raw = "facePicUpdatedTime")
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

service CompanyService {
  // 获取会员节点信息
  MemberNode MemberNode(1: base.IDReq req) (api.post = "/service/member/node")

  base.NilResponse MemberContractList(1: MemberContractListReq req) (api.post = "/service/member/contract-list")

}
