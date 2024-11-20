namespace go system.sys

include "../base/base.thrift"

service SystemService {

     base.SysListResp RoleList(1: base.ListReq req) (api.post = "/service/sys/role/list")

     base.NilResponse Captcha(1: base.Empty req) (api.post = "/service/sys/captcha")

}



