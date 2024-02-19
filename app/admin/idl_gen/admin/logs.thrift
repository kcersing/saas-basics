namespace go admin.logs

include "../base/base.thrift"
include "../base/data.thrift"

//日志列表请求数据
struct LogsListReq {
    1:  string page (api.raw = "page")
    2:  string limit (api.raw = "limit")
    3:  string type (api.raw = "type")
    4:  string method (api.raw = "method")
    5:  string api (api.raw = "api")
    6:  string success (api.raw = "success")
    7:  string operators (api.raw = "operators")
}


service LogsService{
  // Get logs list | 获取日志列表
  base.NilResponse GetLogsList(1: LogsListReq req) (api.post = "/api/admin/logs/list")

  // Delete logs | 删除日志信息
  base.NilResponse DeleteLogs(1: base.Empty req) (api.post = "/api/admin/logs/deleteAll")

}