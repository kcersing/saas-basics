namespace go admin.cart

include "../base/base.thrift"

struct CartItme {
    1: optional i64 productId
    2: optional i64 quantity
}
struct AddItmeReq{
    1: optional CartItme item
    2: optional i64 member (api.raw = "member")
}


service CartService {
    base.NilResponse AddItem(1: AddItmeReq req)  (api.get = "/api/cart/add-item")
    base.NilResponse GetCart(1: base.IDReq req) (api.post = "/api/cart/get-cart")
    base.NilResponse EmptyCart(1: base.IDReq req) (api.post = "/api/cart/empty-cart")
}