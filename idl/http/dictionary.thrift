namespace go admin

include "../base/base.thrift"
include "../base/data.thrift"




// 字典列表请求数据
struct DictionaryPageReq {
    1:  string title (api.raw = "title" )
    2:  string name (api.raw = "name" )
    3:  string page (api.raw = "page" )
    4:  string pageSize (api.raw = "pageSize" )
}

//字典名获取字典键值请求数据
struct DictionaryDetailReq{
    1:  string name (api.raw = "title" )
}

// dictionary service
service dictionary {
  // 创建字典信息
  base.NilResponse CreateDictionary(1: data.DictionaryInfo req) (api.post = "/api/admin/dict/create")

  // 更新字典信息
  base.NilResponse UpdateDictionary(1: data.DictionaryInfo req) (api.post = "/api/admin/dict/update")

  // 删除字典信息
  base.NilResponse DeleteDictionary(1: base.IDReq req) (api.post =  "/api/admin/dict")

  // 获取字典列表
  base.NilResponse DictionaryList(1: DictionaryPageReq req) (api.get = "/api/admin/dict/list")

  // 创建字典键值信息
  base.NilResponse CreateDictionaryDetail(1: data.DictionaryDetail req) (api.post = "/api/admin/dict/detail/create")

  // 更新字典键值信息
  base.NilResponse UpdateDictionaryDetail(1: data.DictionaryDetail req) (api.get = "/api/admin/dict/detail/update")

  // 删除字典键值信息
  base.NilResponse DeleteDictionaryDetail(1: base.IDReq req) (api.get = "/api/admin/dict/detail")

  // 根据字典名获取字典键值列表
  base.NilResponse DetailByDictionaryName(1: DictionaryDetailReq req) (api.get = "/api/admin/dict/detail/list")

}