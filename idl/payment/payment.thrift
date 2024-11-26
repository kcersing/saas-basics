namespace go payment

include "../base/base.thrift"

service PaymentService {

    ChargeResp Charge(1: ChargeReq req)  (api.post = "/service/payment/charge")// 创建订单

    base.NilResponse UnifyPay(1: UnifyPayReq req) (api.post = "/service/order/unifyPay") // 支付

    base.NilResponse QRPay(1: QRPayReq req) (api.post = "/service/order/QRPay") // 订单详情
}

struct ChargeReq {
   1: i64 amount  (api.raw = "amount")
   2: string orderId (api.raw = "orderId")
   3: i64 userId (api.raw = "userId")
   4: i64 memberId (api.raw = "memberId")
}

struct ChargeResp {
  1: string transactionId  (api.raw = "transactionId")
}

struct UnifyPayReq {
    1: string orderSn (api.raw = "orderSn")
    2: double payment (api.raw = "payment")
    3: double remission (api.raw = "remission")
    4: string note (api.raw = "note")
}

struct QRPayReq {
    1: string orderSn (api.raw = "orderSn")
    2: string payType (api.raw = "payType")
}