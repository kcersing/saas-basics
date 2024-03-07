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
    2: required string address
    3: required i64 product_id
    4: required i64 stock_num
    5: required i64 venue_id
    6: i64 status
}

struct CreateOrderResp {
    255: base.BaseResp BaseResp
}

struct CancelOrderReq {
    1: required i64 order_id
}

struct CancelOrderResp {
    255: base.BaseResp BaseResp
}

struct ListOrderReq {
    1: required i64 user_id
    2: optional i64 status
}

struct ListOrderResp {
    1: list<OrderItem> orders
    255: base.BaseResp BaseResp
}

struct GetOrderByIdReq {
    1: required i64 order_id
}

struct GetOrderByIdResp {
    1: OrderItem order
    255: base.BaseResp BaseResp
}

service OrderService {
    CreateOrderResp CreateOrder(1: CreateOrderReq req) // 创建订单
    CancelOrderResp CancelOrder(1: CancelOrderReq req) // 取消订单
    ListOrderResp ListOrder(1: ListOrderReq req) // 订单列表
    GetOrderByIdResp GetOrderById(1: GetOrderByIdReq req) // 订单详情
}
