namespace go company.entry

include "../base/base.thrift"

service CompanyService {
  base.NilResponse CreateEntry(1: CreateEntry req) (api.post = "/service/entry/list")
  EntryListResp EntryList(1: EntryListReq req) (api.post = "/service/entry/list")
}
struct CreateEntry  {
    1:i64 MemberProductId  (api.raw = "memberProductId")
	2:i64 MemberPropertyId (api.raw = "memberPropertyId")
	3:i64 EntryTime        (api.raw = "entryTime")
	4:i64 LeavingTime      (api.raw = "leavingTime")
	5:i64 MemberId         (api.raw = "memberId")
	6:i64 UserId           (api.raw = "userId")
	7:i64 VenueId          (api.raw = "venueId")
}
struct EntryListResp {
    1: base.BaseResp resp
    2: optional list<EntryInfo> extra
}
struct EntryInfo  {
	1:i64 Id
	2:i64 MemberProductId
	3:i64 MemberPropertyId
	4:i64 EntryTime
	5:i64 LeavingTime
	6:i64 MemberId
	7:i64 UserId
	8:i64 VenueId
	9:string CreatedAt
	10:string UpdatedAt

	11:string VenueName
	12:string UserName
	13:string MemberName
	14:string MemberProductName
	15:string MemberPropertyName
}


struct EntryListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "memberId")
    4:  optional i64 venueId (api.raw = "venueId")
    5:  optional i64 memberProductId (api.raw = "memberProductId")
    6:  optional i64 memberPropertyId (api.raw = "memberPropertyId")
    7:  optional i64 userId (api.raw = "userId")
    8:  optional string entryTime (api.raw = "entryTime")
    9:  optional string leavingTime (api.raw = "leavingTime")
}