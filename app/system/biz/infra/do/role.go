package do

type Role interface {
	Create(req RoleInfo) error
	Update(req RoleInfo) error
	Delete(id int64) error
	RoleInfoByID(ID int64) (roleInfo *RoleInfo, err error)
	List(req *RoleListReq) (roleInfoList []*RoleInfo, total int, err error)
	UpdateStatus(ID int64, status int64) error
}

type RoleInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Value         string `json:"value"`
	DefaultRouter string `json:"defaultRouter"`
	Status        int64  `json:"status"`
	Remark        string `json:"remark"`
	OrderNo       int32  `json:"orderNo"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Menus         []int  `json:"menus"`
	Apis          []int  `json:"apis"`
}

type RoleListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}
