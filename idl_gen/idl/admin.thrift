namespace go admin

include "../base/base.thrift"

service AdminService {
    //检查系统状态
    base.NilResponse HealthCheck(1: base.Empty req) (api.get = "/service/health")
    //获取验证码
    base.NilResponse Captcha(1: base.Empty req) (api.post = "/service/captcha")

}