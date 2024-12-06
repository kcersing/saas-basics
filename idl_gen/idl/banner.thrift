
namespace go banner

include "../base/base.thrift"

// Banner
service BannerService {

  base.NilResponse CreateBanner(1: BannerInfo req) (api.post = "/service/banner/create")

  base.NilResponse UpdateBanner(1: BannerInfo req) (api.post = "/service/banner/update")

  base.NilResponse DelBanner(1: base.IDReq req) (api.post = "/service/banner/del")

  base.NilResponse BannerList(1: base.PageInfoReq req) (api.post = "/service/banner/list")

  base.NilResponse UpdateBannerShow(1: base.StatusCodeReq req) (api.post = "/service/banner/show")
}

struct BannerInfo {
  1: string name="" (api.raw = "name")
  2: string pic="" (api.raw = "pic")
  3: string link="" (api.raw = "link")
  4: i64 isShow=1 (api.raw = "isShow")
}

