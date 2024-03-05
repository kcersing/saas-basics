package do

type Api interface {
	Create(req ApiInfo) error
	Update(req ApiInfo) error
	Delete(id int64) error
	List(req ListApiReq) (resp []*ApiInfo, total int, err error)
}

type ApiInfo struct {
	ID          int64  `json:"id,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	Path        string `json:"path,omitempty"`
	Description string `json:"description,omitempty"`
	Group       string `json:"group,omitempty"`
	Method      string `json:"method,omitempty"`
}

type ListApiReq struct {
	Page        int64  `json:"page,omitempty"`
	PageSize    int64  `json:"pageSize,omitempty"`
	Path        string `json:"path,omitempty"`
	Description string `json:"description,omitempty"`
	Method      string `json:"method,omitempty"`
	Group       string `json:"group,omitempty"`
}
