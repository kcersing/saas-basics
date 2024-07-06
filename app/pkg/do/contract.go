package do

type Contract interface {
	Create(req *ContractInfo) error
	Update(req *ContractInfo) error
	UpdateStatus(ID int64, status int64) error
	Info(id int64) (info *ContractInfo, err error)
	List(req *ContractListReq) (list []*ContractInfo, total int, err error)
}

type ContractInfo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Status    int64  `json:"status"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type ContractListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Name     string `json:"name"`
}
