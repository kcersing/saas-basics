namespace go order

include "../base/base.thrift"

service OrderService {

    base.NilResponse Update(1: UpdateOrderReq req) (api.post = "/service/order/update")
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/order/status")
    base.NilResponse ListOrder(1: ListOrderReq req )(api.post = "/service/order/list") // 订单列表
    base.NilResponse GetOrderById(1: base.IDReq req) (api.get = "/service/order/info") // 订单详情

}
struct ListOrderReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string mobile="" (api.raw = "mobile")
    4: optional list<i64> sellId=0 (api.raw = "sellId")
    5: optional list<i64> productId=0 (api.raw = "productId")
    6: optional list<i64> venueId=0 (api.raw = "venueId")
    7: optional list<i64> status=0 (api.raw = "status")
    8: optional string orderSn="" (api.raw = "orderSn")
    9: optional string startCompletionAt="" (api.raw = "startCompletionAt")
    10: optional string endCompletionAt="" (api.raw = "endCompletionAt")
    11: optional string productType="" (api.raw = "productType")
    12: optional string nature="" (api.raw = "nature")
    13: optional string name ="" (api.raw = "name ")
    14: optional string member_name="" (api.raw = "member_name")
}
struct OrderInfo {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string orderSn="" (api.raw = "orderSn")
    3:  optional i64 status=0 (api.raw = "status")
    4:  optional string source="" (api.raw = "source")
    5:  optional string device="" (api.raw = "device")
    6:  optional string nature="" (api.raw = "nature")
    10: optional string productType="" (api.raw = "productType")
    11:  optional i64 venueId=0 (api.raw = "venueId")
    12:  optional i64 memberId=0 (api.raw = "memberId")
    13:  optional i64 createId=0 (api.raw = "createId")
    15:  optional string completionAt="" (api.raw = "completionAt")
    16:  optional string createdAt="" (api.raw = "createdAt")
    17:  optional string updatedAt="" (api.raw = "updatedAt")

    18:  optional string memberName="" (api.raw = "memberName")
    19:  optional string memberMobile="" (api.raw = "memberMobile")

    250: optional OrderAmount orderAmount={} (api.raw = "orderAmount")
    251: optional OrderItem orderItem={} (api.raw = "orderItem")
    252: optional list<OrderPay> orderPay={} (api.raw = "orderPay")
    253: optional list<OrderSales> orderSales={} (api.raw = "orderSales")
}

struct OrderAmount {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional double total=0 (api.raw = "total")
    3:  optional double actual=0 (api.raw = "actual")
    4:  optional double residue=0 (api.raw = "residue")
    5:  optional double remission=0 (api.raw = "remission")
    6:  optional i64 orderId=0 (api.raw = "orderId")
    16:  optional string createdAt="" (api.raw = "createdAt")
    17:  optional string updatedAt="" (api.raw = "updatedAt")
}

struct OrderItem {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional i64 productId=0 (api.raw = "productId")
    3:  optional i64 relatedUserProductId=0 (api.raw = "relatedUserProductId")
    4:  optional i64 contestId=0 (api.raw = "contestId")
    5:  optional i64 bootcampId=0 (api.raw = "bootcampId")
    6:  optional i64 orderId=0 (api.raw = "orderId")
    7:  optional string data="" (api.raw = "data")
    8:  optional string name="" (api.raw = "name")
    16:  optional string createdAt="" (api.raw = "createdAt")
    17:  optional string updatedAt="" (api.raw = "updatedAt")
}

struct OrderPay{
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional double pay=0 (api.raw = "pay")
    3:  optional double remission=0 (api.raw = "remission")
    6:  optional i64 orderId=0 (api.raw = "orderId")
    7:  optional string payWay="" (api.raw = "payWay")
    8:  optional string paySn="" (api.raw = "paySn")
    9:  optional string prepayId="" (api.raw = "prepayId")
    10:  optional string payExtra="" (api.raw = "payExtra")
    16:  optional string createdAt="" (api.raw = "createdAt")
    17:  optional string note="" (api.raw = "note")
    18:  optional string updatedAt="" (api.raw = "updatedAt")
}

struct OrderSales{
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional i64 salesId=0 (api.raw = "salesId")
    3:  optional i64 ratio=0 (api.raw = "ratio")
    6:  optional i64 orderId=0 (api.raw = "orderId")
}


struct UpdateOrderReq {

}

