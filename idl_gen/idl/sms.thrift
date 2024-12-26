namespace go sms

include "../base/base.thrift"

struct SmsOrderListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")

}
struct SmsSendListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string mobile="" (api.raw = "mobile")
}

struct SmsSend{
    1: optional string createdAt = ""  (api.raw = "createdAt")
    2: optional i64 status=0  (api.raw = "status")
    3:  optional string mobile="" (api.raw = "mobile")
    4:  optional string name="" (api.raw = "name")
    5:  optional string code="" (api.raw = "code")
}

service SmsService {
    /**订单记录*/
    base.NilResponse SmsInfo(1: base.IDReq req) (api.post = "/service/sms/info")
    /**订单记录*/
    base.NilResponse SmsOrderList(1: SmsOrderListReq req) (api.post = "/service/sms/order-list")
    /**发送记录*/
    base.NilResponse SmsSendList(1: SmsSendListReq req) (api.post = "/service/sms/send-list")

}