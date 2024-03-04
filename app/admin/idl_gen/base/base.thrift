namespace go base

struct BaseResp {
    1: string StatusMessage = ""
    2: i32 StatusCode = 0
    3: optional map<string, string> Extra
}

struct BaseResponse {
    1: i64 status_code,   // Status code, 0-success, other values-failure
    2: string status_msg, // Return status description
}

struct NilResponse {

}

struct Empty {

}

struct IDReq{
    1: string id,
}

struct PageInfoReq{
    1: string page,
    2: string pageSize,
}

struct StatusCodeReq {
    1: string id,
    2: string status,
}

enum Err {
    Success            = 0,
    NoRoute            = 1,
    NoMethod           = 2,
    BadRequest         = 10000,
    ParamsErr          = 10001,
    AuthorizeFail      = 10002,
    TooManyRequest     = 10003,
    ServiceErr         = 20000,
    RPCUserSrvErr      = 30000,
    UserSrvErr         = 30001,
    RPCBlobSrvErr      = 40000,
    BlobSrvErr         = 40001,
    RPCCarSrvErr       = 50000,
    CarSrvErr          = 50001,
    RPCProfileSrvErr   = 60000,
    ProfileSrvErr      = 60001,
    RPCTripSrvErr      = 70000,
    TripSrvErr         = 70001,
    RecordNotFound     = 80000,
    RecordAlreadyExist = 80001,
    DirtyData          = 80003,
}