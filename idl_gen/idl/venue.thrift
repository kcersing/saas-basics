namespace go venue

include "../base/base.thrift"
struct VenuePlaceListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string name = "" (api.raw = "name")
    11: optional i64 status = 0 (api.raw = "status")
}
struct VenueListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string name="" (api.raw = "name")
    4:  optional string type =""(api.raw = "type")
    5:  optional list<i64> classify =0 (api.raw = "classify")
    11: optional i64 status=0 (api.raw = "status")
}
struct VenueInfo {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string name="" (api.raw = "name")
    3:  optional string address="" (api.raw = "address")
    4:  optional string addressDetail="" (api.raw = "addressDetail")
    5:  optional string latitude="" (api.raw = "latitude")
    6:  optional string longitude =""(api.raw = "longitude")
    7:  optional string mobile="" (api.raw = "mobile")
    8:  optional string pic =""(api.raw = "pic")
    9:  optional string information="" (api.raw = "information")
    10:  optional i64 status=1 (api.raw = "status")
    11:  optional string createdAt="" (api.raw = "createdAt")
    12:  optional string updatedAt="" (api.raw = "updatedAt")
    13:  optional string email =""(api.raw = "email")
    14:  optional string type =""(api.raw = "type")
    15:  optional list<i64> classify =0 (api.raw = "classify")
    16:  optional string seal =""(api.raw = "seal")
    17:  optional list<base.List>  classifyName ="" (api.raw = "classifyName")
}

struct VenuePlaceInfo {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string name="" (api.raw = "name")
    3:  optional i64 venueId=0 (api.raw = "venueId")
    4:  optional string pic="" (api.raw = "pic")
    5:  optional i64 status=1 (api.raw = "status")
    6:  optional string createdAt="" (api.raw = "createdAt")
    7:  optional string updatedAt="" (api.raw = "updatedAt")
    /**可容纳人数*/
    8:  optional i64 number=0 (api.raw = "number")
    /**详情*/
    9:  optional string information="" (api.raw = "information")
    /**分类*/
    10:  optional i64 classify =0(api.raw = "classify")
    /**是否展示:1展示;2不展示*/
    11:  optional i64 isShow=1 (api.raw = "isShow")
    /**是否开放:1开放;2关闭*/
    12:  optional i64 isAccessible=1 (api.raw = "isAccessible")
}

service VenueService {

    /**添加场地*/
    base.NilResponse CreateVenuePlace(1: VenuePlaceInfo req) (api.post = "/service/place/create")
    /**编辑场地*/
    base.NilResponse UpdateVenuePlace(1: VenuePlaceInfo req) (api.post = "/service/place/update")
    /**删除场地*/
    base.NilResponse UpdateVenuePlaceStatus(1: base.StatusCodeReq req) (api.post = "/service/place/status")
    /**场地列表*/
    base.NilResponse VenuePlaceList(1: VenuePlaceListReq req) (api.post = "/service/place/list")
     /**场地詳情*/
    base.NilResponse VenuePlaceInfo(1: base.IDReq  req) (api.post = "/service/place/info")
    /**添加场馆*/
    base.NilResponse CreateVenue(1: VenueInfo req) (api.post = "/service/venue/create")
    /**编辑场馆*/
    base.NilResponse UpdateVenue(1: VenueInfo req) (api.post = "/service/venue/update")
    /**编辑场馆状态*/
    base.NilResponse UpdateVenueStatus(1: base.StatusCodeReq req) (api.post = "/service/venue/status")
    /**场馆列表*/
    base.NilResponse VenueList(1: VenueListReq req) (api.post = "/service/venue/list")
    /**场馆詳情*/
    base.NilResponse VenueInfo(1: base.IDReq  req) (api.post = "/service/venue/info")
}