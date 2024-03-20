
namespace go admin.venue

include "../base/base.thrift"
include "../base/data.thrift"

struct Venue {
    1:  optional i64 id (api.raw = "id")
//    2:  optional i64 limit (api.raw = "limit")
    3:  optional string name (api.raw = "name")
    4:  optional string address (api.raw = "address")
    5:  optional string address_detail (api.raw = "address_detail")
    6:  optional string latitude (api.raw = "latitude")
    7:  optional string longitude (api.raw = "longitude")
    8:  optional string mobile (api.raw = "mobile")
    9:  optional string pic (api.raw = "pic")
    10:  optional string information (api.raw = "information")
    11:  optional i64 status (api.raw = "status")
}


struct VenueListReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 limit (api.raw = "limit")
    3:  optional string name (api.raw = "name")
//    4:  optional string address (api.raw = "address")
//    5:  optional string address_detail (api.raw = "address_detail")
//    6:  optional string latitude (api.raw = "latitude")
//    7:  optional string longitude (api.raw = "longitude")
//    8:  optional string mobile (api.raw = "mobile")
//    9:  optional string pic (api.raw = "pic")
//    10:  optional string information (api.raw = "information")
    11:  optional i64 status (api.raw = "status")
}





service VenueService {
  // 获取用户列表
  base.NilResponse VenueList(1: VenueListReq req) (api.post = "/api/admin/venue/list")
}