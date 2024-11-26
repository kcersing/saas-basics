namespace go base
//bool: 布尔值，对应Java中的boolean，
//byte: 有符号字节，对应Java中的byte，对应MySQL的tinyint
//i16: 16位有符号整型，对应Java中的short，对应MySQL的smallint
//i32: 32位有符号整型，对应Java中的int，对应MySQL的int
//i64: 64位有符号整型，对应Java中的long，对应MySQL的bigint
//double: 64位浮点型，对应Java中的double
//string: 字符串，对应Java中的String
//binary: Blob 类型，对应Java中的byte[]
//list<t1>:一系列的t1类型组成的有序表，元素可以重复
//set<t1>:一系列的t1类型组成的无序集合，元素不可以重复
//map<t1,t2>:以t1为key,t2为value的键值对，t1不可以重复

struct BaseResp {
    1: string message = ""
    2: i32 code = 0
    3: i64 total = 0
}

struct BaseResponse {
    1: i64 code,   // Status code, 0-success, other values-failure
    2: string message, // Return status description
}

struct NilResponse {

}

struct Empty {

}

struct IDReq{
    1: i64 id,
}
struct Ids{
    1: list<i64> ids,
}
struct PageInfoReq{
    1: i64 page,
    2: i64 pageSize,
}

struct StatusCodeReq {
    1: i64 id,
    2: i64 status,
}

struct Tree {
	1:string Title
    2:string Value
	3:string Key
	4:string Method
	5:list<Tree> Children
}

struct ListReq {
    1: optional string name (api.raw = "name")
    2: optional i64 dictionaryId (api.raw = "dictionaryId" )
    3: optional string type (api.raw = "type" )
    4: optional string mobile (api.raw = "mobile" )
    5: optional i64 product (api.raw = "product" )
    6: optional i64 venue (api.raw = "venue" )
}
struct SysList  {
    i64	Id
    string	Name
    string	Key
}
struct SysListResp {
    1: BaseResp resp
    2: optional list<SysList> extra
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