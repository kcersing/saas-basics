namespace go admin.product

include "../base/base.thrift"

struct Property {
    1: i64 productId (api.raw = "productId")
    2: string name (api.raw = "name")// 名称
    3: double price (api.raw = "price")// 定价
    4: i64 duration (api.raw = "duration")// 时长
    5: i64 length (api.raw = "length")// 单次时长
    6: i64 count (api.raw = "count")// 次数
    7: string type (api.raw = "type")// 次数
    8: string data (api.raw = "data")
    9: i64 id (api.raw = "id")
}

struct Product {
    1: string name (api.raw = "name")// 商品名
    2: string pic (api.raw = "pic")// 主图
    3: string description (api.raw = "description")// 详情
    4: list <Property> property (api.raw = "property")// 属性
    5: double price (api.raw = "price")// 价格
    6: i64 stock (api.raw = "stock")// 库存
    7: i64 status (api.raw = "status")// 商品状态
    8: i64 id (api.raw = "id")
}

struct CreateOrUpdatePropertyReq {
    1: optional i64 id (api.raw = "id")
    2: optional string name (api.raw = "name")// 名称
    3: optional double price (api.raw = "price")// 定价
    4: optional i64 duration (api.raw = "duration")// 时长
    5: optional i64 length (api.raw = "length")// 单次时长
    6: optional i64 count (api.raw = "count")// 次数
    7: optional string type (api.raw = "type") // 类型
    8: optional list<i64> venueId (api.raw = "venueId")

}

struct CreateOrUpdateReq {
    1: optional i64 id (api.raw = "id")
    2: optional string name (api.raw = "name") // 商品名
    3: optional string pic (api.raw = "pic") // 主图
    4: optional string description  (api.raw = "description")// 详情
    5: optional double price (api.raw = "price") // 价格
    6: optional i64 stock (api.raw = "stock") // 库存
    7: optional i64 cardProperty (api.raw = "cardProperty")  // 属性
    8: optional list<i64> classProperty  (api.raw = "classProperty") // 属性
    9: optional list<i64> courseProperty  (api.raw = "courseProperty") // 属性
    10: optional i64 CreateId (api.raw = "createId") // 库存
}

struct ListReq {
    1: i64 page (api.raw = "page")
    2: i64 pageSize (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional list<i64> status (api.raw = "status")
    5: optional list<i64> venue (api.raw = "venue")
    6: optional list<string> createdTime (api.raw = "createdTime")
    7: optional string type (api.raw = "type") // 类型
}
struct PropertyListReq{
    1: i64 page (api.raw = "page")
    2: i64 pageSize (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional list<i64> status (api.raw = "status")
    5: optional list<i64> venue (api.raw = "venue")
    6: optional list<string> createdTime (api.raw = "createdTime")
    7: optional string type (api.raw = "type") // 类型
}

service ProductService {
    // 添加属性
    base.NilResponse CreateProperty(1: CreateOrUpdatePropertyReq req) (api.post = "/service/property/create")
    // 编辑属性
    base.NilResponse UpdateProperty(1: CreateOrUpdatePropertyReq req) (api.post = "/service/property/update")
    // 删除属性
    base.NilResponse DeleteProperty(1: base.IDReq req) (api.post = "/service/property/del")
    // 商品列表
    base.NilResponse ListProperty(1: PropertyListReq req) (api.post = "/service/property/list")

    // 添加商品
    base.NilResponse Create(1: CreateOrUpdateReq req) (api.post = "/service/product/create")
    // 编辑商品
    base.NilResponse Update(1: CreateOrUpdateReq req) (api.post = "/service/product/update")
    // 删除商品
    base.NilResponse Delete(1: base.IDReq req) (api.post = "/service/product/del")
    // 商品列表
    base.NilResponse List(1: ListReq req) (api.post = "/service/product/list")
    // 上架0/下架1
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/product/status")

    // 商品详情
    base.NilResponse InfoByID(1: base.IDReq req) (api.post = "/service/product/info")



}



