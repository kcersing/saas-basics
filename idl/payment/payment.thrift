namespace go payment

include "../base/base.thrift"

service PaymentService {

    ChargeResp Charge(1: ChargeReq req)  (api.post = "/service/payment/charge")// 创建订单

    base.NilResponse UnifyPay(1: UnifyPayReq req) (api.post = "/service/order/unifyPay") // 支付

    base.NilResponse QRPay(1: QRPayReq req) (api.post = "/service/order/QRPay") // 订单详情
}

struct ChargeReq {
   1: i64 amount;
   2: string order_id;
   3: i64 user_id;
   4: i64 member_id;
}

struct ChargeResp {
  string transaction_id = 1;
}

struct UnifyPayReq {
    1: string orderSn,
    2: double payment,
    3: double remission,
    4: string note,
}

struct QRPayReq {
    1: string orderSn,
    2: string payType,
}