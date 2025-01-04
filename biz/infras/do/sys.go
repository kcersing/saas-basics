package do

import "saas/idl_gen/model/base"

type Sys interface {
	ProductList(req base.ListReq) (list []SysList, total int64, err error)
	VenueList(req base.ListReq) (list []SysList, total int64, err error)
	MemberList(req base.ListReq) (list []SysList, total int64, err error)
	ContractList(req base.ListReq) (list []SysList, total int64, err error)
	StaffList(req base.ListReq) (list []SysList, total int64, err error)
	PlaceList(req base.ListReq) (list []SysList, total int64, err error)
	RoleList(req base.ListReq) (list []SysList, total int64, err error)
}
type SysList struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}
