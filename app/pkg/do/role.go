package do

type Role interface {
	Create(req RoleInfo) error
	Update(req RoleInfo) error
	Delete(id int64) error
	RoleInfoByID(ID int64) (roleInfo *RoleInfo, err error)
	List(req *RoleListReq) (roleInfoList []*RoleInfo, total int, err error)
	UpdateStatus(ID int64, status int8) error
}

type RoleInfo struct {
	ID            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Value         string `json:"value,omitempty"`
	DefaultRouter string `json:"defaultRouter,omitempty"`
	Status        int64  `json:"status,omitempty"`
	Remark        string `json:"remark,omitempty"`
	OrderNo       int32  `json:"orderNo,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
}

type RoleListReq struct {
	Page     int64 `json:"page,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}
