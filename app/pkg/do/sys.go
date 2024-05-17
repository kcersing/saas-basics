package do

type Sys interface {
	ProductList(req SysListReq) (list []SysList, err error)
	PropertyList(req SysListReq) (list []SysList, err error)
	PropertyType(req SysListReq) (list []SysList, err error)
	VenueList(req SysListReq) (list []SysList, err error)
	MemberList(req SysListReq) (list []SysList, err error)
	ContractList(req SysListReq) (list []SysList, err error)
	StaffList(req SysListReq) (list []SysList, err error)
}
type SysList struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type SysListReq struct {
	Page         int64  `json:"page"`
	PageSize     int64  `json:"pageSize"`
	DictionaryId int64  `json:"dictionaryId"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
}
