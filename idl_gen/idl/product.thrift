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
    12: optional i64 isSales=0 (api.raw = "isSales")
    13: optional string signSalesAt = "" (api.raw = "signSalesAt")
    14: optional string endSalesAt = "" (api.raw = "endSalesAt")
    16: optional string createdAt = ""  (api.raw = "createdAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")
    /**标签*/
    18: optional list<base.List> tags = {}  (api.raw = "tags")
    /**合同*/
    19: optional list<base.List> contracts = {} (api.raw = "contracts")
    20: optional i64 createdId = 0 (api.raw = "createdId")
    21: optional string createdName = "" (api.raw = "createdName")
    /**团课*/
    22: optional list<base.CourseList> lessons =0 (api.raw = "lessons")
    /**团课预约 1支持2不支持*/
    23: optional i64 isLessons =0 (api.raw = "isLessons")
    /**次级类型courseOne一对一私教课 courseMore一对多私教课 cardTerm期限卡 cardSub次卡 lessons团课 coursePackage私教课包*/
    24: optional string subType ="" (api.raw = "subType")
    /**价格(单个/卡价格)*/
    25: optional double price = 0 (api.raw = "price")
    /**次数(次卡)*/
    26:optional i64 times = 0 (api.raw = "times")

     /**课程 课程 1不限2指定*/
     27: optional i64 isCourse =0 (api.raw = "isCourse")
     /**课程-数组*/
     28: optional list<base.CourseList> courses = {}  (api.raw = "courses")
}


struct CreateOrUpdateProductReq {
    1: optional i64 id =0 (api.raw = "id")
    /**商品名称*/
    2: optional string name ="" (api.raw = "name")
    /**商品图片*/
    3: optional string pic=""  (api.raw = "pic")
    /**商品详情*/
    4: optional string description ="" (api.raw = "description")
    /**库存*/
    5: optional i64 stock =0 (api.raw = "stock")
    /**状态[1:未上架上架,2:上架]*/
    6: optional i64 status=1  (api.raw = "status")
    /**有效期(卡期限/课单节期限)*/
    7: optional i64 duration =0 (api.raw = "duration")
    /**课程课时(课)*/
    8: optional i64 length=0  (api.raw = "length")
    /**类型 card:卡 course:课 coursePackage:课包 Lessons:团课 */
    9: optional string type ="" (api.raw = "type")
     /**激活期限*/
    10: optional i64 deadline=0 (api.raw = "deadline")
    /**销售信息数组(多个/课价格信息)*/
    11: optional list<base.Sales> sales = {} (api.raw = "sales")
    /**销售方式 1会员端*/
    12: optional i64 isSales = 0 (api.raw = "isSales")
    /**销售开始时间*/
    13: optional string signSalesAt ="" (api.raw = "signSalesAt")
    /**销售结束时间*/
    14: optional string endSalesAt ="" (api.raw = "endSalesAt")
    /**标签-数组*/
    16: optional list<i64> tagId=0   (api.raw = "tagId")
    /**合同-数组*/
    17: optional list<i64> contractId =0 (api.raw = "contractId")
    /**团课-数组*/
    18: optional list<base.CourseList> lessons =0 (api.raw = "lessons")
    /**团课预约 1支持2不支持*/
    19: optional i64 isLessons =0 (api.raw = "isLessons")
    /**次级类型courseOne一对一私教课 courseMore一对多私教课 cardTerm期限卡 cardSub次卡 lessons团课 coursePackage私教课包*/
    20: optional string subType ="" (api.raw = "subType")
    /**价格(单个/卡价格)*/
    21: optional double price = 0 (api.raw = "price")
    /**次数(次卡)*/
    22:optional i64 times = 0 (api.raw = "times")

    /**课程 课程 1不限2指定*/
    23: optional i64 isCourse =0 (api.raw = "isCourse")
    /**课程-数组*/
    24: optional list<base.CourseList> courses = {}  (api.raw = "courses")

}

struct ProductListReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string name ="" (api.raw = "name")
    4: optional list<i64> status=0 (api.raw = "status")
    7: optional string type="" (api.raw = "type") // 类型
   /**次级类型*/
    8: optional string subType ="" (api.raw = "subType")
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



