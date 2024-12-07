namespace go payment

include "../base/base.thrift"


service PaymentService {
    base.NilResponse WXPay(1: WXPayReq req) (api.post = "/service/payment/WXPay") // 微信小程序支付
    base.NilResponse WXNotify(1: WXNotifyReq req) (api.post = "/service/payment/WXNotify") // 回调
}

struct WXPayReq {
    1: i64 orderId (api.raw = "orderId")
    2: string  orderSn (api.raw = "orderSn")
    3: double  fee (api.raw = "orderSn")
}
struct WXNotifyReq {
    1: string id (api.raw = "id")
    2: string createTime (api.raw = "create_time")
    3: string resourceType (api.raw = "resource_type")
    4: string eventType (api.raw = "event_type")
    5: string summary (api.raw = "summary")
    6: Resource resource (api.raw = "resource")

}
struct Resource {
    1: string original_type (api.raw = "original_type")
    2: string algorithm (api.raw = "algorithm")
    3: string ciphertext (api.raw = "ciphertext")
    4: string associated_data (api.raw = "associated_data")
    5: string nonce (api.raw = "nonce")
}