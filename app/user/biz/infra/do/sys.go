package do

import (
	"rpc_gen/kitex_gen/base"
)

type Sys interface {
	StaffList(req base.ListReq) (list []base.SysList, total int64, err error)
}
