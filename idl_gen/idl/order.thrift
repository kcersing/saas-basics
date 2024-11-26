namespace go order

include "../base/base.thrift"
include "../base/data.thrift"

struct OrderItem {
    1: i64 order_id
    2: i64 user_id
    3: string user_name
    4: string address
    5: i64 product_id
    6: i64 stock_num
    7: string product_snapshot
}

struct CreateOrderReq {
    1: optional string assign_at
    2: optional list<propertyItem> cardProperty
    3: optional list<propertyItem> courseProperty
    4: optional list<propertyItem> classProperty
    5: optional i64 member
    6: optional i64 natureType
    7: optional i64 product
    8: optional double total
    9: optional i64 venue
    10: optional list<staffItem> staffs
    11:optional string signImg
    12:optional list<i64> contract
}

struct propertyItem{
    1:optional i64 property
    2:optional i64 quantity
}
struct staffItem{
    1:optional i64 id
    2:optional i64 ratio
}

struct ListOrderReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 member (api.raw = "member")
    4:  optional list<i64> sell (api.raw = "sell")
    5:  optional list<i64> product (api.raw = "product")
    6:  optional list<i64> venue (api.raw = "venue")
}

struct UpdateOrderReq {
    1: OrderItem order
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

service OrderService {
    base.NilResponse CreateOrder(1: CreateOrderReq req)  (api.post = "/api/idl/order/create")// 创建订单

    base.NilResponse UpdateOrder(1: UpdateOrderReq req) (api.post = "/api/idl/order/update")

    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/idl/order/status")

    base.NilResponse ListOrder(1: ListOrderReq req )(api.post = "/api/idl/order/list") // 订单列表

    base.NilResponse GetOrderById(1: base.IDReq req) (api.get = "/api/idl/order/info") // 订单详情

    base.NilResponse UnifyPay(1: UnifyPayReq req) (api.post = "/api/idl/order/unifyPay") // 支付

    base.NilResponse QRPay(1: QRPayReq req) (api.post = "/api/idl/order/QRPay") // 订单详情

}
