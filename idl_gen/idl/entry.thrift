namespace go entry

include "../base/base.thrift"

service EntryService {
  base.NilResponse CreateEntry(1: CreateEntry req) (api.post = "/service/entry/create")
  base.NilResponse EntryList(1: EntryListReq req) (api.post = "/service/entry/list")
}
struct CreateEntry  {
    1:i64 MemberProductId  (api.raw = "memberProductId")
	3:string EntryTime        (api.raw = "entryTime")
	4:string LeavingTime      (api.raw = "leavingTime")
	5:i64 MemberId         (api.raw = "memberId")
	6:i64 UserId           (api.raw = "userId")
	7:i64 VenueId          (api.raw = "venueId")
}

struct EntryInfo  {
	1:i64 Id
	2:i64 MemberProductId
	4:string EntryTime
	5:string LeavingTime
	6:i64 MemberId
	7:i64 UserId
	8:i64 VenueId
	9:string CreatedAt
	10:string UpdatedAt

	11:string VenueName
	12:string UserName
	13:string MemberName
	14:string MemberProductName
}


struct EntryListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional i64 memberId = 0 (api.raw = "memberId")
    4:  optional i64 venueId= 0 (api.raw = "venueId")
    5:  optional i64 memberProductId= 0 (api.raw = "memberProductId")
    7:  optional i64 userId= 0 (api.raw = "userId")
    8:  optional string entryTime="" (api.raw = "entryTime")
    9:  optional string leavingTime="" (api.raw = "leavingTime")
}