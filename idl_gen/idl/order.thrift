namespace go order

include "../base/base.thrift"

service OrderService {

    base.NilResponse Update(1: UpdateOrderReq req) (api.post = "/service/order/update")
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/order/status")
    base.NilResponse ListOrder(1: ListOrderReq req )(api.post = "/service/order/list") // 订单列表
    base.NilResponse GetOrderById(1: base.IDReq req) (api.get = "/service/order/info") // 订单详情

}
struct OrderInfo {

}
struct UpdateOrderReq {

}

struct ListOrderReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string mobile="" (api.raw = "mobile")
    4: optional list<i64> sellId=0 (api.raw = "sellId")
    5: optional list<i64> productId=0 (api.raw = "productId")
    6: optional list<i64> venueId=0 (api.raw = "venueId")
}