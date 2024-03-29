namespace go admin.order

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
    1: required i64 user_id
    2: required list<i64> sell_id
    3: required i64 product_id
    4: required i64 stock_num
    5: required i64 venue_id
    6: required i64 create_id

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

service OrderService {
    base.NilResponse CreateOrder(1: CreateOrderReq req)  (api.post = "/api/admin/order/create")// 创建订单

    base.NilResponse UpdateOrder(1: UpdateOrderReq req) (api.post = "/api/admin/order/update")

    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/order/status")

    base.NilResponse ListOrder(1: ListOrderReq req )(api.post = "/api/admin/order/list") // 订单列表

    base.NilResponse GetOrderById(1: base.IDReq req) (api.get = "/api/admin/order/info") // 订单详情
}
