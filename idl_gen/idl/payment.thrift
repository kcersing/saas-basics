namespace go payment

include "../base/base.thrift"


service PaymentService {
    base.NilResponse WXPay(1: PayReq req) (api.post = "/service/payment/WXPay") // 微信小程序支付
    base.NilResponse WXQRPay(1: PayReq req) (api.post = "/service/payment/WXQRPay") // 微信小程序支付
    base.NilResponse WXNotify(1: NotifyReq req) (api.post = "/service/payment/WXNotify") // 回调
    base.NilResponse WXRefundOrder(1: RefundOrderReq req) (api.post = "/service/payment/WXRefundOrder") // 退款

}

struct PayReq {
    1:optional i64 orderId (api.raw = "orderId")
    2:optional string  orderSn="" (api.raw = "orderSn")
    3:optional double  total=0 (api.raw = "total")
    4:optional string  openId="" (api.raw = "openId")
}

struct NotifyReq {
    1: string id (api.raw = "id")
    2: string createTime (api.raw = "create_time")
    3: string resourceType (api.raw = "resource_type")
    4: string eventType (api.raw = "event_type")
    5: string summary (api.raw = "summary")
    6: Resource resource (api.raw = "resource")

}
struct RefundOrderReq {
   1: string transactionId (api.raw = "transactionId")
    2: string outRefundNo (api.raw = "outRefundNo")
}

struct Resource {
    1: string original_type (api.raw = "original_type")
    2: string algorithm (api.raw = "algorithm")
    3: string ciphertext (api.raw = "ciphertext")
    4: string associated_data (api.raw = "associated_data")
    5: string nonce (api.raw = "nonce")
}

