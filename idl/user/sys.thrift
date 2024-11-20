namespace go user

include "../base/base.thrift"

service UserService {

    // 员工列表
    base.SysListResp StaffList(1: base.ListReq req) (api.post = "/service/sys/staff/list")

}



