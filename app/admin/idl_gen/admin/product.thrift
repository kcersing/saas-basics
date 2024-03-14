namespace go admin.product

include "../base/base.thrift"
include "../base/data.thrift"

struct Property {
    1: i64 product_id
    2: string spu_name // spu名
    3: i64 spu_price // 定价
    4: i64 duration // 时长
    5: i64 length // 单次时长
    6: i64 count // 次数
    7: string type // 次数
    8: string data
}

struct Product {
    1: string name // 商品名
    2: string pic // 主图
    3: string description // 详情
    4: list <Property> property // 属性
    5: i64 price // 价格
    6: i64 stock // 库存
    7: i64 status // 商品状态
}

struct CreatePropertyReq {
    2: string spu_name // spu名
    3: i64 spu_price // 定价
    4: i64 duration // 时长
    5: i64 length // 单次时长
    6: i64 count // 次数
    7: string type // 次数
}

struct UpdatePropertyReq {
    1: required i64 property_id
    2: string spu_name // spu名
    3: i64 spu_price // 定价
    4: i64 duration // 时长
    5: i64 length // 单次时长
    6: i64 count // 次数
    7: string type // 次数
}

struct DeletePropertyReq {
    1: required i64 property_id
}

struct CreateReq {
    1: required string name // 商品名
    2: required string pic // 主图
    3: required string description // 详情
    4: required list <Property> property // 属性
    5: required i64 price // 价格
    6: required i64 stock // 库存
}

struct UpdateReq {
    1: required i64 product_id
    2: optional string name // 商品名
    3: optional string pic // 主图
    4: optional string description // 详情
    5: optional list <Property> property  // 属性
    6: optional i64 price // 价格
    7: optional i64 stock // 库存
}

struct ListReq {
    1: i64 page,
    2: i64 pageSize,
    3: optional string name
    4: optional string spu_name
    5: optional i64 status
}



service ProductService {
    // 添加属性
    base.NilResponse CreateProperty(1: CreatePropertyReq req) (api.post = "/api/admin/property/create")
    // 编辑属性
    base.NilResponse UpdateProperty(1: UpdatePropertyReq req) (api.post = "/api/admin/property/update")
    // 删除属性
    base.NilResponse DeleteProperty(1: base.IDReq req) (api.post = "/api/admin/property/del")
    // 商品列表
    base.NilResponse ListProperty(1: ListReq req) (api.post = "/api/admin/Property/list")

    // 添加商品
    base.NilResponse Create(1: CreateReq req) (api.post = "/api/admin/product/create")
    // 编辑商品
    base.NilResponse Update(1: UpdateReq req) (api.post = "/api/admin/product/update")
    // 删除商品
    base.NilResponse Delete(1: base.IDReq req) (api.post = "/api/admin/product/del")
    // 上架0/下架1
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/product/status")
    // 商品详情
    base.NilResponse InfoByID(1: base.IDReq req) (api.post = "/api/admin/product/info")
    // 商品列表
    base.NilResponse List(1: ListReq req) (api.post = "/api/admin/product/list")

}


