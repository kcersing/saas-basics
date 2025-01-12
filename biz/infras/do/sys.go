package do

import (
	"saas/idl_gen/model/base"
	"saas/idl_gen/model/sys"
)

type Sys interface {
	ProductList(req sys.SysProductListReq) (list []SysProductList, total int64, err error)
	VenueList(req base.ListReq) (list []SysList, total int64, err error)
	MemberList(req sys.SysMemberListReq) (list []SysMemberList, total int64, err error)
	ContractList(req base.ListReq) (list []SysList, total int64, err error)
	StaffList(req sys.SysStaffListReq) (list []SysStaffList, total int64, err error)
	PlaceList(req sys.SysPlaceListReq) (list []SysPlaceList, total int64, err error)
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
	Tags   []int  `json:"tags"`
}

type SysStaffList struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	Tags      []int    `json:"tags"`
	Functions []string `json:"functions"`
}
type SysMemberList struct {
	Id            int64                  `json:"id"`
	Name          string                 `json:"name"`
	Mobile        string                 `json:"mobile"`
	MemberProduct []SysMemberProductList `json:"memberProduct"`
}
type SysMemberProductList struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	NumberSurplus int64  `json:"numberSurplus"`
	IsCourse      int64  `json:"isCourse"`
	Type          string `json:"type"`
}

type SysPlaceList struct {
	Id       int64          `json:"id"`
	Name     string         `json:"name"`
	Products []int          `json:"products"`
	Seat     [][]*base.Seat `json:"seat"`
	Number   int64          `json:"number"`
	Classify int64          `json:"classify"`
	Type     int64          `json:"type"`
}
