
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
}
struct ParticipantListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional string name (api.raw = "name")
    4:  optional string mobile (api.raw = "mobile")
    5:  optional string sn (api.raw = "sn")
}
struct ParticipantInfo{
    1:  optional i64 contestId (api.raw = "contestId")
    2:  optional i64 fields (api.raw = "fields")
    3:  optional string name (api.raw = "name")
    4:  optional string mobile (api.raw = "mobile")
}
struct ContestInfo{
    1:  optional i64 Id (api.raw = "id")
    2:  optional string Name (api.raw = "name")
    3:  optional i64 SignNumber (api.raw = "signNumber")
    4:  optional string SignStartAt (api.raw = "signStartAt")
    5:  optional string SignEndAt (api.raw = "signEndAt")
    6:  optional i64 Number (api.raw = "number")
    7:  optional string StartAt (api.raw = "startAt")
    8:  optional string EndAt (api.raw = "endAt")
    9:  optional string Pic (api.raw = "pic")

    10: optional string Sponsor (api.raw = "sponsor")
    11:  optional double Fee (api.raw = "fee")
    12:  optional i64 IsCancel (api.raw = "isCancel")
    13:  optional i64 CancelTime (api.raw = "cancelTime")
    14:  optional string Detail (api.raw = "detail")
    15:  optional string SignFields (api.raw = "signFields")
    16:  optional string CreatedAt (api.raw = "createdAt")
    17:  optional string UpdatedAt (api.raw = "updatedAt")

}
