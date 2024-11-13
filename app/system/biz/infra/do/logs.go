package do

import "rpc_gen/kitex_gen/system/auth"

type Logs interface {
	Create(logsReq *auth.LogsInfo) error
	List(req *auth.LogsListReq) (list []*auth.LogsInfo, total int64, err error)
	DeleteAll(ids []int64) error
}
