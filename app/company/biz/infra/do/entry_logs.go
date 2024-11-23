package do

import "rpc_gen/kitex_gen/company/entry"

type EntryLogs interface {
	Create(logsReq *entry.EntryInfo) error
	List(req *entry.EntryListReq) (list []*entry.EntryInfo, total int, err error)
}
