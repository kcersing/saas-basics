package do

type Role interface {
	Create(req RoleInfo) error
	Update(req RoleInfo) error
	Delete(id uint64) error
	RoleInfoByID(ID uint64) (roleInfo *RoleInfo, err error)
	List(req *RoleListReq) (roleInfoList []*RoleInfo, total int, err error)
	UpdateStatus(ID uint64, status uint8) error
}

type RoleInfo struct {
	ID            uint64
	Name          string
	Value         string
	DefaultRouter string
	Status        uint64
	Remark        string
	OrderNo       uint32
	CreatedAt     string
	UpdatedAt     string
}

type RoleListReq struct {
	Page     uint64
	PageSize uint64
}
