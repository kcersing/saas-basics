namespace go product

include "../base/base.thrift"



struct ProductInfo {
    1: optional i64 id=0  (api.raw = "id")
    2: optional string name  = ""(api.raw = "name")
    3: optional string pic= ""  (api.raw = "pic")
    4: optional string description  = ""(api.raw = "description")
    5: optional i64 stock=0  (api.raw = "stock")
    6: optional i64 status=0  (api.raw = "status")
    7: optional i64 duration =0 (api.raw = "duration")
    8: optional i64 length=0  (api.raw = "length")
    9: optional string type = "" (api.raw = "type")
    10: optional i64 deadline=0 (api.raw = "deadline")
    11: optional list<base.Sales> sales = {} (api.raw = "sales")
    12: optional i64 is_sales=0 (api.raw = "isSales")
    13: optional string signSalesAt = "" (api.raw = "signSalesAt")
    14: optional string endSalesAt = "" (api.raw = "endSalesAt")
    16: optional string createdAt = ""  (api.raw = "updatedAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")
    18: optional list<base.List> tags = {}  (api.raw = "tags")
    19: optional list<base.List> contracts = {} (api.raw = "contracts")
    20: optional i64 createdId = 0 (api.raw = "createdId")
    21: optional string createdName = "" (api.raw = "createdName")
    /**团课*/
    22: optional list<base.List> lessons = {} (api.raw = "lessons")
}


struct CreateOrUpdateProductReq {
    1: optional i64 id =0 (api.raw = "id")
    2: optional string name ="" (api.raw = "name")
    3: optional string pic=""  (api.raw = "pic")
    4: optional string description ="" (api.raw = "description")
    5: optional i64 stock =0 (api.raw = "stock")
    6: optional i64 status=0  (api.raw = "status")
    7: optional i64 duration =0 (api.raw = "duration")
    8: optional i64 length=0  (api.raw = "length")
    9: optional string type ="" (api.raw = "type")
    10: optional i64 deadline=0 (api.raw = "deadline")
    11: optional list<base.Sales> sales = {} (api.raw = "sales")
    12: optional i64 is_sales =1(api.raw = "isSales")
    13: optional string signSalesAt ="" (api.raw = "signSalesAt")
    14: optional string endSalesAt ="" (api.raw = "endSalesAt")
    16: optional list<i64> tagId=0   (api.raw = "tagId")
    17: optional list<i64> contractId =0 (api.raw = "contractId")
    18: optional list<i64> lessonsId =0 (api.raw = "lessonsId")
}

struct ProductListReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string name ="" (api.raw = "name")
    4: optional list<i64> status=0 (api.raw = "status")
    7: optional string type="" (api.raw = "type") // 类型
}

service ProductService {

    // 添加商品
    base.NilResponse CreateProduct(1: CreateOrUpdateProductReq req) (api.post = "/service/product/create")
    // 编辑商品
    base.NilResponse UpdateProduct(1: CreateOrUpdateProductReq req) (api.post = "/service/product/update")
    // 删除商品
    base.NilResponse DeleteProduct(1: base.IDReq req) (api.post = "/service/product/del")
    // 商品列表
    base.NilResponse ProductList(1: ProductListReq req) (api.post = "/service/product/list")
    // 上架0/下架1
    base.NilResponse UpdateProductStatus(1: base.StatusCodeReq req) (api.post = "/service/product/status")

    // 商品详情
    base.NilResponse ProductInfo(1: base.IDReq req) (api.post = "/service/product/info")

    base.NilResponse ProductListExport(1: ProductListReq req) (api.post = "/service/product/export")


}



