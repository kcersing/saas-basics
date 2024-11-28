package do

import "saas/idl_gen/model/entry"

type EntryLogs interface {
	Create(logsReq *entry.EntryInfo) error
	List(req *entry.EntryListReq) (list []*entry.EntryInfo, total int, err error)
}
