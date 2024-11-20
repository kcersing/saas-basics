namespace go system.sys

include "../base/base.thrift"

service ProductService {
    // 商品列表
    base.SysListResp ProductList(1: base.ListReq req) (api.post = "/service/sys/product/list")

    // 属性列表
    base.SysListResp PropertyList(1: base.ListReq req) (api.post = "/service/sys/property/list")

    // 属性类型
    base.SysListResp PropertyType(1: base.ListReq req) (api.post = "/service/sys/property/type")

}



