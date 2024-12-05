
namespace go contest

include "../base/base.thrift"

service BootcampService {

  base.NilResponse CreateBootcamp(1: BootcampInfo req) (api.post = "/service/bootcamp/create")

  base.NilResponse UpdateBootcamp(1: BootcampInfo req) (api.post = "/service/bootcamp/update")

  base.NilResponse BootcampInfo(1: base.IDReq req) (api.post = "/service/bootcamp/info")

  base.NilResponse BootcampList(1: BootcampListReq req) (api.post = "/service/bootcamp/list")

  base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/bootcamp/status")


  base.NilResponse CreateParticipant(1: ParticipantInfo req) (api.post = "/service/participant/create")

  base.NilResponse UpdateParticipant(1: ParticipantInfo req) (api.post = "/service/participant/update")

  base.NilResponse BootcampParticipantInfo(1: base.IDReq req) (api.post = "/service/participant/info")

  base.NilResponse ParticipantListList(1: ParticipantListReq req) (api.post = "/service/participant/list")

  base.NilResponse UpdateParticipantStatus(1: base.StatusCodeReq req) (api.post = "/service/participant/status")


  base.NilResponse ResultsUpload(1: ResultsUploadReq req) (api.post = "/service/bootcamp/results-upload")

  base.NilResponse PromotionalLinks(1: base.IDReq req) (api.post = "/service/bootcamp/promotional-links")

}
struct ResultsUploadReq{
    1:  optional string pic="" (api.raw = "pic")
}
struct BootcampListReq{
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
    6:  optional i64 bootcampId=0 (api.raw = "bootcampId")
}
struct ParticipantInfo{
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional i64 bootcampId=0 (api.raw = "bootcampId")
    3:  optional string fields="" (api.raw = "fields")
    4:  optional string name="" (api.raw = "name")
    5:  optional string mobile="" (api.raw = "mobile")
}
struct BootcampInfo{
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
    
    18:  optional string condition="" (api.raw = "condition")
    19:  optional i64 isFee=0 (api.raw = "isFee")

}
