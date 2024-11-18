package do

import "rpc_gen/kitex_gen/system/sys"

type Sys interface {
	ProductList(req sys.ListReq) (list []sys.SysList, total int64, err error)
	PropertyList(req sys.ListReq) (list []sys.SysList, total int64, err error)
	PropertyType(req sys.ListReq) (list []sys.SysList, total int64, err error)
	VenueList(req sys.ListReq) (list []sys.SysList, total int64, err error)
	MemberList(req sys.ListReq) (list []sys.SysList, total int64, err error)
	ContractList(req sys.ListReq) (list []sys.SysList, total int64, err error)
	StaffList(req sys.ListReq) (list []sys.SysList, total int64, err error)
	PlaceList(req sys.ListReq) (list []sys.SysList, total int64, err error)
	RoleList(req sys.ListReq) (list []sys.SysList, total int64, err error)
}
