package do

import (
	"rpc_gen/kitex_gen/base"
)

type Sys interface {
	ProductList(req base.ListReq) (list []base.SysList, total int64, err error)
	PropertyList(req base.ListReq) (list []base.SysList, total int64, err error)
	PropertyType(req base.ListReq) (list []base.SysList, total int64, err error)
	VenueList(req base.ListReq) (list []base.SysList, total int64, err error)
	MemberList(req base.ListReq) (list []base.SysList, total int64, err error)
	ContractList(req base.ListReq) (list []base.SysList, total int64, err error)
	StaffList(req base.ListReq) (list []base.SysList, total int64, err error)
	PlaceList(req base.ListReq) (list []base.SysList, total int64, err error)
	RoleList(req base.ListReq) (list []base.SysList, total int64, err error)
}
