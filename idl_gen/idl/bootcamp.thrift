
namespace go contest

include "../base/base.thrift"

service ContestService {

  base.NilResponse CreateContest(1: ContestInfo req) (api.post = "/service/contest/create")

  base.NilResponse UpdateContest(1: ContestInfo req) (api.post = "/service/contest/update")

  base.NilResponse ContestInfo(1: base.IDReq req) (api.post = "/service/contest/info")

  base.NilResponse ContestList(1: ContestListReq req) (api.post = "/service/contest/list")

  base.NilResponse UpdateContestStatus(1: base.StatusCodeReq req) (api.post = "/service/contest/status")

  base.NilResponse UpdateContestShow (1: base.StatusCodeReq req) (api.post = "/service/contest/show")

  base.NilResponse CreateParticipant(1: ParticipantInfo req) (api.post = "/service/participant/create")

  base.NilResponse UpdateParticipant(1: ParticipantInfo req) (api.post = "/service/participant/update")

  base.NilResponse ContestParticipantInfo(1: base.IDReq req) (api.post = "/service/participant/info")

  base.NilResponse ParticipantListList(1: ParticipantListReq req) (api.post = "/service/participant/list")

  base.NilResponse UpdateParticipantStatus(1: base.StatusCodeReq req) (api.post = "/service/participant/status")

  base.NilResponse ParticipantListListExport(1: ParticipantListReq req) (api.post = "/service/participant/export")

  base.NilResponse ResultsUpload(1: ResultsUploadReq req) (api.post = "/service/contest/results-upload")

  base.NilResponse PromotionalLinks(1: base.IDReq req) (api.post = "/service/contest/promotional-links")

  base.NilResponse DelContest(1: base.IDReq req) (api.post = "/service/contest/del")

  base.NilResponse ParticipantFinish(1: ParticipantFinishReq req) (api.post = "/service/participant/finish")
}

struct ParticipantFinishReq{
  1:  optional i64 contestId=0 (api.raw = "contestId")
  2:  optional list<i64> memberId (api.raw = "memberId")
}

struct ResultsUploadReq{
    1:  optional string pic="" (api.raw = "pic")
}
struct ContestListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string name = "" (api.raw = "name")
    4:  optional string signStartAt = "" (api.raw = "signStartAt")
    5:  optional string signEndAt = "" (api.raw = "signEndAt")
    6:  optional string startAt  = ""(api.raw = "startAt")
    7:  optional string endAt  = ""(api.raw = "endAt")
    8:  optional i64 condition=0 (api.raw = "condition")

}
struct ParticipantListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string name  = ""(api.raw = "name")
    4:  optional string mobile  = ""(api.raw = "mobile")
    5:  optional string sn = "" (api.raw = "sn")
    6:  optional i64 contestId=0 (api.raw = "contestId")
}
struct ParticipantInfo{
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional i64 contestId=0 (api.raw = "contestId")
    3:  optional string fields="" (api.raw = "fields")
    4:  optional string name="" (api.raw = "name")
    5:  optional string mobile="" (api.raw = "mobile")
    6:  optional string createdAt="" (api.raw = "createdAt")

    8:  optional string orderSn="" (api.raw = "orderSn")
    9:  optional double fee=0 (api.raw = "fee")
    10:  optional i64 status=1 (api.raw = "status")

}

struct ContestInfo{
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string name="" (api.raw = "name")
    3:  optional i64 signNumber=0 (api.raw = "signNumber")
    4:  optional string signStartAt="" (api.raw = "signStartAt")
    5:  optional string signEndAt="" (api.raw = "signEndAt")
    6:  optional i64 number=0 (api.raw = "number")
    7:  optional string startAt="" (api.raw = "startAt")
    8:  optional string endAt="" (api.raw = "endAt")
    9:  optional string pic="" (api.raw = "pic")

    10: optional string sponsor="" (api.raw = "sponsor")
    11:  optional double fee=0 (api.raw = "fee")
    12:  optional i64 isCancel=0 (api.raw = "isCancel")
    13:  optional i64 cancelTime=0 (api.raw = "cancelTime")
    14:  optional string detail="" (api.raw = "detail")
    15:  optional string signFields="" (api.raw = "signFields")
    16:  optional string createdAt="" (api.raw = "createdAt")
    17:  optional string updatedAt="" (api.raw = "updatedAt")

    18:  optional i64 condition=0 (api.raw = "condition")
    19:  optional i64 isFee=0 (api.raw = "isFee")

    20:  optional i64 createdId=0 (api.raw = "createdId")
    21:  optional string createdName="0" (api.raw = "createdName")
    22:  optional i64 isShow=1 (api.raw = "isShow")

}
