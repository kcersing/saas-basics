package do

type Contract interface {
	ContractList(req *MenuListReq) (list []*ContractInfo, total int, err error)
	ContractCreate(req *ContractInfo) error
	ContractUpdate(req *ContractInfo) error
	ContractUpdateStatus(ID int64, status int64) error
	ContractByID(id int64) (*ContractInfo, error)
}

type ContractInfo struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Status  int64  `json:"status"`
	Content string `json:"content"`
}
