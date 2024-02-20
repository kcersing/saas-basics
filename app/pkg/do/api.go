package do

type Api interface {
	Create(req ApiInfo) error
	Update(req ApiInfo) error
	Delete(id uint64) error
	List(req ListApiReq) (resp []*ApiInfo, total int, err error)
}

type ApiInfo struct {
	ID          uint64
	CreatedAt   string
	UpdatedAt   string
	Path        string
	Description string
	Group       string
	Method      string
}

type ListApiReq struct {
	Page        uint64
	PageSize    uint64
	Path        string
	Description string
	Method      string
	Group       string
}
