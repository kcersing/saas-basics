
namespace go contest

include "../base/base.thrift"

service ContestService {

  base.NilResponse CreateContest(1: ContestInfo req) (api.post = "/service/contest/create")

  base.NilResponse UpdateContest(1: ContestInfo req) (api.post = "/service/contest/update")

  base.NilResponse ContestInfo(1: base.IDReq req) (api.post = "/service/contest/info")

  base.NilResponse ContestList(1: ContestListReq req) (api.post = "/service/contest/list")

  base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/contest/status")


  base.NilResponse CreateParticipant(1: ParticipantInfo req) (api.post = "/service/participant/create")

  base.NilResponse UpdateParticipant(1: ParticipantInfo req) (api.post = "/service/participant/update")

  base.NilResponse ContestParticipantInfo(1: base.IDReq req) (api.post = "/service/participant/info")

  base.NilResponse ParticipantListList(1: ParticipantListReq req) (api.post = "/service/participant/list")

  base.NilResponse UpdateParticipantStatus(1: base.StatusCodeReq req) (api.post = "/service/participant/status")

}
struct ContestListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional string name (api.raw = "name")
    4:  optional string signStartAt (api.raw = "signStartAt")
    5:  optional string signEndAt (api.raw = "signEndAt")
    6:  optional string startAt (api.raw = "startAt")
    7:  optional string endAt (api.raw = "endAt")
    8:  optional i64 condition (api.raw = "condition")

}
struct ParticipantListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional string name (api.raw = "name")
    4:  optional string mobile (api.raw = "mobile")
    5:  optional string sn (api.raw = "sn")
    6:  optional i64 contestId (api.raw = "contestId")
}
struct ParticipantInfo{
    1:  optional i64 id (api.raw = "id")
    2:  optional i64 contestId (api.raw = "contestId")
    3:  optional string fields (api.raw = "fields")
    4:  optional string name (api.raw = "name")
    5:  optional string mobile (api.raw = "mobile")
}
struct ContestInfo{
    1:  optional i64 id (api.raw = "id")
    2:  optional string name (api.raw = "name")
    3:  optional i64 signNumber (api.raw = "signNumber")
    4:  optional string signStartAt (api.raw = "signStartAt")
    5:  optional string signEndAt (api.raw = "signEndAt")
    6:  optional i64 number (api.raw = "number")
    7:  optional string startAt (api.raw = "startAt")
    8:  optional string endAt (api.raw = "endAt")
    9:  optional string pic (api.raw = "pic")

    10: optional string sponsor (api.raw = "sponsor")
    11:  optional double fee (api.raw = "fee")
    12:  optional i64 isCancel (api.raw = "isCancel")
    13:  optional i64 cancelTime (api.raw = "cancelTime")
    14:  optional string detail (api.raw = "detail")
    15:  optional string signFields (api.raw = "signFields")
    16:  optional string createdAt (api.raw = "createdAt")
    17:  optional string updatedAt (api.raw = "updatedAt")
    
    18:  optional string condition (api.raw = "condition")

}
