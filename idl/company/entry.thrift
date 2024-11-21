namespace go admin.entry

include "../base/base.thrift"
include "../base/data.thrift"


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
service CompanyService {
  base.NilResponse EntryList(1: EntryListReq req) (api.post = "/service/entry/list")
}