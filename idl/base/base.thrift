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

}

struct PageInfoReq{

}

struct StatusCodeReq {
    1: string id,
    2: string status,
}