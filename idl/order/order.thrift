namespace go order

include "../base/base.thrift"

struct OrderItem {
    1: i64 orderId (api.raw = "orderId")
    2: i64 userId (api.raw = "userId")
    3: string userName (api.raw = "userName")
    4: string address (api.raw = "address")
    5: i64 productId (api.raw = "productId")
    6: i64 stockNum (api.raw = "stockNum")
    7: string productSnapshot (api.raw = "productSnapshot")
}

struct CreateOrderReq {
    1: optional string assignAt (api.raw = "assignAt")
    2: optional list<propertyItem> cardProperty (api.raw = "cardProperty")
    3: optional list<propertyItem> courseProperty (api.raw = "courseProperty")
    4: optional list<propertyItem> classProperty (api.raw = "classProperty")
    5: optional i64 memberId (api.raw = "memberId")
    6: optional i64 natureType (api.raw = "natureType")
    7: optional i64 productId (api.raw = "productId")
    8: optional double total (api.raw = "total")
    9: optional i64 venueId (api.raw = "venueId")
    10: optional list<staffItem> staffs (api.raw = "staffs")
    11:optional string signImg (api.raw = "signImg")
    12:optional list<i64> contract (api.raw = "contract")
}

struct propertyItem{
    1:optional i64 property (api.raw = "property")
    2:optional i64 quantity (api.raw = "quantity")
}
struct staffItem{
    1:optional i64 id (api.raw = "id")
    2:optional i64 ratio (api.raw = "ratio")
}

struct ListOrderReq {
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")
    3:  optional i64 memberId (api.raw = "memberId")
    4:  optional list<i64> sell (api.raw = "sell")
    5:  optional list<i64> product (api.raw = "product")
    6:  optional list<i64> venue (api.raw = "venue")
}

struct UpdateOrderReq {
    1: OrderItem order
}


service OrderService {

    base.NilResponse CreateOrder(1: CreateOrderReq req)  (api.post = "/service/order/create")// 创建订单

    base.NilResponse UpdateOrder(1: UpdateOrderReq req) (api.post = "/service/order/update")

    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/order/status")

    base.NilResponse ListOrder(1: ListOrderReq req )(api.post = "/service/order/list") // 订单列表

    base.NilResponse GetOrderById(1: base.IDReq req) (api.get = "/service/order/info") // 订单详情



}