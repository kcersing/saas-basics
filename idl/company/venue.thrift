namespace go company.venue

include "../base/base.thrift"

struct VenuePlaceListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional string name (api.raw = "name")
    11:  optional i64 status (api.raw = "status")
}
struct VenueListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional string name (api.raw = "name")
    11:  optional i64 status (api.raw = "status")
}
struct VenueInfo {
    1:  optional i64 id (api.raw = "id")
    2:  optional string name (api.raw = "name")
    3:  optional string address (api.raw = "address")
    4:  optional string addressDetail (api.raw = "addressDetail")
    5:  optional string latitude (api.raw = "latitude")
    6:  optional string longitude (api.raw = "longitude")
    7:  optional string mobile (api.raw = "mobile")
    8:  optional string pic (api.raw = "pic")
    9:  optional string information (api.raw = "information")
    10:  optional i64 status (api.raw = "status")
    11:  optional string CreatedAt (api.raw = "createdAt")
    12:  optional string UpdatedAt (api.raw = "updatedAt")
}

struct VenuePlaceInfo {
    1:  optional i64 id (api.raw = "id")
    2:  optional string name (api.raw = "name")
    3:  optional i64 venueId (api.raw = "venueId")
    4:  optional string pic (api.raw = "pic")
    5:  optional i64 status (api.raw = "status")
    6:  optional string CreatedAt (api.raw = "createdAt")
    7:  optional string UpdatedAt (api.raw = "updatedAt")
}
struct VenueListResp {
    1: base.BaseResp resp
    2: optional list<VenueInfo> extra
}
struct VenuePlaceListResp {
    1: base.BaseResp resp
    2: optional list<VenuePlaceInfo> extra
}
service VenueService {
    // 添加
    base.NilResponse CreateVenuePlace(1: VenuePlaceInfo req) (api.post = "/service/place/create")
    // 编辑
    base.NilResponse UpdateVenuePlace(1: VenuePlaceInfo req) (api.post = "/service/place/update")
    // 删除
    base.NilResponse UpdateVenuePlaceStatus(1: base.StatusCodeReq req) (api.post = "/service/place/status")
    // 列表
    VenuePlaceListResp VenuePlaceList(1: VenuePlaceListReq req) (api.post = "/service/place/list")

    // 添加
    base.NilResponse CreateVenue(1: VenueInfo req) (api.post = "/service/venue/create")
    // 编辑
    base.NilResponse UpdateVenue(1: VenueInfo req) (api.post = "/service/venue/update")
    base.NilResponse UpdateVenueStatus(1: base.StatusCodeReq req) (api.post = "/service/venue/status")
    // 列表
    VenueListResp VenueList(1: VenueListReq req) (api.post = "/service/venue/list")
}




