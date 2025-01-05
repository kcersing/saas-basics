package do

import (
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/sys"
)

type Sys interface {
	ProductList(req sys.SysProductListReq) (list []SysProductList, total int64, err error)
	VenueList(req base.ListReq) (list []SysList, total int64, err error)
	MemberList(req base.ListReq) (list []SysList, total int64, err error)
	ContractList(req base.ListReq) (list []SysList, total int64, err error)
	StaffList(req base.ListReq) (list []SysList, total int64, err error)
	PlaceList(req base.ListReq) (list []SysList, total int64, err error)
	RoleList(req base.ListReq) (list []SysList, total int64, err error)
}
type SysList struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}
type SysProductList struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Length int64  `json:"Length"`
}
