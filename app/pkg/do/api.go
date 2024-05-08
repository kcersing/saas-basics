package do

type Api interface {
	Create(req ApiInfo) error
	Update(req ApiInfo) error
	Delete(id int64) error
	List(req ListApiReq) (resp []*ApiInfo, total int, err error)
}

type ApiInfo struct {
	ID          int64  `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Method      string `json:"method"`
}

type ListApiReq struct {
	Page        int64  `json:"page"`
	PageSize    int64  `json:"pageSize"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Method      string `json:"method"`
	Group       string `json:"group"`
}
