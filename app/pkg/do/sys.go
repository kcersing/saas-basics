package do

type Sys interface {
	ProductList(req SysListReq) (list []SysList, total int64, err error)
	PropertyList(req SysListReq) (list []SysList, total int64, err error)
	PropertyType(req SysListReq) (list []SysList, total int64, err error)
	VenueList(req SysListReq) (list []SysList, total int64, err error)
	MemberList(req SysListReq) (list []SysList, total int64, err error)
	ContractList(req SysListReq) (list []SysList, total int64, err error)
	StaffList(req SysListReq) (list []SysList, total int64, err error)
}
type SysList struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}
type SysListReq struct {
	DictionaryId int64  `json:"dictionaryId"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	Type         string `json:"type"`
	Product      int64  `json:"product"`
}
