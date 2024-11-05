namespace go admin.admin

include "../base/base.thrift"
include "../base/data.thrift"


service AdminService {
    //检查系统状态
    base.NilResponse HealthCheck(1: base.Empty req) (api.get = "/api/health")
    //获取验证码
    base.NilResponse Captcha(1: base.Empty req) (api.post = "/api/captcha")

}