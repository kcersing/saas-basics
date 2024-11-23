namespace go company.entry

include "../base/base.thrift"

service CompanyService {
  base.NilResponse CreateEntry(1: CreateEntry req) (api.post = "/service/entry/list")
  EntryListResp EntryList(1: EntryListReq req) (api.post = "/service/entry/list")
}
struct CreateEntry  {
    1:i64 MemberProductId  (api.raw = "member_product_id")
	2:i64 MemberPropertyId (api.raw = "member_property_id")
	3:i64 EntryTime        (api.raw = "entry_time")
	4:i64 LeavingTime      (api.raw = "leaving_time")
	5:i64 MemberId         (api.raw = "member_id")
	6:i64 UserId           (api.raw = "user_id")
	7:i64 VenueId          (api.raw = "venue_id")
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

    3:  optional i64 memberId (api.raw = "member_id")
    4:  optional i64 venueId (api.raw = "venue_id")

    5:  optional i64 memberProductId (api.raw = "member_product_id")
    6:  optional i64 memberPropertyId (api.raw = "member_property_id")
    7:  optional i64 userId (api.raw = "user_id")
    8:  optional string entryTime (api.raw = "entry_time")
    9:  optional string leavingTime (api.raw = "leaving_time")

}